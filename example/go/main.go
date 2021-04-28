package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type User struct {
  Fullname string `validate:"required,lowercase"`
	Email string `validate:"required,email"`
  Password string `validate:"required,gte=8"`
}

func main() {

  var input User
  var validate *validator.Validate
  validators := gpc.NewValidator(validate)
  bind := gpc.NewBindValidator(validators)

  input.Fullname = "Restu Wahyu Saputra"
  input.Email = "restuwahyu13@#zetmail.com"
  input.Password = "qwerty"

  config := gpc.ErrorConfig{
  Options: []gpc.ErrorMetaConfig{
      gpc.ErrorMetaConfig{
        Tag: "lowercase",
        Field: "Fullname",
        Message: "fullname is required",
      },
      gpc.ErrorMetaConfig{
        Tag: "lowercase",
        Field: "Fullname",
        Message: "fullname must be a lowercase",
      },
      gpc.ErrorMetaConfig{
        Tag: "email",
        Field: "Email",
        Message: "email is required",
      },
      gpc.ErrorMetaConfig{
        Tag: "email",
        Field: "Email",
        Message: "email format is not valid",
      },
      gpc.ErrorMetaConfig{
        Tag: "required",
        Field: "Password",
        Message: "password is required",
      },
      gpc.ErrorMetaConfig{
        Tag: "gte",
        Field: "Password",
        Message: "password must be greater 7",
      },
    },
  }
  errResponse := bind.BindValidator(&input, config.Options)
  fmt.Println(errResponse)
}
