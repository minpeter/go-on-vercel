package api

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "subpath - pong!")
}
