package ctx

import "chat/pkg/env"

const (
	httpPort       = "HTTP_PORT"
	logLevel       = "LOG_LEVEL"
	OPENAI_API_KEY = "OPENAI_API_KEY"
)

type Conf struct {
	httpPort  string
	logLevel  string
	openAIKey string
}

func loadConf() *Conf {
	return &Conf{
		httpPort:  env.GetEnvValueWithFallback(httpPort, "8080"),
		logLevel:  env.GetEnvValueWithFallback(logLevel, "info"),
		openAIKey: env.GetEnvValueWithFallback(OPENAI_API_KEY, ""),
	}
}

func (c *Conf) LogLevel() string {
	return c.logLevel
}

func (c *Conf) HttpPort() string {
	return c.httpPort
}

func (c *Conf) OpenAIKey() string {
	return c.openAIKey
}
