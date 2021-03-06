package handler

import (
	"net/http"
	"strconv"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
	"github.com/gin-gonic/gin"
)

type voteCommentHandler struct {
	vcR repository.VoteCommentRepository
}

type VoteCommentHandler interface {
	IncreaseVoteComment(*gin.Context)
	RevokeVoteComment(*gin.Context)
	FindVoteCommentIdOfVoted(*gin.Context)
}

func NewVoteCommentHandler(vcR repository.VoteCommentRepository) VoteCommentHandler {
	return &voteCommentHandler{vcR: vcR}
}

// good/bad を増やす
func (vcH *voteCommentHandler) IncreaseVoteComment(ctx *gin.Context) {
	comment_vote := &model.VoteComment{}
	if err := ctx.Bind(comment_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := vcH.vcR.IncreaseVoteComment(comment_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// good/bad の取り消し
func (vcH *voteCommentHandler) RevokeVoteComment(ctx *gin.Context) {
	comment_vote := &model.VoteComment{}
	if err := ctx.Bind(comment_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := vcH.vcR.RevokeVoteComment(comment_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// good/bad済みか
// もしすでにしてるものがあればそのcomment_idを返す

// type ReqBodyVoteComment struct{
// 	ThreadId int `json:"thread_id"`
// 	UserId string `json:"uid"`
// }

func (vcH *voteCommentHandler) FindVoteCommentIdOfVoted(ctx *gin.Context) {

	uid := ctx.Param("uid")

	threadIdString := ctx.Param("thread_id")
	threadId, err := strconv.Atoi(threadIdString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	vote_comments, err := vcH.vcR.FindVoteCommentIdOfVoted(uid, threadId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vote_comments)
}
