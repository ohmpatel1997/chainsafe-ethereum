// Package gerritapi implements a read-only change.Service using Gerrit API client.
package gerritapi

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"

	"dmitri.shuralyov.com/service/change"
	"github.com/andygrunwald/go-gerrit"
	"github.com/shurcooL/users"
)

// NewService creates a Gerrit-backed issues.Service using given Gerrit client.
// client must be non-nil.
func NewService(client *gerrit.Client) change.Service {
	return service{
		cl:     client,
		domain: client.BaseURL().Host,
	}
}

type service struct {
	cl     *gerrit.Client
	domain string
}

func (s service) List(ctx context.Context, rs string, opt change.ListOptions) ([]change.Change, error) {
	project := project(rs)
	var query string
	switch opt.Filter {
	case change.FilterOpen:
		query = fmt.Sprintf("project:%s status:open", project)
	case change.FilterClosedMerged:
		// "status:closed" is equivalent to "(status:abandoned OR status:merged)".
		query = fmt.Sprintf("project:%s status:closed", project)
	case change.FilterAll:
		query = fmt.Sprintf("project:%s", project)
	}
	cs, _, err := s.cl.Changes.QueryChanges(&gerrit.QueryChangeOptions{
		QueryOptions: gerrit.QueryOptions{
			Query: []string{query},
			Limit: 25,
		},
		ChangeOptions: gerrit.ChangeOptions{
			AdditionalFields: []string{"DETAILED_ACCOUNTS", "MESSAGES"},
		},
	})
	if err != nil {
		return nil, err
	}
	var is []change.Change
	for _, chg := range *cs {
		if chg.Status == "DRAFT" {
			continue
		}
		is = append(is, change.Change{
			ID:        uint64(chg.Number),
			State:     state(chg.Status),
			Title:     chg.Subject,
			Author:    s.gerritUser(chg.Owner),
			CreatedAt: chg.Created.Time,
			Replies:   len(chg.Messages),
		})
	}
	//sort.Sort(sort.Reverse(byID(is))) // For some reason, IDs don't completely line up with created times.
	sort.Slice(is, func(i, j int) bool {
		return is[i].CreatedAt.After(is[j].CreatedAt)
	})
	return is, nil
}

func (s service) Count(_ context.Context, repo string, opt change.ListOptions) (uint64, error) {
	// TODO.
	return 0, nil
}

func (s service) Get(ctx context.Context, _ string, id uint64) (change.Change, error) {
	chg, _, err := s.cl.Changes.GetChange(fmt.Sprint(id), &gerrit.ChangeOptions{
		AdditionalFields: []string{"DETAILED_ACCOUNTS", "MESSAGES", "ALL_REVISIONS"},
	})
	if err != nil {
		return change.Change{}, err
	}
	if chg.Status == "DRAFT" {
		return change.Change{}, os.ErrNotExist
	}
	return change.Change{
		ID:           id,
		State:        state(chg.Status),
		Title:        chg.Subject,
		Author:       s.gerritUser(chg.Owner),
		CreatedAt:    chg.Created.Time,
		Replies:      len(chg.Messages),
		Commits:      len(chg.Revisions),
		ChangedFiles: 0, // TODO.
	}, nil
}

func state(status string) change.State {
	switch status {
	case "NEW":
		return change.OpenState
	case "ABANDONED":
		return change.ClosedState
	case "MERGED":
		return change.MergedState
	case "DRAFT":
		panic("not sure how to deal with DRAFT status")
	default:
		panic("unreachable")
	}
}

func (s service) ListCommits(ctx context.Context, _ string, id uint64) ([]change.Commit, error) {
	chg, _, err := s.cl.Changes.GetChange(fmt.Sprint(id), &gerrit.ChangeOptions{
		AdditionalFields: []string{"DETAILED_ACCOUNTS", "ALL_REVISIONS"},
		//AdditionalFields: []string{"ALL_REVISIONS", "ALL_COMMITS"}, // TODO: Consider using git committer/author instead...
	})
	if err != nil {
		return nil, err
	}
	if chg.Status == "DRAFT" {
		return nil, os.ErrNotExist
	}
	commits := make([]change.Commit, len(chg.Revisions))
	for sha, r := range chg.Revisions {
		commits[r.Number-1] = change.Commit{
			SHA:     sha,
			Message: fmt.Sprintf("Patch Set %d", r.Number),
			// TODO: r.Uploader and r.Created describe the committer, not author.
			Author:     s.gerritUser(r.Uploader),
			AuthorTime: r.Created.Time,
		}
	}
	return commits, nil
}

func (s service) GetDiff(ctx context.Context, _ string, id uint64, opt *change.GetDiffOptions) ([]byte, error) {
	switch opt {
	case nil:
		diff, _, err := s.cl.Changes.GetPatch(fmt.Sprint(id), "current", nil)
		if err != nil {
			return nil, err
		}
		return []byte(*diff), nil
	default:
		chg, _, err := s.cl.Changes.GetChange(fmt.Sprint(id), &gerrit.ChangeOptions{
			AdditionalFields: []string{"ALL_REVISIONS"},
		})
		if err != nil {
			return nil, err
		}
		if chg.Status == "DRAFT" {
			return nil, os.ErrNotExist
		}
		r, ok := chg.Revisions[opt.Commit]
		if !ok {
			return nil, os.ErrNotExist
		}
		var base string
		switch r.Number {
		case 1:
			base = ""
		default:
			base = fmt.Sprint(r.Number - 1)
		}
		files, _, err := s.cl.Changes.ListFiles(fmt.Sprint(id), opt.Commit, &gerrit.FilesOptions{
			Base: base,
		})
		if err != nil {
			return nil, err
		}
		var sortedFiles []string
		for file := range files {
			sortedFiles = append(sortedFiles, file)
		}
		sort.Strings(sortedFiles)
		var diff string
		for _, file := range sortedFiles {
			diffInfo, _, err := s.cl.Changes.GetDiff(fmt.Sprint(id), opt.Commit, file, &gerrit.DiffOptions{
				Base:    base,
				Context: "5",
			})
			if err != nil {
				return nil, err
			}
			diff += strings.Join(diffInfo.DiffHeader, "\n") + "\n"
			for i, c := range diffInfo.Content {
				if i == 0 {
					diff += "@@ -154,6 +154,7 @@\n" // TODO.
				}
				switch {
				case len(c.AB) > 0:
					if len(c.AB) <= 10 {
						for _, line := range c.AB {
							diff += " " + line + "\n"
						}
					} else {
						switch i {
						case 0:
							for _, line := range c.AB[len(c.AB)-5:] {
								diff += " " + line + "\n"
							}
						default:
							for _, line := range c.AB[:5] {
								diff += " " + line + "\n"
							}
							diff += "@@ -154,6 +154,7 @@\n" // TODO.
							for _, line := range c.AB[len(c.AB)-5:] {
								diff += " " + line + "\n"
							}
						case len(diffInfo.Content) - 1:
							for _, line := range c.AB[:5] {
								diff += " " + line + "\n"
							}
						}
					}
				case len(c.A) > 0 || len(c.B) > 0:
					for _, line := range c.A {
						diff += "-" + line + "\n"
					}
					for _, line := range c.B {
						diff += "+" + line + "\n"
					}
				}
			}
		}
		return []byte(diff), nil
	}
}

func (s service) ListTimeline(ctx context.Context, _ string, id uint64, opt *change.ListTimelineOptions) ([]interface{}, error) {
	// TODO: Pagination. Respect opt.Start and opt.Length, if given.

	chg, _, err := s.cl.Changes.GetChangeDetail(fmt.Sprint(id), nil)
	if err != nil {
		return nil, err
	}
	comments, _, err := s.cl.Changes.ListChangeComments(fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	var timeline []interface{}
	timeline = append(timeline, change.Comment{ // CL description.
		ID:        "0",
		User:      s.gerritUser(chg.Owner),
		CreatedAt: chg.Created.Time,
		Body:      "", // THINK: Include commit message or no?
		Editable:  false,
	})
	for idx, message := range chg.Messages {
		if strings.HasPrefix(message.Tag, "autogenerated:") {
			switch message.Tag[len("autogenerated:"):] {
			case "gerrit:merged":
				timeline = append(timeline, change.TimelineItem{
					Actor:     s.gerritUser(message.Author),
					CreatedAt: message.Date.Time,
					Payload: change.MergedEvent{
						CommitID: message.Message[46:86], // TODO: Make safer.
						RefName:  chg.Branch,
					},
				})
			}
			continue
		}
		labels, body, ok := parseMessage(message.Message)
		if !ok {
			continue
		}
		var cs []change.InlineComment
		for file, comments := range *comments {
			for _, c := range comments {
				if c.Updated.Equal(message.Date.Time) {
					cs = append(cs, change.InlineComment{
						File: file,
						Line: c.Line,
						Body: c.Message,
					})
				}
			}
		}
		timeline = append(timeline, change.Review{
			ID:        fmt.Sprint(idx), // TODO: message.ID is not uint64; e.g., "bfba753d015916303152305cee7152ea7a112fe0".
			User:      s.gerritUser(message.Author),
			CreatedAt: message.Date.Time,
			State:     reviewState(labels),
			Body:      body,
			Editable:  false,
			Comments:  cs,
		})
	}
	return timeline, nil
}

func parseMessage(m string) (labels string, body string, ok bool) {
	// "Patch Set ".
	if !strings.HasPrefix(m, "Patch Set ") {
		return "", "", false
	}
	m = m[len("Patch Set "):]

	// "123".
	i := strings.IndexFunc(m, func(c rune) bool { return !unicode.IsNumber(c) })
	if i == -1 {
		return "", "", false
	}
	m = m[i:]

	// ":".
	if len(m) < 1 || m[0] != ':' {
		return "", "", false
	}
	m = m[1:]

	switch i = strings.IndexByte(m, '\n'); i {
	case -1:
		labels = m
	default:
		labels = m[:i]
		body = m[i+1:]
	}

	if labels != "" {
		// " ".
		if len(labels) < 1 || labels[0] != ' ' {
			return "", "", false
		}
		labels = labels[1:]
	}

	if body != "" {
		// "\n".
		if len(body) < 1 || body[0] != '\n' {
			return "", "", false
		}
		body = body[1:]
	}

	return labels, body, true
}

func reviewState(labels string) change.ReviewState {
	for _, label := range strings.Split(labels, " ") {
		switch label {
		case "Code-Review+2":
			return change.Approved
		case "Code-Review-2":
			return change.ChangesRequested
		}
	}
	return change.Commented
}

func (service) EditComment(_ context.Context, repo string, id uint64, cr change.CommentRequest) (change.Comment, error) {
	return change.Comment{}, fmt.Errorf("EditComment: not implemented")
}

func (s service) gerritUser(user gerrit.AccountInfo) users.User {
	var avatarURL string
	for _, avatar := range user.Avatars {
		if avatar.Height == 100 {
			avatarURL = avatar.URL
		}
	}
	return users.User{
		UserSpec: users.UserSpec{
			ID:     uint64(user.AccountID),
			Domain: s.domain,
		},
		Login: user.Name, //user.Username, // TODO.
		Name:  user.Name,
		//Email:     user.Email,
		AvatarURL: avatarURL,
	}
}

func project(repo string) string {
	i := strings.IndexByte(repo, '/')
	if i == -1 {
		return ""
	}
	return repo[i+1:]
}
