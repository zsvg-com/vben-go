package constants

/*公共常量*/
// menu
const (
	AdminId = 1

	IsFrameYes = "0"
	IsFrameNo  = "1"

	isCacheYes = "0"
	isCacheNo  = "1"

	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

const (
	LoginCacheKey        = "go_login_tokens:"
	CaptchaCodesKey      = "captcha_codes:"
	SysDictCacheKey      = "sys_dict:"
	SysConfigCacheKey    = "sys_config:"
	PwdErrCntCacheKey    = "pwd_err_cnt:"
	RepeatSubmitCacheKey = "repeat_submit:"
	RateLimitCacheKey    = "rate_limit:"
	ScanCountMax         = 1000
)
