package global

import "time"

const (
	MethodPost     = "POST"
	MethodGet      = "GET"
	MethodPut      = "PUT"
	MethodDelete   = "DELETE"
	FILEPATH       = "./static"
	ScratchFile    = "./static/scratch"
	LogFilePath    = "./logs"
	ConfigFilePath = "./config"
	TempPath       = "./temp"
	Pong           = "pong"
	MaxUploadSize  = 10 << 20 // 10M
	PAGE_START     = 1
	PAGE_SIZE_MAX  = 1000
	PAGE_SIZE_MIN  = 10

	ValidatorKey        = "ValidatorKey"
	TranslatorKey       = "TranslatorKey"
	AdminSessionInfoKey = "AdminSessionInfoKey"

	LoadTypeHTTP = 0
	LoadTypeTCP  = 1
	LoadTypeGRPC = 2

	HTTPRuleTypePrefixURL = 0
	HTTPRuleTypeDomain    = 1

	RedisFlowDayKey  = "flow_day_count"
	RedisFlowHourKey = "flow_hour_count"

	FlowTotal         = "flow_total"
	FlowServicePrefix = "flow_service_"
	FlowAppPrefix     = "flow_app_"

	JwtSignKey = "my_sign_key"
	JwtExpires = 60 * 60

	DataAlive = time.Duration(86400*2) * time.Second
)
