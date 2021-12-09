package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddPost(username, txt string) error {
	err := dao.Post(username, txt)
	if err != nil {
		return err
	}
	return err
}

func SelectByPostID(username, txt string) (int, error) {
	PostID, err := dao.SelectByPostId(username, txt)
	if err != nil {
		return PostID, err
	}
	return PostID, err
}

func GetPost(PostID int) (error, []model.Post, []model.Comment) {
	err, posts, comments := dao.SelectPostAndCommentByPostID(PostID)
	if err != nil {
		return err, posts, comments
	}
	return err, posts, comments
}

func DeletePost(postId int, post string) error {
	err := dao.DeletePost(postId, post)
	if err != nil {
		return err
	}
	return err
}

func ChangePost(newPost string, PostID int) error {
	err := dao.ChangePost(newPost, PostID)
	if err != nil {
		return err
	}
	return err
}

func GetAllPost() (error, []model.Post, []string) {
	err, user, Time := dao.GetAllPost()
	if err != nil {
		return err, user, Time
	}
	return err, user, Time
}
