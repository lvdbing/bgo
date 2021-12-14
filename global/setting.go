package global

import "time"

var (
	ServerSetting   *ServerSettings
	AppSetting      *AppSettings
	DatabaseSetting *DatabaseSettings
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
