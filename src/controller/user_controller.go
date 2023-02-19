package controller

import (
	entities "main/src/entities/user"
	userCreate "main/src/handler/user/create"
	userFind "main/src/handler/user/find"
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

// @Summary      Access token
// @Description  Pegar usuário logado
// @Tags         Usuário
// @Accept       json
// @Produce      json
// @Success      200  {object}  userFind.ResponseFindUser
// @Failure      400  {object}  error.CustomError
// @Router       /user [get]
// @Security ApiKeyAuth
func findUserByToken(rg *gin.RouterGroup) {
	rg.GET("/", func(ctx *gin.Context) {
		u := userFind.New(language.Make(utils.GetHeaderValue(ctx, "Content-Language")))
		var userToken, _ = ctx.Get("User")
		var list, err = u.FindUser(userToken.(*entities.UserToken))
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
	user := r.Group("/user")
	{
		user.Use(utils.UserAuth)
		findUserByToken(user)
	}
}
