package validation_test

import (
	"fmt"

	"github.com/dabao-zhao/validation"
)

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Customer struct {
	Name    string
	Gender  string
	Email   string
	Address Address `json:"address"`
}

func ExampleValidation_ValidateStruct() {
	c := Customer{
		Name:  "dabao",
		Email: "977904037@qq.com",
		Address: Address{
			Street: "5",
			City:   "Beijing",
			State:  "Beijing",
			Zip:    "0000000",
		},
	}
	valid := validation.Make(&c,
		validation.Field(&c.Address, validation.Required),
		validation.Field(&c.Address.City, validation.Required, validation.Length(20, 100)),
		validation.Field(&c.Address.Street, validation.Required, validation.RuneLength(20, 1000)),
	)

	err := valid.Validate()
	fmt.Println(err)

	// Output:
	// ["the address.City field length must be between 20 and 100","the address.Street field length must be between 20 and 1000"]
}

func ExampleValidation_ValidateMap() {
	c := map[string]interface{}{
		"name":  "dabao",
		"email": "977904037@qq.com",
		"address": map[string]interface{}{
			"street": "5",
			"city":   "Beijing",
			"state":  "Beijing",
			"zip":    "0000000",
		},
	}
	valid := validation.Make(&c,
		validation.Field("address", validation.Required),
		validation.Field("address.city", validation.Required, validation.Length(20, 100)),
	)

	err := valid.Validate()
	fmt.Println(err)

	// Output:
	// ["the address.city field length must be between 20 and 100"]
}
