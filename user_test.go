package tester

import "testing"

func TestUserTimeline(t *testing.T) {
	u := NewUser()
	u.Login()
	u.Track()
}
