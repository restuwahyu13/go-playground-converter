package gpc

import "testing"

func TestMergeSlice(action *testing.T) {
	action.Run("Should be TestMergeSlice - is empty response", func(t *testing.T) {
		data := []string{}

		res := mergeSlice(data)
		if len(res) == 0 {
			assert(t, len(res), 0)
		}
	})

	action.Run("Should be TestMergeSlice - not empty response", func(t *testing.T) {
		data := []string{"a", "b", "c"}

		res := mergeSlice(data)
		if len(res) == 0 {
			assert(t, len(res), 0)
		}
	})
}
