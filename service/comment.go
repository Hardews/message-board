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

	err, id := dao.SelectCommentsSectionID(cUser)
	if err != nil {
		return err
	}

	err = dao.DeleteComment2(cUser.PostID, id)
	if err != nil {
		return err
	}
	return err
}

func ChangeComment(newComment string, cUser model.Comment) error {
	err, id := dao.SelectCommentsSectionID(cUser)
	if err != nil {
		return err
	}

	err = dao.ChangeComment(newComment, cUser.CommentId, id)
	if err != nil {
		return err
	}
	return err
}
