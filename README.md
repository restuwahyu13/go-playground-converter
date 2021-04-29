## Go Playground Converter

**go-playground-converter** is formatter error response inspiration like `express-validator` in nodejs build on top in
`go-playground-validator`.

- [GPC API Documentation](<#Agtran-API-Documentation-(Next-Version)>)
  - [Installation](#Installation)
  - [Function Reference](#API-Reference)
    - [NewValidator](#NewValidator)
    - [NewBindValidator](#NewBindValidator)
    - [BindValidator](#BindValidator)
  - [Struct Reference](#Struct-Reference)
    - [ErrorConfig](#ErrorConfig)
    - [ErrorMetaConfig](#ErrorMetaConfig)
  - [Bugs](#Bugs)
  - [Contributing](#Contributing)
  - [License](#License)

### Installation

```sh
$ go get -u https://github.com/restuwahyu13/go-playground-converter
```

### Function Reference

| Method Name          | Description                                                  |
| -------------------- | ------------------------------------------------------------ |
| **NewValidator**     | wrapper struct validation schema for go playground validator |
| **NewBindValidator** | change the struct to an error response                       |
| **BindValidator**    | for returning error response                                 |

### Struct Reference

| Struct Name         | Description                                         |
| ------------------- | --------------------------------------------------- |
| **ErrorConfig**     | config for creating error message                   |
| **ErrorMetaConfig** | meta config defination for returning error response |

### Example Usage

this is an example usage for go-playground-converter see more about
[example usage](https://github.com/restuwahyu13/go-playground-converter/tree/master/example) for go-playground -converter.

```go
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
        ctx.JSON(http.StatusOk, gin.H{"message": "register new account successfully"})
      }
  })

  router.Run(":3000")
}
```

### Bugs

For information on bugs related to package libraries, please visit
[here](https://github.com/restuwahyu13/go-playground-converter/issues)

### Contributing

Want to make **Midtrans Node** more perfect ? Let's contribute and follow the
[contribution guide.](https://github.com/restuwahyu13/go-playground-converter/blob/main/CONTRIBUTING.md)

### License

- [MIT License](https://github.com/restuwahyu13/go-playground-converter/blob/main/LICENSE.md)

<p align="right" style="padding: 5px; border-radius: 100%; background-color: red; font-size: 2rem;">
  <b><a href="#midtrans-node">BACK TO TOP</a></b>
</p>
