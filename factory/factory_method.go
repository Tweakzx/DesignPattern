package factory

import "fmt"

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type yamlRuleConfigParserFactory struct{}

func (Y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	fmt.Println("yamlRuleConfigParserFactory step 1")
	fmt.Println("yamlRuleConfigParserFactory step 2")
	return yamlRuleConfigParser{}
}

type jsonRuleConfigParserFactory struct{}

func (J jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	fmt.Println("jsonRuleConfigParserFactory step 1")
	fmt.Println("jsonRuleConfigParserFactory step 2")
	return jsonRuleConfigParser{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	default:
		return nil
	}
}
