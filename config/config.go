package config

var Config *config

func init() {
	Config = &config{}
}

type route struct {
	Method string `yaml:"method"`
	Result string `yaml:"result"`
}

type config struct {
	Routes map[string]*route `yaml:"routes"`
}
