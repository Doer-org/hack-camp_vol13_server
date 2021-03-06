package handler

import (
	"net/http"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	uR repository.UserRepository
}

type UserHandler interface {
	CreateUser(*gin.Context)
	DeleteUser(*gin.Context)
	UpdateUser(*gin.Context)
	FindUserById(*gin.Context)
	FindAllUser(*gin.Context)
	GetUserRanking(ctx *gin.Context)
}

func NewUserHandler(uR repository.UserRepository) UserHandler {
	return &userHandler{uR: uR}
}

// user作成
func (uH *userHandler) CreateUser(ctx *gin.Context) {

	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	resUser, err := uH.uR.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resUser)
}

// user 削除
func (uH *userHandler) DeleteUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	err := uH.uR.DeleteUser(&model.User{Id: uid})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// user 更新
func (uH *userHandler) UpdateUser(ctx *gin.Context) {
	user := &model.User{}

	if err := ctx.Bind(user); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	uid := ctx.Param("uid")
	user.Id = uid

	err := uH.uR.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// uid で userを検索
func (uH *userHandler) FindUserById(ctx *gin.Context) {
	uid := ctx.Param("uid")
	user, err := uH.uR.FindUserById(uid)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

// すべてのuserを返す
func (uH *userHandler) FindAllUser(ctx *gin.Context) {
	users, err := uH.uR.FindAllUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// 上位10名の userを表示 (Scoreを基準とする)
func (uH *userHandler) GetUserRanking(ctx *gin.Context) {
	top_users, err := uH.uR.GetUserRanking()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, top_users)
}
