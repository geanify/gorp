package editor

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/browser"
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	err := browser.OpenURL("localhost:8080")

	if err != nil {
		log.Println("Could not open browser")
	}
}
