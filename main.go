package main

import (
	"errors"
	"fmt"
	"github.com/awouda/notes-api/domain"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"math/rand"
	"net/http"
	"strings"
)


type AppNote domain.Note



func main() {

    //example commandline parsing, maybe we need it later
	//var count = flag.Int("count", 5, "the count of items")
	//flag.Parse()
	//fmt.Println("count value ", *count)



	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/notes", func(r chi.Router) {
		r.Get("/", ListNotes)
		r.Post("/", CreateNote)
	})

	http.ListenAndServe(":3000", r)

}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, NewArticleListResponse(notes)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	data := &NoteRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	article := data.AppNote
	dbNewNote(article)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewNoteResponse(article))
}


func NewNoteResponse(note *AppNote) *NoteResponse {
	resp := &NoteResponse{AppNote: note}

	return resp
}


type NoteListResponse []*NoteResponse

func NewArticleListResponse(notes []*AppNote) []render.Renderer {
	list := []render.Renderer{}
	for _, note := range notes {
		list = append(list, NewNoteResponse(note))
	}
	return list
}

func dbNewNote(note *AppNote) (string, error) {
	note.ID = fmt.Sprintf("%d", rand.Intn(100)+10)
	notes = append(notes, note)
	fmt.Println(len(notes))
	return note.ID, nil
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

//func (n *domain.Note) Render(w http.ResponseWriter, r *http.Request) render.Renderer {
//	return nil
//}



type NoteResponse struct {
	*AppNote
	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}
type NoteRequest struct {
	*AppNote
	// We add an additional field to the response here.. such as this
	// elapsed computed property
	User string `json:"user,omitempty"`
}

func (n *NoteRequest) Bind(r *http.Request) error {
	// a.Article is nil if no Article fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if n.AppNote == nil {
		return errors.New("missing required Note fields.")
	}


	if n.User == "" {
       fmt.Println("user was empty")
	} else {
		fmt.Println("user was ", n.User)
	}



	n.AppNote.Content = strings.ToLower(n.Content) // as an example, we down-case
	return nil
}

func (rd *NoteResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}




func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}



var notes = []*AppNote{
	{ID: "1", Content:"notitie 1"},
	{ID: "2", Content:"notitie 2"},
	{ID: "3", Content:"notitie 3"},
}
