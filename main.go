package main

import (
	"github.com/awouda/notes-api/domain"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"net/http"
)

type Note domain.Note

func main() {

	//example commandline parsing, maybe we need it later
	//var count = flag.Int("count", 5, "the count of items")
	//flag.Parse()
	//fmt.Println("count value ", *count)
	domain.InitDb()
	domain.DB.AutoMigrate(&Note{})

	defer domain.DB.Close()

	e := echo.New()

	n := e.Group("/notes")
	n.GET("", listNotes)
	n.POST("", createNote)

	e.File("/", "public/index.html")

	e.Logger.Fatal(e.Start(":3000"))
}

func listNotes(c echo.Context) error {

	domain.DB.Find(&notes)
	return c.JSON(http.StatusOK, &notes)

}

func createNote(c echo.Context) (err error) {
	n := new(Note)

	if err = c.Bind(n); err != nil {
		return
	}

	dbNewNote(n)
	return c.JSON(http.StatusCreated, n)
}

func dbNewNote(note *Note) (*Note, error) {

	domain.DB.Create(note)

	return note, nil
}

var notes []*Note
