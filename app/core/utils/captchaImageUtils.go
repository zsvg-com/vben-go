package utils

import (
	"image/color"
	"vben/common/config"
	"vben/pkg/cache"

	"github.com/mojocn/base64Captcha"
)

// var store = base64Captcha.NewMemoryStore(20240, 3*time.Minute)
// 使用redis缓存
var store base64Captcha.Store = cache.RedisStore{}

// CreateImageCaptcha 生产 验证码
func CreateImageCaptcha() (string, string, error) {
	var driver base64Captcha.Driver

	if config.LoginVerificationCode == "Number" {
		driver = letterConfig()
	}

	if config.LoginVerificationCode == "Arithmetic" {
		driver = mathConfig()
	}
	if driver == nil {
		panic("生成验证码的类型没有配置，请在yaml文件中配置完再次重试启动项目")
	}
	// 创建验证码并传入创建的类型的配置，以及存储的对象
	c := base64Captcha.NewCaptcha(driver, store)
	return c.Generate()
}

// VerifyCaptcha 验证 验证码
func VerifyCaptcha(Uuid string, Code string) bool {
	return store.Verify(Uuid, Code, true)
}

// 配置 算数 验证码
func mathConfig() *base64Captcha.DriverMath {
	mathType := &base64Captcha.DriverMath{
		Height:          50,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: []string{"RitaSmith.ttf"},
	}
	return mathType
}

// 配置 数字 验证码
func letterConfig() *base64Captcha.DriverString {
	driverString := base64Captcha.DriverString{
		Height:          50,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"RitaSmith.ttf"},
	}

	return driverString.ConvertFonts()
}
