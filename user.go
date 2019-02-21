package tester

import (
	"github.com/yields/phony/pkg/phony"
)

type User struct {
	UserId      string
	AnonymousId string
	Traits      map[string]interface{}
	Device      string
	Timeline    []Call
}

func NewUser() *User {
	u := User{}
	u.NewSession()
	return &u
}

func (u *User) NewSession() {
	u.AnonymousId = phony.Get("id")
	u.Traits = make(map[string]interface{})
}

func (u *User) Login() {
	if u.UserId == "" {
		u.UserId = phony.Get("id")
	}

	u.Traits["email"] = phony.Get("email")
}

func (u *User) Track() {
	t := NewTrack()
	u.addCommonFields(&t.Message)
	t.Event = "foo"
}

func (u *User) addCommonFields(m *Message) {
	m.AnonymousID = u.AnonymousId
	m.UserID = u.UserId
}
