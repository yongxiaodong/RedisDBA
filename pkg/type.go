package pkg

type Config struct {
	Addr           string `yaml:"Addr"`
	Port           int    `yaml:"Port"`
	Password       string `yaml:"Password"`
	DB             int    `yaml:"DB"`
	PullKeysCount  int64  `yaml:"PullKeysCount"`
	PipeQueryCount int    `yaml:"PipeQueryCount"`
	ConsumerNum    int    `yaml:"ConsumerNum"`
}
