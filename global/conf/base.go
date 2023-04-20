package conf

type BaseConf struct {
	Base    Base    `mapstructure:"base"`
	Http    Http    `mapstructure:"http"`
	Log     Log     `mapstructure:"log"`
	Swagger Swagger `mapstructure:"swagger"`
	Cluster Cluster `mapstructure:"cluster"`
}

type Base struct {
	DebugMode    string `mapstructure:"debug_mode"`
	TimeLocation string `mapstructure:"time_location"`
}

type Http struct {
	Addr           string   `mapstructure:"addr"`
	ReadTimeout    int      `mapstructure:"read_timeout"`
	WriteTimeout   int      `mapstructure:"write_timeout"`
	MaxHeaderBytes int      `mapstructure:"max_header_bytes"`
	AllowIp        []string `mapstructure:"allow_ip"`
}

type Log struct {
	AppName       string        `mapstructure:"name"`
	LogFileDir    string        `mapstructure:"logFileDir"`
	MaxSize       int           `mapstructure:"maxSize"`
	MaxBackups    int           `mapstructure:"maxBackups"`
	MaxAge        int           `mapstructure:"maxAge"`
	Dev           bool          `mapstructure:"dev"`
	LogLevel      string        `mapstructure:"log_level"`
	FileWriter    FileWriter    `mapstructure:"file_writer"`
	ConsoleWriter ConsoleWriter `mapstructure:"console_writer"`
}

type FileWriter struct {
	On                 bool   `mapstructure:"on"`
	LogPath            string `mapstructure:"log_path"`
	RotateLogPath      string `mapstructure:"rotate_log_path"`
	WfLogPath          string `mapstructure:"wf_log_path"`
	rotate_wf_log_path string `mapstructure:"rotate_wf_log"`
}

type ConsoleWriter struct {
	On    bool `mapstructure:"on"`
	Color bool `mapstructure:"color"`
}

type Swagger struct {
	Title    string `mapstructure:"title"`
	Desc     string `mapstructure:"desc"`
	Host     string `mapstructure:"host"`
	BasePath string `mapstructure:"base_path"`
}

type Cluster struct {
	ClusterIp      string `mapstructure:"cluster_ip"`
	ClusterPort    string `mapstructure:"cluster_port"`
	ClusterSSLPort string `mapstructure:"cluster_ssl_port"`
}
