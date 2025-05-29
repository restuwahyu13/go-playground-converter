## Go Playground Converter

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/restuwahyu13/go-playground-converter?style=flat)
[![Go Report Card](https://goreportcard.com/badge/github.com/restuwahyu13/go-playground-converter)](https://goreportcard.com/report/github.com/restuwahyu13/go-playground-converter) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/restuwahyu13/go-playground-converter/blob/master/CONTRIBUTING.md)

**go-playground-converter** is formatter error response inspiration like express-validator in Node.js. `go-playground-converter` builds on top of
`go-playground/validator`, see more about struct references and follow this [documentation](https://github.com/go-playground/validator), And for the new version of `go-playground-converter`, you can use custom messages using `gpc` struct tags, and you need the core `go-playground/validator` you can access use  `GoValidator`.

- [Go Playground Converter](#go-playground-converter)
  - [Installation](#installation)
  - [Example Usage](#example-usage)
- [Testing](#testing)
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
   validate "github.com/restuwahyu13/go-playground-converter"
  )

  type Login struct {
  	Email    string `validate:"required" json:"email"`
  	Password string `validate:"required" json:"password"`
  }

  func main() {
  	payload := Login{Email: "", Password: ""}
  	validator := validate.New() // if not errors, validator return res & err nil value

  	res, err := validate.Struct(req)
  	if err != nil {
  		panic(err)
  	}

  	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  		w.Header().Set("Content-Type", "application/json")
  		json.NewEncoder(w).Encode(&res)
  	})

  	http.ListenAndServe(":3000", nil)
  }

  //  [
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
```

## Testing

- Testing Via Local

  ```sh
   go test --race -v --failfast .
  ```

- Testing Via Docker

  ```sh
  docker build -t go-playground-converter --compress . && docker run go-playground-converter go test --race -v --failfast .
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
