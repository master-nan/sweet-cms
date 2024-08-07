package config

type Server struct {
	Name     string  `mapStructure:"name"`
	Version  string  `mapStructure:"version"`
	Port     int     `mapStructure:"port"`
	DB       DB      `mapStructure:"db"`
	Redis    Redis   `mapStructure:"redis"`
	Session  Session `mapStructure:"session"`
	WorkerID int64   `mapStructure:"worker_id"`
	Conf     Conf    `mapStructure:"conf"`
}

type DB struct {
	Host     string `mapStructure:"host"`
	Port     int    `mapStructure:"port"`
	Name     string `mapStructure:"name"`
	User     string `mapStructure:"user"`
	Password string `mapStructure:"password"`
	Prefix   string `mapStructure:"prefix"`
}

type Redis struct {
	Host            string `mapStructure:"host"`
	Port            int    `mapStructure:"port"`
	DB              int    `mapStructure:"db"`
	Password        string `mapStructure:"password"`
	PoolSize        int    `mapStructure:"pool_size"`
	MinIdleConns    int    `mapStructure:"min_idle_conns"`
	MaxIdleConns    int    `mapStructure:"max_idle_conns"`
	ConnMaxIdleTime int    `mapStructure:"conn_max_idle_time"`
}

type Session struct {
	Secret string `mapStructure:"secret"`
}

type Conf struct {
	Salt string `mapStructure:"salt"`
}
