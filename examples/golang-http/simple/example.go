package main

import (
	"fmt"
	"net/http"
)

func myHttpHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/localreply_by_config" {
		body := fmt.Sprintf("%s, path: %s\r\n", "f.config.echoBody", path)
		w.Write([]byte(body))
	}

}
