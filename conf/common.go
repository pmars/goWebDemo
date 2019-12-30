package conf

var (
	Config = &ConfigObj{}
	Hosts  = &SourceHost{}
)

type ConfigObj struct {
	Service struct {
		Addr string
		Mode string
		Html string
	}

	Mysql struct {
		Conn      string
		MaxActive int
		MaxIdle   int
	}

	Redis struct {
		Host        string
		IsAuth      bool
		Password    string
		MaxIdle     int
		MaxActive   int
		DialTimeout int
	}

	WechatPush struct {
		Need      bool
		AppId     string
		Secret    string
		TmpId     string
		UserIds   string
		RedisConn string
		RedisAuth string
		RedisKey  string
	}
	EncryptSecret string
	MiniSecret    string
	MiniAppId     string
	FuncTimeSecs  float64
}

type SourceHost struct {
}
