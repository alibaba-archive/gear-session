package session

import (
	"fmt"
	"testing"
	"time"

	"github.com/DavidCai1993/request"
	"github.com/go-http-utils/cookie-session"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

// Session ...
type Session struct {
	*sessions.Meta `json:"-"`
	Name           string `json:"name"`
	Age            int64  `json:"age"`
	Authed         int64  `json:"authed"`
}

// Save ...
func (s *Session) Save() error {
	return s.GetStore().Save(s)
}

var cookieSession = New("GSess", func() sessions.Sessions {
	return &Session{Meta: &sessions.Meta{}}
})

func FromCtx(ctx *gear.Context) (*Session, error) {
	val, err := ctx.Any(cookieSession)
	return val.(*Session), err
}

func TestGearSession(t *testing.T) {
	t.Run("should work", func(t *testing.T) {
		assert := assert.New(t)

		app := gear.New()
		app.Set(gear.SetKeys, []string{"some key"})
		app.Use(func(ctx *gear.Context) error {
			sess, err := FromCtx(ctx)
			if err != nil {
				assert.True(sess.IsNew())
				sess.Name = "gear"
				sess.Age = 18
				sess.Authed = time.Now().Unix()
				sess.Save()
			} else {
				assert.False(sess.IsNew())
				assert.Equal("gear", sess.Name)
				assert.Equal(18, sess.Age)
				assert.True(sess.Authed > 0)
			}
			return ctx.JSON(200, sess)
		})
		srv := app.Start()
		defer srv.Close()
		url := "http://" + srv.Addr().String()

		res, err := request.Get(url).JSON()
		fmt.Println(11111, res, err)
		assert.Nil(err)
	})
}
