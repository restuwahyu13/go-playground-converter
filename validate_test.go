package validate

import (
	"context"
	"testing"
)

var (
	v Validate = New()
)

type (
	Address struct {
		Street  string   `json:"street" validate:"required,min=3,max=50"`
		City    string   `json:"city" validate:"required,min=2,max=30"`
		Country string   `json:"country" validate:"required,oneof=ID US UK JP SG"`
		Status  []string `json:"status" validate:"required,gt=0,dive,oneof=active inactive"`
	}

	User struct {
		Name      string    `json:"name" validate:"required,min=3,max=50"`
		Email     string    `json:"email" validate:"required,email"`
		Age       int       `json:"age" validate:"gte=18,lte=100"`
		Addresses []Address `json:"addresses" validate:"dive"`
	}

	User2 struct {
		Name      string    `json:"name" validate:"required,min=3,max=50"`
		Email     string    `json:"email" validate:"required,email"`
		Age       int       `json:"age" validate:"gte=18,lte=100"`
		Addresses []Address `json:"addresses" validate:"dive"`
	}
)

var (
	rules map[string]interface{} = map[string]interface{}{"Street": "required,min=3,max=50", "City": "required,min=2,max=30"}
)

func TestRegisterValidate(t *testing.T) {

	t.Run("[Struct] should be register standart validate is success", func(t *testing.T) {
		body := Address{Street: "", City: "", Country: "", Status: []string{"x"}}

		res, err := v.Struct(body)
		t.Logf("should be register standart validate is success, Error: %v", err)
		t.Log(res)

		if err != nil {
			t.FailNow()
		}

		if len(res) < 1 {
			t.FailNow()
		}
	})

	t.Run("[Struct] should be register nested validate is success", func(t *testing.T) {
		body := User{
			Name:      "Max Cavalera",
			Email:     "maxcavalera13@gmail.com",
			Age:       30,
			Addresses: []Address{{Street: "", City: "", Country: ""}},
		}

		res, err := v.Struct(body)
		t.Logf("should be register nested validate is success, Error: %v", err)
		t.Log(len(res))

		if err != nil {
			t.FailNow()
		}

		if len(res) < 1 {
			t.FailNow()
		}
	})

	t.Run("[StructCtx] should be register standart validate is success", func(t *testing.T) {
		body := Address{Street: "", City: "", Country: ""}
		ctx := context.Background()

		res, err := v.StructCtx(ctx, body)

		t.Logf("should be register standart validate is success, Error: %v", err)
		t.Log(res)

		if err != nil {
			t.FailNow()
		}

		if len(res) < 1 {
			t.FailNow()
		}
	})

	t.Run("[StructCtx] should be register nested validate is success", func(t *testing.T) {
		body := User{
			Name:      "Max Cavalera",
			Email:     "maxcavalera13@gmail.com",
			Age:       30,
			Addresses: []Address{{Street: "", City: "", Country: ""}},
		}

		ctx := context.Background()
		res, err := v.StructCtx(ctx, body)

		t.Logf("should be register nested validate is success, Error: %v", err)
		t.Log(len(res))

		if err != nil {
			t.FailNow()
		}

		if len(res) < 1 {
			t.FailNow()
		}
	})

	t.Run("[StructExcept] should be register standart validate is success", func(t *testing.T) {
		body := Address{Street: "", City: "", Country: ""}

		res, err := v.StructExcept(body, "Street", "Country")

		t.Logf("should be register standart validate is success, Error: %v", err)
		t.Log(res)

		if err != nil {
			t.FailNow()
		}

		if len(res) < 1 {
			t.FailNow()
		}
	})

	t.Run("[StructExceptCtx] should be register nested standart validate is success", func(t *testing.T) {
		body := Address{Street: "", City: "", Country: ""}
		ctx := context.Background()

		res, err := v.StructExceptCtx(ctx, body, "Street", "Country")

		t.Logf("should be register nested validate is success, Error: %v", err)
		t.Log(res)

		if err != nil {
			t.FailNow()
		}

		if len(res) < 1 {
			t.FailNow()
		}
	})

	t.Run("[StructPartial] should be register standart validate is success", func(t *testing.T) {
		body := Address{Street: "", City: "", Country: ""}
		res, err := v.StructPartial(body, "Street", "Country")

		t.Logf("should be register standart validate is success, Error: %v", err)
		t.Log(res)

		if err != nil {
			t.FailNow()
		}

		if len(res) < 1 {
			t.FailNow()
		}
	})

	t.Run("[StructPartialCtx] should be register standart validate is success", func(t *testing.T) {
		body := Address{Street: "", City: "", Country: ""}
		ctx := context.Background()

		res, err := v.StructPartialCtx(ctx, body, "Street", "Country")

		t.Logf("should be register nested validate is success, Error: %v", err)
		t.Log(res)

		if err != nil {
			t.FailNow()
		}

		if len(res) < 1 {
			t.FailNow()
		}
	})
}
