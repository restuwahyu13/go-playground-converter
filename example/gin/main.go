package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type User struct {
  Fullname string `validate:"required,lowercase"`
	Email string `validate:"required,email"`
  Password string `validate:"required,gte=8"`
}

func main() {
  router := gin.Default()
  router.POST("/", func(ctx *gin.Context) {

    var input User
    ctx.ShouldBindJSON(&input)

    var validate *validator.Validate
    validators := gpc.NewValidator(validate)
    bind := gpc.NewBindValidator(validators)

    config := gpc.ErrorConfig{
      Options: []gpc.ErrorMetaConfig{
          gpc.ErrorMetaConfig{
            Tag: "required",
            Field: "Fullname",
            Message: "fullname is required",
          },
          gpc.ErrorMetaConfig{
            Tag: "lowercase",
            Field: "Fullname",
            Message: "fullname must be a lowercase",
          },
          gpc.ErrorMetaConfig{
            Tag: "required",
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

      errResponse, errCount := bind.BindValidator(&input, config.Options)

      if errCount > 0 {
        ctx.JSON(http.StatusBadRequest, errResponse)
        ctx.AbortWithStatus(http.StatusBadRequest)
        return
      } else {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "register new account successfully"})
      }
  })

  router.Run(":3000")
}
