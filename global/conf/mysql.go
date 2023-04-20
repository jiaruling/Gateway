package conf

type MySQLConf struct {
	List map[string]*MySQL `json:"list"`
}

type MySQL struct {
	Name      string `mapstructure:"name"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Ip        string `mapstructure:"ip"`
	Port      int    `mapstructure:"port"`
	DB        string `mapstructure:"db"`
	Parameter string `mapstructure:"parameter"`
}
