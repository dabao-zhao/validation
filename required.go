package validation

var (
	ErrRequired = NewError("the {{.field}} field is required.")
)

type RequiredRule struct {
	err Error
}

// Required 验证非空.
var Required = RequiredRule{
	err: ErrRequired,
}

func (r RequiredRule) Validate(key, value interface{}) error {
	v, isNil := Indirect(value)
	if isNil || IsEmpty(v) {
		return r.err.Parse(map[string]interface{}{"field": key})
	}
	return nil
}

func (r RequiredRule) Error(message string) RequiredRule {
	if r.err == nil {
		r.err = ErrRequired
	}
	r.err = r.err.SetMessage(message)
	return r
}
