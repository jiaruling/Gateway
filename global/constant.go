package global

const (
	MethodPost          = "POST"
	MethodGet           = "GET"
	MethodPut           = "PUT"
	MethodDelete        = "DELETE"
	FILEPATH            = "./static"
	ScratchFile         = "./static/scratch"
	LogFilePath         = "./logs"
	ConfigFilePath      = "./config"
	TempPath            = "./temp"
	Pong                = "pong"
	ValidatorKey        = "ValidatorKey"
	TranslatorKey       = "TranslatorKey"
	AdminSessionInfoKey = "AdminSessionInfoKey"
	MaxUploadSize       = 10 << 20 // 10M
	PAGE_START          = 1
	PAGE_SIZE_MAX       = 1000
	PAGE_SIZE_MIN       = 10
)
