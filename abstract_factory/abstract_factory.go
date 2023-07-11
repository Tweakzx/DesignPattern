package factory

// 规则配置解析
type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct{}

func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// 系统配置解析
type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

type jsonSystemConfigParser struct{}

func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

// 抽象工厂
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct{}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
