package validate

import (
	"context"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type (
	Validate interface {
		RegisterAlias(alias string, tags string)
		RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...interface{})
		RegisterStructValidation(fn validator.StructLevelFunc, types ...interface{})
		RegisterStructValidationCtx(fn validator.StructLevelFuncCtx, types ...interface{})
		RegisterStructValidationMapRules(rules map[string]string, types ...interface{})
		RegisterTagNameFunc(fn validator.TagNameFunc)
		RegisterTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc, translationFn validator.TranslationFunc) (err error)
		RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error
		RegisterValidationCtx(tag string, fn validator.FuncCtx, callValidationEvenIfNull ...bool) error
		Struct(s interface{}) ([]Errors, error)
		StructCtx(ctx context.Context, s interface{}) ([]Errors, error)
		StructExcept(s interface{}, fields ...string) ([]Errors, error)
		StructExceptCtx(ctx context.Context, s interface{}, fields ...string) ([]Errors, error)
		StructFiltered(s interface{}, fn validator.FilterFunc) ([]Errors, error)
		StructFilteredCtx(ctx context.Context, s interface{}, fn validator.FilterFunc) ([]Errors, error)
		StructPartial(s interface{}, fields ...string) ([]Errors, error)
		StructPartialCtx(ctx context.Context, s interface{}, fields ...string) ([]Errors, error)
		Var(field interface{}, tag string) error
		VarCtx(ctx context.Context, field interface{}, tag string) error
		VarWithValueCtx(ctx context.Context, field interface{}, other interface{}, tag string) error
		ValidateMap(data map[string]interface{}, rules map[string]interface{}) map[string]interface{}
		ValidateMapCtx(ctx context.Context, data map[string]interface{}, rules map[string]interface{}) map[string]interface{}
		SetTagName(name string)
	}

	validate struct {
		validator *validator.Validate
	}
)

func New(options ...validator.Option) Validate {
	return validate{validator: validator.New(options...)}
}

/*
RegisterAlias registers a mapping of a single validation tag that defines a common or complex set of validation(s) to simplify adding validation to structs.

NOTE: this function is not thread-safe it is intended that these all be registered prior to any validation
*/
func (v validate) RegisterAlias(alias string, tags string) {
	v.validator.RegisterAlias(alias, tags)
}

/*
RegisterCustomTypeFunc registers a CustomTypeFunc against a number of types

NOTE: this method is not thread-safe it is intended that these all be registered prior to any validation
*/
func (v validate) RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...interface{}) {
	v.validator.RegisterCustomTypeFunc(fn, types...)
}

/*
RegisterStructValidation registers a StructLevelFunc against a number of types.

NOTE: - this method is not thread-safe it is intended that these all be registered prior to any validation
*/
func (v validate) RegisterStructValidation(fn validator.StructLevelFunc, types ...interface{}) {
	v.validator.RegisterStructValidation(fn, types...)
}

/*
RegisterStructValidationCtx registers a StructLevelFuncCtx against a number of types and allows passing of contextual validation information via context.Context.

NOTE: - this method is not thread-safe it is intended that these all be registered prior to any validation
*/
func (v validate) RegisterStructValidationCtx(fn validator.StructLevelFuncCtx, types ...interface{}) {
	v.validator.RegisterStructValidationCtx(fn, types...)
}

/*
RegisterStructValidationMapRules registers validate map rules. Be aware that map validation rules supersede those defined on a/the struct if present.

NOTE: this method is not thread-safe it is intended that these all be registered prior to any validation
*/
func (v validate) RegisterStructValidationMapRules(rules map[string]string, types ...interface{}) {
	v.validator.RegisterStructValidationMapRules(rules, types...)
}

/*
RegisterTagNameFunc registers a function to get alternate names for StructFields.

eg. to use the names which have been specified for JSON representations of structs, rather than normal Go field names:

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
	    name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	    // skip if tag key says it should be ignored
	    if name == "-" {
	        return ""
	    }
	    return name
	})
*/
func (v validate) RegisterTagNameFunc(fn validator.TagNameFunc) {
	v.validator.RegisterTagNameFunc(fn)
}

/*
RegisterTranslation registers translations against the provided tag.
*/
func (v validate) RegisterTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc, translationFn validator.TranslationFunc) (err error) {
	return v.validator.RegisterTranslation(tag, trans, registerFn, translationFn)
}

/*
RegisterValidation adds a validation with the given tag

NOTES: - if the key already exists, the previous validation function will be replaced. - this method is not thread-safe it is intended that these all be registered prior to any validation
*/
func (v validate) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return v.validator.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

/*
RegisterValidationCtx does the same as RegisterValidation on accepts a FuncCtx validation allowing context.Context validation support.
*/
func (v validate) RegisterValidationCtx(tag string, fn validator.FuncCtx, callValidationEvenIfNull ...bool) error {
	return v.validator.RegisterValidationCtx(tag, fn, callValidationEvenIfNull...)
}

/*
RegisterValidate, this is custom validaton without context and you not manual configure the validator.
*/
// func (v validate) RegisterValidate(s interface{}) ([]Errors, error) {
// 	return handlerValidator(nil, s, v.validator)
// }

// /*
// RegisterValidateCtx does the same as RegisterValidate, this is custom validaton with context and you not manual configure the validator.
// */
// func (v validate) RegisterValidateCtx(ctx context.Context, s interface{}) ([]Errors, error) {
// 	return handlerValidator(ctx, s, v.validator)
// }

/*
Struct validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors.
*/
func (v validate) Struct(s interface{}) ([]Errors, error) {
	if err := check(s, v.validator); err != nil {
		return nil, err
	}

	err := v.validator.Struct(s)
	if err == nil {
		return nil, nil
	}

	return translator(err, s, v.validator)
}

/*
StructCtx validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified and also allows passing of context.Context for contextual validation information.

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors.
*/
func (v validate) StructCtx(ctx context.Context, s interface{}) ([]Errors, error) {
	if err := check(s, v.validator); err != nil {
		return nil, err
	}

	err := v.validator.StructCtx(ctx, s)
	if err == nil {
		return nil, nil
	}

	return translator(err, s, v.validator)
}

/*
StructExcept validates all fields except the ones passed in. Fields may be provided in a namespaced fashion relative to the struct provided i.e. NestedStruct.Field or NestedArrayField[0].Struct.Name

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors.
*/
func (v validate) StructExcept(s interface{}, fields ...string) ([]Errors, error) {
	if err := check(s, v.validator); err != nil {
		return nil, err
	}

	err := v.validator.StructExcept(s, fields...)
	if err == nil {
		return nil, nil
	}

	return translator(err, s, v.validator)
}

/*
StructExceptCtx validates all fields except the ones passed in and allows passing of contextual validation information via context.Context Fields may be provided in a namespaced fashion relative to the struct provided i.e. NestedStruct.Field or NestedArrayField[0].Struct.Name

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors.
*/
func (v validate) StructExceptCtx(ctx context.Context, s interface{}, fields ...string) ([]Errors, error) {
	if err := check(s, v.validator); err != nil {
		return nil, err
	}

	err := v.validator.StructExceptCtx(ctx, s, fields...)
	if err == nil {
		return nil, nil
	}

	return translator(err, s, v.validator)
}

/*
StructFiltered validates a structs exposed fields, that pass the FilterFunc check and automatically validates nested structs, unless otherwise specified.

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors.
*/
func (v validate) StructFiltered(s interface{}, fn validator.FilterFunc) ([]Errors, error) {
	if err := check(s, v.validator); err != nil {
		return nil, err
	}

	err := v.validator.StructFiltered(s, fn)
	if err == nil {
		return nil, nil
	}

	return translator(err, s, v.validator)
}

/*
StructFilteredCtx validates a structs exposed fields, that pass the FilterFunc check and automatically validates nested structs, unless otherwise specified and also allows passing of contextual validation information via context.Context

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors.
*/
func (v validate) StructFilteredCtx(ctx context.Context, s interface{}, fn validator.FilterFunc) ([]Errors, error) {
	if err := check(s, v.validator); err != nil {
		return nil, err
	}

	err := v.validator.StructFilteredCtx(ctx, s, fn)
	if err == nil {
		return nil, nil
	}

	return translator(err, s, v.validator)
}

/*
StructPartial validates the fields passed in only, ignoring all others. Fields may be provided in a namespaced fashion relative to the struct provided eg. NestedStruct.Field or NestedArrayField[0].Struct.Name

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors.
*/
func (v validate) StructPartial(s interface{}, fields ...string) ([]Errors, error) {
	if err := check(s, v.validator); err != nil {
		return nil, err
	}

	err := v.validator.StructPartial(s, fields...)
	if err == nil {
		return nil, nil
	}

	return translator(err, s, v.validator)
}

/*
StructPartialCtx validates the fields passed in only, ignoring all others and allows passing of contextual validation information via context.Context Fields may be provided in a namespaced fashion relative to the struct provided eg. NestedStruct.Field or NestedArrayField[0].Struct.Name

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors.
*/
func (v validate) StructPartialCtx(ctx context.Context, s interface{}, fields ...string) ([]Errors, error) {
	if err := check(s, v.validator); err != nil {
		return nil, err
	}

	err := v.validator.StructPartialCtx(ctx, s, fields...)
	if err == nil {
		return nil, nil
	}

	return translator(err, s, v.validator)
}

/*
Var validates a single variable using tag style validation. eg. var i int validate.Var(i, "gt=1,lt=10")

WARNING: a struct can be passed for validation eg. time.Time is a struct or if you have a custom type and have registered a custom type handler, so must allow it; however unforeseen validations will occur if trying to validate a struct that is meant to be passed to 'validate.Struct'

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors. validate Array, Slice and maps fields which may contain more than one error
*/
func (v validate) Var(field interface{}, tag string) error {
	return v.validator.Var(field, tag)
}

/*
VarCtx validates a single variable using tag style validation and allows passing of contextual validation information via context.Context. eg. var i int validate.Var(i, "gt=1,lt=10")

WARNING: a struct can be passed for validation eg. time.Time is a struct or if you have a custom type and have registered a custom type handler, so must allow it; however unforeseen validations will occur if trying to validate a struct that is meant to be passed to 'validate.Struct'

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors. validate Array, Slice and maps fields which may contain more than one error
*/
func (v validate) VarCtx(ctx context.Context, field interface{}, tag string) (err error) {
	return v.validator.VarCtx(ctx, field, tag)
}

/*
VarWithValueCtx validates a single variable, against another variable/field's value using tag style validation and allows passing of contextual validation validation information via context.Context. eg. s1 := "abcd" s2 := "abcd" validate.VarWithValue(s1, s2, "eqcsfield") // returns true

WARNING: a struct can be passed for validation eg. time.Time is a struct or if you have a custom type and have registered a custom type handler, so must allow it; however unforeseen validations will occur if trying to validate a struct that is meant to be passed to 'validate.Struct'

It returns InvalidValidationError for bad values passed in and nil or ValidationErrors as error otherwise. You will need to assert the error if it's not nil eg. err.(validator.ValidationErrors) to access the array of errors. validate Array, Slice and maps fields which may contain more than one error
*/
func (v validate) VarWithValueCtx(ctx context.Context, field interface{}, other interface{}, tag string) (err error) {
	return v.validator.VarWithValueCtx(ctx, field, other, tag)
}

// SetTagName allows for changing of the default tag name of 'validate'
func (v validate) SetTagName(name string) {
	v.validator.SetTagName(name)
}

// ValidateMap validates map data from a map of tags
func (v validate) ValidateMap(data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	return v.validator.ValidateMap(data, rules)
}

// ValidateMapCtx validates a map using a map of validation rules and allows passing of contextual validation information via context.Context.
func (v validate) ValidateMapCtx(ctx context.Context, data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	return v.validator.ValidateMapCtx(ctx, data, rules)
}
