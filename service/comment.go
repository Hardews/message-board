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

func DeleteComment(commentId, PostID int) error {
	err := dao.DeleteComment(commentId, PostID)
	if err != nil {
		return err
	}
	return err
}

func ChangeComment(newComment string, commentID int) error {
	err := dao.ChangeComment(newComment, commentID)
	if err != nil {
		return err
	}
	return err
}
