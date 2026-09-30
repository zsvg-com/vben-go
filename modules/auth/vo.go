package auth

type LoginVo struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	ExpireIn        int64  `json:"expire_in"`
	RefreshExpireIn int64  `json:"refresh_expire_in"`
	ClientId        string `json:"client_id"`
	Scope           string `json:"scope"`
	Openid          string `json:"openid"`
}

type CaptchaVo struct {
	CaptchaEnabled bool   `json:"captchaEnabled"`
	Uuid           string `json:"uuid"`
	Img            string `json:"img"`
}
