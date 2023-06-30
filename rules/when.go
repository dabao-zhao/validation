package rules

import "github.com/dabao-zhao/validation"

type WhenRule struct {
	condition bool
	rule      validation.Rule
	elseRule  validation.Rule
}

// When 当 condition 为 true 时验证 rule，为 false 时验证 elseRule
func When(condition bool, rule validation.Rule) WhenRule {
	return WhenRule{
		condition: condition,
		rule:      rule,
	}
}

func (r WhenRule) Else(rule validation.Rule) WhenRule {
	r.elseRule = rule
	return r
}

func (r WhenRule) Validate(key, value interface{}) error {
	if r.condition {
		return r.rule.Validate(key, value)
	}
	if r.elseRule != nil {
		return r.elseRule.Validate(key, value)
	}
	return nil
}
