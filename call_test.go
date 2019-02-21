package tester

import (
	"testing"
)

func TestTrack(t *testing.T) {
	track := NewTrack()
	j := track.JSON()
	if j == "" {
		t.Error("error converting track to json", j)
	}
}
