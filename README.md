# notes-api

Go get some dependencies:

```
go get github.com/labstack/echo
go get github.com/mattn/go-sqlite3
```

Then run
```
go run main.go
```

And navigate to [http://localhost:3000](http://localhost:3000).


## Test json api

I'm using [Httpie](https://httpie.org/)

Get some notes
```
http localhost:3000/notes
```
Create a note
```
http POST localhost:3000/notes content="interesting note" 
```






