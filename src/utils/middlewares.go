package utils

import (
	custom_error "main/src/entities/custom_error"
	entities "main/src/entities/user"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func LangVerify(c *gin.Context) {
	var lang = GetHeaderValue(c, "Content-Language")
	if lang == "" {
		c.Request.Header["Content-Language"] = []string{"pt-BR"}
	} else if lang != "pt-BR" && lang != "en" {
		c.AbortWithStatusJSON(400, custom_error.NewCode("The language must be pt-BR or en", "LANGUAGE_ERR"))
	} else {
		c.Next()
	}
}

/*
Recupera as informacoes do token, se estiverem validas
*/
func getToken(c *gin.Context) (*entities.UserToken, *custom_error.CustomError) {
	var token = GetHeaderValue(c, "Authorization")
	var lang = language.Make(GetHeaderValue(c, "Content-Language"))
	if token == "" {
		err := custom_error.NewCode(GetString("missingToken", lang), "TOKEN")
		return nil, &err
	}
	jwt := strings.Split(token, " ")
	if len(jwt) != 2 {
		err := custom_error.NewCode(GetString("invalidToken", lang), "TOKEN")
		return nil, &err
	}
	var values, err = ValidateJWT(jwt[1])
	if err != nil {
		return nil, err
	}
	var s, _ = ParseInt(values["status"])
	user := entities.UserToken{Email: values["email"], Status: int8(s), Id: values["user"]}
	return &user, nil
}

/*
Para acessar rota, precisa está autenticado
*/
func UserAuth(c *gin.Context) {
	var userToken, err = getToken(c)
	if err != nil {
		c.AbortWithStatusJSON(401, err)
		return
	}
	c.Set("User", userToken)
	c.Next()
}

/*
Para acessar rota, precisa está autenticado
*/
func AdminAuth(c *gin.Context) {
	var userToken, err = getToken(c)
	if err != nil {
		c.AbortWithStatusJSON(401, err)
		return
	}
	if userToken.Status != 2 {
		var lang = language.Make(GetHeaderValue(c, "Content-Language"))
		err := custom_error.NewCode(GetString("adminError", lang), "WITHOUT_PERMISSION")
		c.AbortWithStatusJSON(401, err)
		return
	}
	c.Set("User", userToken)
	c.Next()
}
