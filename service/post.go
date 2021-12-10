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

func GetPost(PostID int) (error, []model.Post) {
	err, posts := dao.SelectCommentsSection(PostID)
	if err != nil {
		return err, posts
	}
	return err, posts
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

func CreateCommentsSection(ID int, Post model.Post) error {
	err := dao.CreateCommentsSection(ID, Post)
	if err != nil {
		return err
	}
	return err
}
