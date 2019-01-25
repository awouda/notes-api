package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)


	r.Route("/notes", func(r chi.Router) {
			r.Get("/", listNotes)
		})


	http.ListenAndServe(":3000", r	)







}





func listNotes(w http.ResponseWriter, r *http.Request) {

	render.Render(w, r, NewNoteResponse(notes[0]))
	//if err := render.Render(w, r, Note{"a","aaa"}.Render(w,r)); err != nil {
	//	render.Render(w, r, ErrRender(err))
	//	return
	//}
}

func NewNoteResponse(note *Note) *NoteResponse {
	resp := &NoteResponse{Note: note}


	return resp
}



type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}



func (n *Note) Render(w http.ResponseWriter, r *http.Request) render.Renderer {
	return nil
}

type NoteResponse struct {
	*Note
	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

func (rd *NoteResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}





