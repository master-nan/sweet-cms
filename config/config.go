package config

type Server struct {
	Name      string    `mapStructure:"name"`
	Version   string    `mapStructure:"version"`
	Port      int       `mapStructure:"port"`
	DB        DB        `mapStructure:"db"`
	Redis     Redis     `mapStructure:"redis"`
	Session   Session   `mapStructure:"session"`
	Configure Configure `mapStructure:"conf"`
}

type DB struct {
	Host     string `mapStructure:"host"`
	Port     int    `mapStructure:"port"`
	Name     string `mapStructure:"name"`
	User     string `mapStructure:"user"`
	Password string `mapStructure:"password"`
}

type Redis struct {
	Host string `mapStructure:"host"`
	Port int    `mapStructure:"port"`
}

type Session struct {
	Secret string `mapStructure:"secret"`
}

type Configure struct {
	Salt string `mapStructure:"salt"`
}
