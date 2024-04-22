package server

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

const (
	token_cookie = "vt"
)

func (s *server) WithTokenVerification(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookie, err := r.Cookie(token_cookie)

		// would only return ErrNoCookie according to documentation so
		// no cookie would mean the request is unauthenticated
		if err != nil {
			fmt.Println("no vt cookie supplied")
			w.WriteHeader(401)
			return
		}

		if cookie.Valid() != nil {
			fmt.Println("cookie was not valid")
			w.WriteHeader(403)
			return
		}

		if cookie.Value != s.token {
			fmt.Printf("token is invalid. expected %s but received %s\n", s.token, cookie.Value)
			w.WriteHeader(403)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     token_cookie,
			Value:    s.token,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})

		h(w, r, ps)
	}
}

// updateToken updates the server's internal token
func (s *server) updateToken() error {
	f, err := os.OpenFile(s.tokenFile, os.O_RDONLY, 0443)
	if err != nil {
		return err
	}

	defer f.Close()
	token, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	s.token = string(token)
	return nil
}
