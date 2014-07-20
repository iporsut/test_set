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

func (S) Before() {

}

func (S) After() {

}

func (S) TestA(t *testing.T) {

}

func (S) TestB(t *testing.T) {

}

func (S) TestC(t *testing.T) {

}
