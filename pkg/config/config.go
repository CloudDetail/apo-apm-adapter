package config

type AdapterConfig struct {
	HttpPort      int                   `mapstructure:"http_port"`
	Timeout       int64                 `mapstructure:"timeout"`
	TraceApi      *TraceApiConfig       `mapstructure:"trace_api"` // Default
	ClusterAPIMap ClusterTraceAPIConfig `mapstructure:"clusters"`  // ClusterScope API
}

type ClusterTraceAPIConfig map[string]*TraceApiConfig

type TraceApiConfig struct {
	ApmList    []string          `mapstructure:"apm_list"`
	Skywalking *SkywalkingConfig `mapstructure:"skywalking"`
	Jaeger     *JaegerConfig     `mapstructure:"jaeger"`
	Elastic    *ElasticConfig    `mapstructure:"elastic"`
	Pinpoint   *PinpointConfig   `mapstructure:"pinpoint"`
}

type SkywalkingConfig struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type JaegerConfig struct {
	Address string `mapstructure:"address"`
}

type ElasticConfig struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type PinpointConfig struct {
	Address string `mapstructure:"address"`
}
