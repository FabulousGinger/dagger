package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	wavingHand  = '\U0001f44b'
	worldMap    = '\U0001f5fa'
	YellowHeart = '\U0001f49b'
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got / request from %s\n", r.RemoteAddr)
		io.WriteString(w, greeting())
	})

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func greeting() string {

	// return fmt.Sprintf("%c %c", wavingHand, worldMap)
	return fmt.Sprintf("%c", YellowHeart)
}
