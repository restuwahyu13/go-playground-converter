package helper

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

func Exception(key string, format any) error {
	msg := sync.Map{}

	if format != nil {
		msg.Store("key_exist_input_invalid", fmt.Errorf("validator value not supported, because %v is not struct", reflect.TypeOf(format).Kind().String()))
		msg.Store("err_format_input_not_pointer", fmt.Errorf("validator value not supported, because %v is struct pointer", format))
		msg.Store("err_format_input_not_struct", fmt.Errorf("validator value not supported, because %v is not struct", format))
		msg.Store("err_format_input_empty_struct", fmt.Errorf("validator value can't be empty struct %v", format))
	} else {
		msg.Store("err_format_not_err_validator", errors.New("err is not validator error"))
		msg.Store("err_format_transalator_not_found", errors.New("translator not found"))
	}

	result, ok := msg.LoadAndDelete(key)
	if ok {
		return result.(error)
	}

	return nil
}
