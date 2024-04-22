package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/elahi-arman/util-server/internal/groupings"
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
	router.GET("/randomize", s.WithTokenVerification(s.randomize))
	router.GET("/up", s.WithTokenVerification(s.healthcheck))

	return s, nil
}

func (s *server) GetRouter() *httprouter.Router {
	return s.router
}

func (s *server) randomize(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("received error while reading body: ", err)
		w.WriteHeader(400)
		return
	}

	var names []string
	if err = json.Unmarshal(bytes, &names); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Could not unmarshal body into list of strings"))
		return
	}

	queryParams := r.URL.Query()
	groupCount := 2
	if groupsParam := queryParams.Get("groups"); groupsParam != "" {
		p, err := strconv.Atoi(groupsParam)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("expected groups query param to be a parseable integer"))
			return
		}

		if p > 2 {
			groupCount = p
		}
	}

	output := groupings.RandomEvenLists(names, groupCount)
	d, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("unable to marshal ouptut into a list", err)
	}

	w.Write(d)
}
