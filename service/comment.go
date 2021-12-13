package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddComment(commentUser model.Comment) error {
	err := dao.AddComment(commentUser)
	if err != nil {
		return err
	}
	return err
}

func SelectCommentID(cUser model.Comment) (int, error) {
	err, cid := dao.SelectByCommentId(cUser)
	if err != nil {
		return cid, err
	}
	return cid, err
}

func DeleteComment(cUser model.Comment) error {
	err := dao.DeleteComment(cUser.CommentId, cUser.PostID)
	if err != nil {
		return err
	}

	return err
}

func ChangeComment(newComment string, cUser model.Comment) error {
	err := dao.ChangeComment(newComment, cUser.CommentId)
	if err != nil {
		return err
	}
	return err
}

func SelectByPostID(username, txt string) (int, error) {
	id, err := dao.SelectByPostID(username, txt)
	if err != nil {
		return id, err
	}
	return id, err
}
