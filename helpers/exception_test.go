package helper

import "testing"

func TestExceptionForValidator(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		format any
		want   string
	}{
		{
			name:   "invalid input - not struct",
			key:    "key_exist_input_invalid",
			format: "string input",
			want:   "validator value not supported, because string is not struct",
		},
		{
			name:   "invalid input - struct pointer",
			key:    "err_format_input_not_pointer",
			format: &struct{}{},
			want:   "validator value not supported, because &{} is struct pointer",
		},
		{
			name:   "invalid input - not struct type",
			key:    "err_format_input_not_struct",
			format: []string{"test"},
			want:   "validator value not supported, because []string is not struct",
		},
		{
			name:   "empty struct input",
			key:    "err_format_input_empty_struct",
			format: struct{}{},
			want:   "validator value can't be empty struct {}",
		},
		{
			name:   "not validator error",
			key:    "err_format_not_err_validator",
			format: nil,
			want:   "err is not validator error",
		},
		{
			name:   "translator not found",
			key:    "err_format_transalator_not_found",
			format: nil,
			want:   "translator not found",
		},
		{
			name:   "non-existent error key with format",
			key:    "non_existent_key",
			format: struct{ Name string }{Name: "test"},
			want:   "",
		},
		{
			name:   "non-existent error key without format",
			key:    "non_existent_key",
			format: nil,
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Exception(tt.key, tt.format)
			if got == nil && tt.want != "" {
				t.Errorf("Exception() = nil, want %v", tt.want)
			}
			if got != nil && got.Error() != tt.want {
				t.Errorf("Exception() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}
