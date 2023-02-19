package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	custom_error "main/src/entities/custom_error"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetHeaderValue(ctx *gin.Context, key string) string {
	if ctx.Request.Header[key] == nil {
		return ""
	}
	return ctx.Request.Header[key][0]
}

func ParseInt(val string) (n int, err *custom_error.CustomError) {
	var er error
	n, er = strconv.Atoi(val)
	if er != nil {
		err := custom_error.New("Parâmetro deve ser um número")
		return n, &err
	}
	return n, err
}

func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func GenerateJWT(data map[string]string) (string, *custom_error.CustomError) {
	var clains = jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	}
	for k, v := range data {
		clains[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
	var tokenString, e = token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if e != nil {
		err := custom_error.New("GenerateJWT: " + e.Error())
		return "", &err
	}
	return tokenString, nil
}

func ValidateJWT(token string) (map[string]string, *custom_error.CustomError) {
	claims := jwt.MapClaims{}
	var _, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		var ce = custom_error.NewCode(err.Error(), "TOKEN_VALIDATE")
		if err.Error() == "Token is expired" {
			ce = custom_error.NewCode(err.Error(), "TOKEN_EXPIRED")
		}
		return nil, &ce
	}

	values := make(map[string]string)
	for k, v := range claims {
		switch v.(type) {
		case string:
			values[k] = v.(string)
			break
		default:
			break
		}
	}
	return values, nil
}

func Md5(val string) string {
	hasher := md5.New()
	hasher.Write([]byte(val))
	return hex.EncodeToString(hasher.Sum(nil))
}
