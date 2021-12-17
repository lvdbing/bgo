package global

import "time"

var (
	ServerSetting   *ServerSettings
	AppSetting      *AppSettings
	DatabaseSetting *DatabaseSettings
	JWTSetting      *JWTSettings
	EmailSetting    *EmailSettings
	LimiterSetting  *LimiterSettings
)

type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettings struct {
	DefaultPagesize    int
	MaxPagesize        int
	LogPath            string
	LogFilename        string
	LogFileExt         string
	UploadPath         string
	UploadUrl          string
	UploadMaxSizeImg   int
	UploadMaxSizeExcel int
	UploadMaxSizeWord  int
	UploadMaxSizePPT   int
	UploadMaxSizeTxt   int
	UploadExtsImg      []string
	UploadExtsExcel    []string
	UploadExtsWord     []string
	UploadExtsPPT      []string
	UploadExtsTxt      []string
	ContextTimeout     time.Duration
}

type DatabaseSettings struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSettings struct {
	From     string
	Host     string
	Port     int
	Username string
	Password string
	IsSSL    bool
	To       []string
}

type LimiterSettings struct {
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
