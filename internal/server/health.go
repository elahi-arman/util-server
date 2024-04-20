package server

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// healthcheck opens the healthCheck file established in the constructor
// and attempts to read it. Any error in reading (can't open or can't parse)
// results in the healthcheck failing. If the status code can be parsed,
// then the written status code will be returned.
func (s *server) healthcheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	f, err := os.OpenFile(s.healthFile, os.O_RDONLY, 0443)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	defer f.Close()
	status, err := io.ReadAll(f)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	statusCode, err := strconv.Atoi(string(status))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(statusCode)
}
