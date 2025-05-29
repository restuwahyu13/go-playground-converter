package helper

import (
	"testing"
)

func TestMergeSlice(action *testing.T) {
	action.Run("Should be TestMergeSlice - is empty response", func(t *testing.T) {
		var (
			data []string = []string{}
			res  []string = mergeSlice(data)
		)

		if len(res) != 0 {
			t.FailNow()
		}
	})

	action.Run("Should be TestMergeSlice - not empty response", func(t *testing.T) {
		data := []string{"a", "b", "c"}

		res := mergeSlice(data)
		if len(res) == 0 {
			t.FailNow()
		}
	})
}
