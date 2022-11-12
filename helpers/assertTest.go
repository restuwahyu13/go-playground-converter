package helpers

import "testing"

func AssertTest(t *testing.T, val1, val2 interface{}) {
	defer t.Cleanup(func() {
		val1 = nil
		val2 = nil
	})

	if val1 != val2 {
		t.FailNow()
	}
}
