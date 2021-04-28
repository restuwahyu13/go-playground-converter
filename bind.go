package gpc

type ErrorResponse struct {
	Results    interface{} `json:"results"`
}

type Binding interface {
	BindValidator(s interface{}) *ErrorResponse
}

type binding struct {
	validator Validator
}

func NewBindValidator(validator Validator) *binding {
	return &binding{validator: validator}
}

func (v *binding) BindValidator(s interface{}) *ErrorResponse {

	var errors ErrorResponse
  validate := v.validator.Validator(s)

	errDataCollection := make(map[string][]map[string]interface{})

	for i, v := range validate {
		errorResults := make(map[string]interface{})
		errorResults[i] = v
		errDataCollection["errors"] = append(errDataCollection["errors"], errorResults)
	}

	errors.Results = errDataCollection

	return &errors
}
