package stack

import (
	"testing"
)

func TestPushPop(t *testing.T) {
	c := new(stack)
	c.Push(5)
	if c.Pop() != 5 {
		t.Log("Pop not 5")
		t.Fail()
	}
}
