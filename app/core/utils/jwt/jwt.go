package jwt

import (
	"net/http"
	"strings"
	"time"
	"vben/app/admin/model/constants"
	config2 "vben/common/config"
	"vben/pkg/cache/redisCache"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateToken(UserName string, UserId int, DeptId int, uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_name": UserName,
		"user_id":   UserId,
		"dept_id":   DeptId,
		"exp":       time.Now().Unix() + int64(config2.Jwt.JwtTtl),
		"iss":       "vben-go",
		"uuid":      uuid,
	})

	mySigningKey := []byte(config2.Jwt.Secret)

	return token.SignedString(mySigningKey)
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	mySigningKey := []byte(config2.Jwt.Secret)
	tokenStr = strings.ReplaceAll(tokenStr, config2.HeaderSignTokenStr, "")
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
}

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 根据实际情况取TOKEN, 这里从request header取
		header := ctx.Request.Header
		tokenStr := header.Get(config2.HeaderSignToken)
		if len(tokenStr) < 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "令牌不能为空",
				"code": http.StatusInternalServerError,
			})
			ctx.Abort()
			return
		}

		token, err := VerifyToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "认证失败",
				"code": http.StatusUnauthorized,
			})
			ctx.Abort()
			return
		}
		uuid := token.Claims.(jwt.MapClaims)["uuid"]
		loginUser, err := redisCache.NewRedisCache().Get(constants.LoginCacheKey + uuid.(string))
		if len(loginUser) <= 0 || err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  "认证失败",
				"code": http.StatusUnauthorized,
			})
			ctx.Abort()
			return
		}
		userId := token.Claims.(jwt.MapClaims)["user_id"]
		deptId := token.Claims.(jwt.MapClaims)["dept_id"]
		userName := token.Claims.(jwt.MapClaims)["user_name"]
		// 此处已经通过了, 可以把Claims中的有效信息拿出来放入上下文使用
		ctx.Set("userId", userId)
		ctx.Set("deptId", deptId)
		ctx.Set("userName", userName)
		ctx.Next()
	}
}

func GetJwtUuid(ctx *gin.Context) (string, error) {
	tokenStr := ctx.Request.Header.Get(config2.HeaderSignToken)
	if len(tokenStr) <= 0 {
		return "", nil
	}
	token, err := VerifyToken(tokenStr)
	if err != nil {
		return "", err
	}
	uuid := token.Claims.(jwt.MapClaims)["uuid"]
	return uuid.(string), nil

}
