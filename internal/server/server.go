package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	healthFile string
	router     *httprouter.Router
	token      string
	tokenFile  string
}

func NewServer(healthFile string, tokenFile string) (*server, error) {
	router := httprouter.New()

	s := &server{
		healthFile: healthFile,
		router:     router,
		tokenFile:  tokenFile,
	}

	s.updateToken()
	router.GET("/randomize", s.randomize)
	router.GET("/up", s.healthcheck)

	return s, nil
}

func (s *server) GetRouter() *httprouter.Router {
	return s.router
}

func (s *server) randomize(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("received error while reading body: ", err)
		w.WriteHeader(400)
		return
	}

	var names []string
	if err = json.Unmarshal(bytes, &names); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Could not unmarshal body into list of strings"))
		return
	}

}
