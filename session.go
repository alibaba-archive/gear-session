package session

import (
	"github.com/go-http-utils/cookie-session"
	"github.com/teambition/gear"
)

// GearSession is a useful wrap of sessions.Store and sessions.Sessions
type GearSession struct {
	name  string
	store sessions.Store
	sess  func() sessions.Sessions
}

// New implements Gear.Any interface
func (gs *GearSession) New(ctx *gear.Context) (interface{}, error) {
	session := gs.sess()
	ctx.SetAny(gs, session)
	return session, gs.store.Load(gs.name, session, ctx.Cookies)
}

// New return a GearSession instance
func New(name string, store sessions.Store, sess func() sessions.Sessions) *GearSession {
	return &GearSession{name, store, sess}
}
