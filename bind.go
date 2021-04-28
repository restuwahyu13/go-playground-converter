package gpc

type ErrorResponse struct {
	Results    interface{} `json:"results"`
}

type Binding interface {
	BindValidator(s interface{}) *ErrorResponse
}

type binding struct {
	validators Validators
}

func NewBindValidator(validators Validators) *binding {
	return &binding{validators: validators}
}

func (v *binding) BindValidator(s *interface{}) *ErrorResponse {

	var errors ErrorResponse
  validate := v.validators.Validator(&s)

	errDataCollection := make(map[string][]map[string]interface{})

	for i, v := range validate {
		errorResults := make(map[string]interface{})
		errorResults[i] = v
		errDataCollection["errors"] = append(errDataCollection["errors"], errorResults)
	}

	errors.Results = errDataCollection

	return &errors
}
