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
  - [Example Usage](#example-usage)
  - [Bugs](#bugs)
  - [Contributing](#contributing)
  - [License](#license)

### Installation

```sh
$ go get -u github.com/restuwahyu13/go-playground-converter
```

### Example Usage

```go
  package main

  import (
  "fmt"
   gpc "github.com/restuwahyu13/go-playground-converter"
  )

  type Login struct {
  	Email    string `validate:"required,email"`
  	Password string `validate:"required"`
  }

  func main() {
   	payload := Login{Email: "", Password: ""}
  		err := Validator(payload)
  		fmt.Println(err.Errors) // if not errors validator return nil value
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
