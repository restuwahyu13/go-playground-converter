## Go Playground Converter

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/restuwahyu13/go-playground-converter?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/restuwahyu13/go-playground-converter?style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/restuwahyu13/go-playground-converter?style=for-the-badge)
![GitHub closed issues](https://img.shields.io/github/issues-closed/restuwahyu13/go-playground-converter?style=for-the-badge)
![GitHub contributors](https://img.shields.io/github/contributors/restuwahyu13/go-palyground-converter?style=for-the-badge)

**go-playground-converter** is formatter error response inspiration like express-validator in nodejs build on top in
go-playground-validator, see more about struct reference follow [this](https://github.com/go-playground/validator).

- [Go Playground Converter](#go-playground-converter)
  - [Installation](#installation)
  - [Function Reference](#function-reference)
  - [Struct Reference](#struct-reference)
  - [Example Usage](#example-usage)
  - [Custom Usage](#custom-usage)
  - [Bugs](#bugs)
  - [Contributing](#contributing)
  - [License](#license)

### Installation

```sh
$ go get -u github.com/restuwahyu13/go-playground-converter
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
[example usage](https://github.com/restuwahyu13/go-playground-converter/tree/master/example) for go-playground-converter.

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
  // error response like this
  //   {
  //     "results": {
  //         "errors": [
  //             {
  //                 "Fullname": {
  //                     "message": "fullname must be a lowercase",
  //                     "value": "Restu Wahyu Saputra",
  //                     "param": "Fullname",
  //                     "tag": "lowercase"
  //                 }
  //             },
  //             {
  //                 "Email": {
  //                     "message": "email format is not valid",
  //                     "value": "restuwahyu13@#gmail.com",
  //                     "param": "Email",
  //                     "tag": "email"
  //                 }
  //             },
  //             {
  //                 "Password": {
  //                     "message": "password must be greater 7",
  //                     "value": "qwerty",
  //                     "param": "Password",
  //                     "tag": "gte"
  //                 }
  //             }
  //         ]
  //     }
  // }
```

### Custom Usage

```go
package util

import (
  "github.com/go-playground/validator/v10"
  gpc "github.com/restuwahyu13/go-playground-converter"
)


func GoValidator(s interface{}, config []gpc.ErrorMetaConfig) (interface{}, int) {
    var validate *validator.Validate
    validators := gpc.NewValidator(validate)
    bind := gpc.NewBindValidator(validators)

    errResponse, errCount := bind.BindValidator(s, config)
    return errResponse, errCount
}
```

### Bugs

For information on bugs related to package libraries, please visit
[here](https://github.com/restuwahyu13/go-playground-converter/issues)

### Contributing

Want to make **Go Playground Converter** more perfect ? Let's contribute and follow the
[contribution guide.](https://github.com/restuwahyu13/go-playground-converter/blob/master/CONTRIBUTING.md)

### License

- [MIT License](https://github.com/restuwahyu13/go-playground-converter/blob/master/LICENSE.md)

<p align="right" style="padding: 5px; border-radius: 100%; background-color: red; font-size: 2rem;">
  <b><a href="#go-playground-converter">BACK TO TOP</a></b>
</p>
