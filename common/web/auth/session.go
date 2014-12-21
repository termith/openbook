package auth

import (
	"git.2rll.net/openbook/common/model/user"
	"github.com/gorilla/context"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func Set(u *user.User, rw http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "auth")
	session.Values["id"] = u.Id
	session.Save(req, rw)
}

func Get(req *http.Request) *user.User {
	session, _ := store.Get(req, "auth")
	_, ok := session.Values["id"]
	if _, ok := session.Values["id"]; !ok {
		return nil
	}

	_, ok = context.GetOk(req, "user")
	if !ok {
		context.Set(req, "user", user.Find(session.Values["id"].(int64)))
	}
	return context.Get(req, "user").(*user.User)
}