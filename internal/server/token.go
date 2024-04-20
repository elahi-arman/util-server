package server

import (
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

const (
	token_cookie = "vt"
)

func (s *server) VerifyToken(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookie, err := r.Cookie(token_cookie)

		// would only return ErrNoCookie according to documentation so
		// no cookie would mean the request is unauthenticated
		if err != nil {
			w.WriteHeader(401)
			return
		}

		if cookie.Valid() != nil || cookie.Value != s.token {
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
	f, err := os.OpenFile(s.healthFile, os.O_RDONLY, 0443)
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
