## validation

```go
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
    validation.Field(&c.Address, rules.Required),
    validation.Field(&c.Address.City, rules.Required, rules.Length(20, 100)),
    validation.Field(&c.Address.Street, rules.Required, rules.RuneLength(20, 1000)),
)

err := valid.Validate()
```

```go
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
    validation.Field("address", rules.Required),
    validation.Field("address.city", rules.Required, rules.Length(20, 100)),
)

err := valid.Validate()
```