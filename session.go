// Inspired by https://github.com/gorilla/sessions.

package session

import (
	"github.com/go-http-utils/cookie-session"
	"github.com/teambition/gear"
)

// CookieSession ...
type CookieSession struct {
	name    string
	newSess func() sessions.Sessions
	store   sessions.Store
}

// New ...
func (cs *CookieSession) New(ctx *gear.Context) (interface{}, error) {
	session := cs.newSess()
	ctx.SetAny(cs, session)
	return session, cs.store.Load(cs.name, session, ctx.Cookies)
}

// New ...
func New(name string, newSess func() sessions.Sessions, options ...*sessions.Options) *CookieSession {
	return &CookieSession{
		name:    name,
		newSess: newSess,
		store:   sessions.New(options...),
	}
}
