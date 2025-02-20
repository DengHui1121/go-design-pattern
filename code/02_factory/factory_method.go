package _2_factory

type ConfigFactory interface {
	createParser() ConfigParser
}

type jsonConfigFactory struct{}

func (factory jsonConfigFactory) createParser() ConfigParser {
	return &jsonConfigParser{}
}

type yamlConfigFactory struct{}

func (factory yamlConfigFactory) createParser() ConfigParser {
	return &yamlConfigParser{}
}

func NewJsonConfigFactory(t string) ConfigFactory {
	switch t {
	case "json":
		return jsonConfigFactory{}
	case "yaml":
		return yamlConfigFactory{}
	}
	return nil
}
