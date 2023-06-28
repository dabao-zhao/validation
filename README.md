## validation

```go
c := map[string]interface{}{
	"Name":  "Qiang Xue",
	"Email": "q",
	"Address": map[string]interface{}{
		"Street": "123",
		"City":   "Unknown",
		"State":  "Virginia",
		"Zip":    "12345",
	},
}

err := validation.Validate(c,
	validation.Map(
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Key("Name", validation.Required.Error("{{.field}} is required"), validation.Length(5, 20)),
		// Email cannot be empty and should be in a valid email format.
		validation.Key("Email", validation.Required, is.Email),
		// Validate Address using its own validation rules
		validation.Key("Address", validation.Map(
			// Street cannot be empty, and the length must between 5 and 50
			validation.Key("Street", validation.Required, validation.Length(5, 50)),
			// City cannot be empty, and the length must between 5 and 50
			validation.Key("City", validation.Required, validation.Length(5, 50)),
			// State cannot be empty, and must be a string consisting of two letters in upper case
			validation.Key("State", validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
			// State cannot be empty, and must be a string consisting of five digits
			validation.Key("Zip", validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		)),
	),
)
```