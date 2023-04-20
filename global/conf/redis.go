package conf

type RedisConf struct {
	List map[string]*Redis `mapstructure:"list"`
}

type Redis struct {
	Name     string `mapstructure:"name"`
	Ip       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
