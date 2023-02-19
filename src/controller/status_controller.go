package controller

import (
	s "main/src/entities/status"
	statusListAll "main/src/handler/status/list_all"
	status "main/src/handler/status/update_delete_create_status"
	"main/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListStatus godoc
// @Summary      List status
// @Description  get all status
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object}  statusListAll.ResponseListStatus
// @Failure      400  {object}  error.CustomError
// @Failure      404  {object}  error.CustomError
// @Failure      500  {object}  error.CustomError
// @Router       /status [get]
func listAll(rg *gin.RouterGroup) {
	rg.GET("/", func(ctx *gin.Context) {
		var list, err = statusListAll.ListAllStatus()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			// var userToken, _ = ctx.Get("User")
			// var user, _ = utils.TypeConverter[u.User](userToken)
			// fmt.Println(user.Id)
			ctx.JSON(http.StatusOK, list)
		}
	})
}

// create godoc
// @Summary      Criação de um novo status
// @Description  Cria um novo status de usuário
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object}  status.ResponseStatus
// @Failure      400  {object}  error.CustomError
// @Failure      404  {object}  error.CustomError
// @Failure      500  {object}  error.CustomError
// @Router       /status [post]
// @Param status   body s.Status true "Status"
func create(rg *gin.RouterGroup) {
	rg.POST("/", func(ctx *gin.Context) {
		var requestBody s.Status
		ctx.BindJSON(&requestBody)
		var list, err = status.CreateStatus(requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			ctx.JSON(http.StatusOK, list)
		}
	})
}

// create godoc
// @Summary      Atualiza um novo status
// @Description  Atualiza status de usuário
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object}  status.ResponseStatus
// @Failure      400  {object}  error.CustomError
// @Failure      404  {object}  error.CustomError
// @Failure      500  {object}  error.CustomError
// @Router       /status/{id} [patch]
// @Param id   path string true "Id"
// @Param status   body s.Status true "o ID passado no body é desconsiderado e pego a da URL"
func update(rg *gin.RouterGroup) {
	rg.PATCH("/:id", func(ctx *gin.Context) {
		var requestBody s.Status
		ctx.BindJSON(&requestBody)
		var list, err = status.UpdateStatus(requestBody, ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			ctx.JSON(http.StatusOK, list)
		}
	})
}

// create godoc
// @Summary      Deleta um status
// @Description  Deleta status de usuário
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object}  status.ResponseStatus
// @Failure      400  {object}  error.CustomError
// @Failure      404  {object}  error.CustomError
// @Failure      500  {object}  error.CustomError
// @Router       /status/{id} [delete]
// @Param id   path string true "Id"
func delete(rg *gin.RouterGroup) {
	rg.DELETE("/:id", func(ctx *gin.Context) {
		var list, err = status.DeleteStatus(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			ctx.JSON(http.StatusOK, list)
		}
	})
}

// create godoc
// @Summary      Pegar por id
// @Description  Pega status por id
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object}  statusListAll.ResponseListStatus
// @Failure      400  {object}  error.CustomError
// @Failure      404  {object}  error.CustomError
// @Failure      500  {object}  error.CustomError
// @Router       /status/{id} [get]
// @Param id   path string true "Id"
func byId(rg *gin.RouterGroup) {
	rg.GET("/:id", func(ctx *gin.Context) {
		var list, err = statusListAll.ListStatusById(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			ctx.JSON(http.StatusOK, list)
		}
	})
}

func StatusController(r *gin.Engine) {
	rg := r.Group("/status")
	{
		rg.Use(utils.UserAuth)
		listAll(rg)
		byId(rg)
	}
	rgAdmin := r.Group("/status")
	{
		rgAdmin.Use(utils.AdminAuth)
		create(rgAdmin)
		update(rgAdmin)
		delete(rgAdmin)
	}
}
