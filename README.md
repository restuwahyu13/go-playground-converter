## Go Playground Converter

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/restuwahyu13/go-playground-converter?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/restuwahyu13/go-playground-converter)](https://goreportcard.com/report/github.com/restuwahyu13/go-playground-converter)

**go-playground-converter** is formatter error response inspiration like express-validator in nodejs build on top in
go-playground-validator, see more about struct reference follow [this](https://github.com/go-playground/validator) and for new version you can use custom message using `gcp`.

- [Go Playground Converter](#go-playground-converter)
  - [Installation](#installation)
  - [Example Usage Without GPC Tags](#example-usage-without-gpc-tags)
  - [Example Usage With GPC Tags](#example-usage-with-gpc-tags)
- [Testing](#testing)
  - [Bugs](#bugs)
  - [Contributing](#contributing)
  - [License](#license)

### Installation

```sh
$ go get -u github.com/restuwahyu13/go-playground-converter
```

### Example Usage Without GPC Tags

```go
  package main

  import (
  "fmt"
   gpc "github.com/restuwahyu13/go-playground-converter"
  )

  type Login struct {
  	Email    string `validate:"required"`
  	Password string `validate:"required"`
  }

  func main() {
   	  payload := Login{Email: "", Password: ""}
  		res, err := gpc.Validator(payload)

      if err != nil {
        panic(err)
      }

      fmt.Println(res) // if not errors, validator return nil value
  }

  // {
  //   "errors": [
  //     {
  //       "msg": "Email is a required field",
  //       "param": "Email",
  //       "tag": "required"
  //     },
  //     {
  //       "msg": "Password is a required field",
  //       "param": "Password",
  //       "tag": "required"
  //     }
  //   ]
  // }
```

### Example Usage With GPC Tags

```go
  package main

  import (
  "fmt"
   gpc "github.com/restuwahyu13/go-playground-converter"
  )

  type Login struct {
  	Email    string `validate:"required" gpc:"required=Email tidak boleh kosong"`
  	Password string `validate:"required" gpc:"required=Password tidak boleh kosong"`
  }

  func main() {
   	  payload := Login{Email: "", Password: ""}
  		res, err := gpc.Validator(payload)

      if err != nil {
        panic(err)
      }

      fmt.Println(res) // if not errors, validator return nil value
  }

  // {
  //   "errors": [
  //     {
  //       "msg": "Email tidak boleh kosong",
  //       "param": "Email",
  //       "tag": "required"
  //     },
  //     {
  //       "msg": "Password tidak boleh kosong",
  //       "param": "Password",
  //       "tag": "required"
  //     }
  //   ]
  // }
```

## Testing

- Testing Via Local

  ```sh
  go test .
  ```

- Testing Via Docker

  ```sh
  docker build -t go-playground-converter --compress . && docker run go-playground-converter go test .
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
