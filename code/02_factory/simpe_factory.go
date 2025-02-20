package _2_factory

type ConfigParser interface {
	Parser(data []byte)
}

type jsonConfigParser struct{}

func (j *jsonConfigParser) Parser(data []byte) {}

type tomlConfigParser struct{}

func (t *tomlConfigParser) Parser(data []byte) {}

type yamlConfigParser struct{}

func (y *yamlConfigParser) Parser(data []byte) {}

func NewConfigParser(s string) ConfigParser {
	switch s {
	case "json":
		return &jsonConfigParser{}
	case "toml":
		return &tomlConfigParser{}
	case "yaml":
		return &yamlConfigParser{}

	}
	return nil
}
