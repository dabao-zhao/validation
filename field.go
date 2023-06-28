package validation

type (
	FieldRules struct {
		field interface{}
		rules []Rule
	}
)

func Field(field interface{}, rules ...Rule) *FieldRules {
	return &FieldRules{
		field: field,
		rules: rules,
	}
}
