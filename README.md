# gear-session

[![Build Status](http://img.shields.io/travis/teambition/gear-session.svg?style=flat-square)](https://travis-ci.org/teambition/gear-session)
[![Coverage Status](http://img.shields.io/coveralls/teambition/gear-session.svg?style=flat-square)](https://coveralls.io/r/teambition/gear-session)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/teambition/gear-session/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/teambition/gear-session)

Cookie session middleware for Gear, base on github.com/go-http-utils/cookie-session.

## Demo

```go
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

// Destroy is helpful method
func (s *Session) Destroy() error {
  return s.GetStore().Destroy(s)
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
```
