# minimux
Just another URL pattern muxer


# Example

```
package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gsirbiladze/minimux"
)

func main() {
	mux := minimux.New()

	// matches "/path/*". NOT "/path"
	mux.Get("/path/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "this is '"+r.URL.Path+"'")
	}))

	// matches exact "/path1"
	mux.Get("/path1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "this is '"+r.URL.Path+"'")
	}))

	// matches exact "/path1/path2"
	mux.Get("/path1/path2", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "this is '"+r.URL.Path+"'")
	}))

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}
}
```
