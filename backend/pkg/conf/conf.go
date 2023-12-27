package conf

var Conf Config

type Config struct {
	Base  Base  `yaml:"base"`
	Log   Log   `yaml:"log"`
	Proof Proof `yaml:"proof"`
}

type Base struct {
	IsDebug        bool   `mapstructure:"is_debug"`
	Domain         string `mapstructure:"domain"`
	MaxConcurrency int64  `mapstructure:"max_concurrency"`
}

type Log struct {
	InfoPath string `yaml:"InfoPath"`
	ErrPath  string `yaml:"ErrPath"`
}

type Proof struct {
	InfoLog string `mapstructure:"info_log"`
	ErrLog  string `mapstructure:"err_log"`
}
