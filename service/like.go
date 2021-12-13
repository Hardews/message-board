package service

import (
	"message-board/dao"
	"message-board/model"
)

func SelectCommentLikeNum(commentId int) (error, int) {
	var err error
	var user model.Comment
	user.LikeNum, err = dao.SelectCommentLikeNum(commentId, user)
	if err != nil {
		return err, user.LikeNum
	}
	return err, user.LikeNum
}

func LikeComment(LikeNum int, info model.Comment, username string) error {
	LikeNum += 1
	err := dao.LikeComment(LikeNum, info, username)
	if err != nil {
		return err
	}
	return err
}

func SelectPostNum(postId int) (error, int) {
	var err error
	var user model.Post
	user.LikeNum, err = dao.SelectPostLikeNum(postId, user)
	if err != nil {
		return err, user.LikeNum
	}
	return err, user.LikeNum
}

func LikePost(LikeNum, postID int, username string) error {
	LikeNum += 1
	err := dao.LikePost(LikeNum, postID, username)
	if err != nil {
		return err
	}
	return err
}
