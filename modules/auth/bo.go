package auth

type LoginBo struct {
	ClientId  string
	GrantType string
	TenantId  string
	Code      string
	Uuid      string
}
