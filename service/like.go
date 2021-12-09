package service

import (
	"message-board/dao"
	"message-board/model"
)

func SelectCommentLikeNum(commentId int) (error, int) {
	var err error
	var user model.Like
	user.CommentLikeNum, err = dao.SelectCommentLikeNum(commentId, user)
	if err != nil {
		return err, user.CommentLikeNum
	}
	return err, user.CommentLikeNum
}

func LikeComment(LikeNum int, info model.Comment, username string) error {
	err := dao.LikeComment(LikeNum, info, username)
	if err != nil {
		return err
	}
	return err
}

func SelectPostNum(postId int) (error, int) {
	var err error
	var user model.Like
	user.PostLikeNum, err = dao.SelectPostLikeNum(postId, user)
	if err != nil {
		return err, user.PostLikeNum
	}
	return err, user.PostLikeNum
}

func LikePost(LikeNum int, info model.Post, username string) error {
	err := dao.LikePost(LikeNum, info, username)
	if err != nil {
		return err
	}
	return err
}
