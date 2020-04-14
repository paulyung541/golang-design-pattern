package templatemethod

import "testing"

func TestGameFlowA(t *testing.T) {
	var playA Player = NewGameFlowA()
	playA.Play()
}
