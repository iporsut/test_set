package before

import (
	. "github.com/iporsut/test_set"
	"testing"
)

func TestExampleSuite(t *testing.T) {
	RunSuite(S{}, t)
}

type S struct {
}

func (S) Before(t *testing.T) {
	t.Log("Before")
}

func (S) After(t *testing.T) {
	t.Log("After")
}

func (S) TestA(t *testing.T) {
	t.Log("Test A")
}

func (S) TestB(t *testing.T) {
	t.Log("Test B")
}

func (S) TestC(t *testing.T) {
	t.Log("Test C")
}
