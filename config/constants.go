package config

const (
	// MinGoVersion 最小 Go 版本
	MinGoVersion = "1.16"

	// ProjectVersion 项目版本
	ProjectVersion = "v1.0.2"

	// ProjectName 项目名称
	ProjectName = "vben-go"

	// ProjectDomain 项目域名
	ProjectDomain = "http://127.0.0.1"

	// HeaderSignToken 签名验证 Authorization，Header 中传递的参数
	HeaderSignToken = "Authorization"

	// HeaderSignTokenStr 签名验证 Authorization，Header 中传递的参数
	HeaderSignTokenStr = "Bearer "

	// LoginVerificationCode 登录验证码
	LoginVerificationCode = "Number" //Arithmetic 算数 、Number数字

	// profile 上传路径
	FileProfile = "./static/images/"

	// 查看文件路径
	ShowFileProfile = "/profile/"
)
