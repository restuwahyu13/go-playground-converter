package validate

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

var (
	vx *validator.Validate = validator.New()
)

func TestUtilityFunctions(t *testing.T) {
	t.Run("[check] should fail for pointer input", func(t *testing.T) {
		addr := &Address{Street: "123 Main St", City: "TestCity"}
		err := check(addr, vx)
		if err == nil {
			t.Error("Expected error for pointer input, got nil")
		}
	})

	t.Run("[check] should fail for non-struct input", func(t *testing.T) {
		nonStruct := "test string"
		err := check(nonStruct, vx)
		if err == nil {
			t.Error("Expected error for non-struct input, got nil")
		}
	})

	t.Run("[check] should fail for empty struct", func(t *testing.T) {
		emptyStruct := struct{}{}
		err := check(emptyStruct, vx)
		if err == nil {
			t.Error("Expected error for empty struct, got nil")
		}
	})
}
