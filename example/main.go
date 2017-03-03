package main

import (
	"fmt"
	"time"

	"github.com/go-http-utils/cookie-session"
	"github.com/teambition/gear"
	"github.com/teambition/gear-session"
)

// Session is a custom session struct.
type Session struct {
	*sessions.Meta `json:"-"`
	UserID         string `json:"_userId"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Avatar         string `json:"avatar"`
	Authed         int64  `json:"authed"`
}

// Save is helpful method
func (s *Session) Save() error {
	return s.GetStore().Save(s)
}

var gearSession = session.New("Sess", sessions.New(), func() sessions.Sessions {
	return &Session{Meta: &sessions.Meta{}}
})

// FromCtx is helpful function to read session from gear.Context
func FromCtx(ctx *gear.Context) (*Session, error) {
	val, err := ctx.Any(gearSession)
	return val.(*Session), err
}

func main() {
	app := gear.New()
	app.Set(gear.SetKeys, []string{"some key"})

	app.Use(func(ctx *gear.Context) error {
		sess, err := FromCtx(ctx)
		if err != nil {
			fmt.Println(sess.IsNew()) // true
			sess.UserID = "xxxxID"
			sess.Name = "gear"
			sess.Email = "gear@teambition.com"
			sess.Authed = time.Now().Unix()
			sess.Save()
		} else {
			fmt.Println(sess.IsNew()) // true
			fmt.Println(sess)

			// update session
			sess.Avatar = "avatar.png"
			sess.Authed = time.Now().Unix()
			sess.Save()
		}
		return ctx.JSON(200, sess)
	})
	app.Listen(":3000")
}
