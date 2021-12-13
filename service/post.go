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

func SelectAllByPostID(username, txt string) (model.Post, error) {
	user, err := dao.SelectAllByPostId(username, txt)
	if err != nil {
		return user, err
	}
	return user, err
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

func GetCommentsSection(PostID int) (error, []model.Comment) {
	err, comments := dao.GetCommentsSection(PostID)
	if err != nil {
		return err, comments
	}
	return err, comments
}

func GetAllPosts() (error, []model.Post) {
	err, user := dao.SelectAllPost()
	if err != nil {
		return err, user
	}
	return err, user
}
