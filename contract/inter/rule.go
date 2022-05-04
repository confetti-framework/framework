package inter

import "github.com/confetti-framework/framework/support"

type Rule interface {
	Verify(value support.Value) error
}

type RuleWithRequirements interface {
	Rule
	Requirements() []Rule
}

type RuleWithApp interface {
	Rule
	SetApp(app AppReader) Rule
}
