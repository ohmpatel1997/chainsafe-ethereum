package httphandler

import (
	"log"
	"net/http"

	"github.com/shurcooL/home/httputil"
	"github.com/shurcooL/httperror"
	"github.com/shurcooL/reactions"
)

// Reactions is an API handler for reactions.Service.
type Reactions struct {
	Reactions reactions.Service
}

func (h Reactions) List(w http.ResponseWriter, req *http.Request) error {
	if req.Method != http.MethodGet {
		return httperror.Method{Allowed: []string{http.MethodGet}}
	}
	reactableURL := req.URL.Query().Get("ReactableURL")
	reactions, err := h.Reactions.List(req.Context(), reactableURL)
	if err != nil {
		return err
	}
	return httperror.JSONResponse{V: reactions}
}

func (h Reactions) GetOrToggle(w http.ResponseWriter, req *http.Request) error {
	if err := httputil.AllowMethods(req, http.MethodGet, http.MethodPost); err != nil {
		return err
	}
	if err := req.ParseForm(); err != nil {
		log.Println("req.ParseForm:", err)
		return httperror.BadRequest{Err: err}
	}
	reactableURL := req.Form.Get("reactableURL")
	reactableID := req.Form.Get("reactableID")
	switch req.Method {
	case http.MethodGet:
		reactions, err := h.Reactions.Get(req.Context(), reactableURL, reactableID)
		if err != nil {
			return err
		}
		return httperror.JSONResponse{V: reactions}
	case http.MethodPost:
		tr := reactions.ToggleRequest{
			Reaction: reactions.EmojiID(req.PostForm.Get("reaction")),
		}
		reactions, err := h.Reactions.Toggle(req.Context(), reactableURL, reactableID, tr)
		if err != nil {
			return err
		}
		return httperror.JSONResponse{V: reactions}
	default:
		panic("unreachable")
	}
}
