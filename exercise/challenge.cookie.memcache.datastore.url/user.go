package mem

import (
	"errors"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

func newVisitor(req *http.Request) (*http.Cookie, error) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Println("ERROR newVisitor uuid.NewV4", err)
		return nil, err
	}
	m := initialModel(id.String())
	return makeCookie(m, req)
}

func currentVisitor(m model, req *http.Request) (*http.Cookie, error) {
	return makeCookie(m, req)
}

func initialModel(id string) model {
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"one.jpg",
		},
		ID: id,
	}
	return m
}

// Gets user's Id from cookie, if it does not exists it will look into
// url and set it on the cookie and then returns it.
func getId(res http.ResponseWriter, req *http.Request) (string, error) {

	// Looking into cookie
	log.Println("Looking into cookie for session-id...")
	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie {
		// Looking into URL
		log.Println("Looking into URL...")
		id := req.FormValue("id")
		if id == "" {
			return id, errors.New("[id] not found on URL")
		}

		// Setting user's cookie
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    id,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	return cookie.Value, nil
}
