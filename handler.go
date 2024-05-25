package main

type RuleHandler struct {
	rule Rule
}

func NewRuleHandlers(config Config) (ruleHandlers []RuleHandler) {
	for _, rule := range config.Rules {
		ruleHandlers = append(ruleHandlers, RuleHandler{rule})
	}
	return ruleHandlers
}
