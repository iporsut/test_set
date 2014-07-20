package test_set

import (
	"reflect"
	"strings"
	"testing"
)

func RunSuite(s interface{}, t *testing.T) {
	st := reflect.TypeOf(s)
	sv := reflect.ValueOf(s)
	tv := reflect.ValueOf(t)
	args := make([]reflect.Value, 2)
	args[0] = sv
	args[1] = tv
	numOfTest := st.NumMethod()
	for i := 0; i < numOfTest; i++ {
		if strings.HasPrefix(st.Method(i).Name, "Test") {
			if before, ok := st.MethodByName("Before"); ok {
				before.Func.Call(args)
			}
			st.Method(i).Func.Call(args)
			if after, ok := st.MethodByName("After"); ok {
				after.Func.Call(args)
			}
		}
	}
}
