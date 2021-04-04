// Package state defines states for domain types.
package state

// Issue represents the possible states of an issue.
type Issue string

// The possible states of an issue.
const (
	IssueOpen   Issue = "open"   // An issue that is still open.
	IssueClosed Issue = "closed" // An issue that has been closed.
)

// Change represents the possible states of a change.
type Change string

// The possible states of a change.
const (
	ChangeOpen   Change = "open"   // A change that is still open.
	ChangeClosed Change = "closed" // A change that has been closed without being merged.
	ChangeMerged Change = "merged" // A change that has been closed by being merged.
)
