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

func (s S) Before() {

}

func (s S) After() {

}

func (s S) TestA(t *testing.T) {

}

func (s S) TestB(t *testing.T) {

}

func (s S) TestC(t *testing.T) {

}
