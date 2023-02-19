package controller

import (
	userCreate "main/src/handler/user/create"
	userLogin "main/src/handler/user/login"
	utils "main/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

// @Summary      Crated a user to system
// @Description   Crated a user to system
// @Tags         Usuário
// @Accept       json
// @Produce      json
// @Success      200  {object}  userCreate.ResponseCreateUser
// @Failure      400  {object}  error.CustomError
// @Failure      404  {object}  error.CustomError
// @Failure      500  {object}  error.CustomError
// @Router       /user [post]
// @Param user   body userCreate.RequestRegisterUser true "User"
func createUser(rg *gin.RouterGroup) {
	rg.POST("/", func(ctx *gin.Context) {
		var requestBody userCreate.RequestRegisterUser
		ctx.BindJSON(&requestBody)
		u := userCreate.New(language.Make(utils.GetHeaderValue(ctx, "Content-Language")))
		var list, err = u.CreateUser(requestBody, utils.GetHeaderValue(ctx, "User-Agent"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			ctx.JSON(http.StatusOK, list)
		}
	})
}

// @Summary      Access token
// @Description  Get access token
// @Tags         Usuário
// @Accept       json
// @Produce      json
// @Success      200  {object}  userLogin.ResponseLoginUser
// @Failure      400  {object}  error.CustomError
// @Failure      404  {object}  error.CustomError
// @Failure      500  {object}  error.CustomError
// @Router       /user/auth [patch]
// @Param user   body userLogin.RequestLoginUser true "User"
func findUserByEmailAndPassword(rg *gin.RouterGroup) {
	rg.PATCH("/auth", func(ctx *gin.Context) {
		var requestBody userLogin.RequestLoginUser
		ctx.BindJSON(&requestBody)
		u := userLogin.New(language.Make(utils.GetHeaderValue(ctx, "Content-Language")))
		var list, err = u.FindUserByAuth(requestBody, utils.GetHeaderValue(ctx, "User-Agent"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			ctx.JSON(http.StatusOK, list)
		}
	})
}

func UserController(r *gin.Engine) {
	rg := r.Group("/user")
	{
		createUser(rg)
		findUserByEmailAndPassword(rg)
	}
}
