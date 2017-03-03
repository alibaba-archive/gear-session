package session

import (
	"testing"
	"time"

	"github.com/DavidCai1993/request"
	"github.com/go-http-utils/cookie-session"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
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

var gearSession = New("GSess", sessions.New(), func() sessions.Sessions {
	return &Session{Meta: &sessions.Meta{}}
})

// FromCtx is helpful function to read session from gear.Context
func FromCtx(ctx *gear.Context) (*Session, error) {
	val, err := ctx.Any(gearSession)
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
				sess.UserID = "xxxxID"
				sess.Name = "gear"
				sess.Email = "gear@teambition.com"
				sess.Authed = time.Now().Unix()
				sess.Save()
			} else {
				assert.False(sess.IsNew())
				assert.Equal("xxxxID", sess.UserID)
				assert.Equal("gear", sess.Name)
				assert.Equal("gear@teambition.com", sess.Email)
				assert.Equal("", sess.Avatar)
				assert.True(sess.Authed > 0)

				sess.Avatar = "avatar.png"
				sess.Save()
			}
			return ctx.JSON(200, sess)
		})
		srv := app.Start()
		defer srv.Close()
		url := "http://" + srv.Addr().String()

		res, err := request.Get(url).End()
		assert.Nil(err)

		data, _ := res.JSON()
		s := *(data.(*map[string]interface{}))
		assert.Nil(err)
		assert.Equal("xxxxID", s["_userId"].(string))
		assert.Equal("gear", s["name"].(string))
		assert.Equal("gear@teambition.com", s["email"].(string))
		assert.Equal("", s["avatar"].(string))
		assert.True(s["authed"].(float64) > 0)

		cli := request.Get(url)
		for _, c := range res.Cookies() {
			cli.Cookie(c)
		}
		res, err = cli.End()
		assert.Nil(err)

		data, _ = res.JSON()
		s = *(data.(*map[string]interface{}))
		assert.Equal("xxxxID", s["_userId"].(string))
		assert.Equal("gear", s["name"].(string))
		assert.Equal("gear@teambition.com", s["email"].(string))
		assert.Equal("avatar.png", s["avatar"].(string))
		assert.True(s["authed"].(float64) > 0)
	})
}
