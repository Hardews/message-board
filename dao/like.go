package dao

import (
	"message-board/model"
)

func SelectCommentLikeNum(id int, user model.Comment) (int, error) {
	sqlStr := "select commentLikeNum from user_Comment where id = ?"
	err := dB.QueryRow(sqlStr, id).Scan(&user.LikeNum)
	if err != nil {
		return 0, err
	}

	return user.LikeNum, err
}

func LikeComment(LikeNum int, info model.Comment, username string) error {
	sqlStr1 := "update user_Comment set commentLikeNum = ? where id = ?"
	_, err := dB.Exec(sqlStr1, LikeNum, info.CommentId)
	if err != nil {
		return err
	}

	sqlStr2 := "insert into comment_Like (username,PostID,CommentID) values (?,?,?)"
	_, err = dB.Exec(sqlStr2, username, info.PostID, info.CommentId)
	if err != nil {
		return err
	}

	return nil
}

func SelectPostLikeNum(id int, user model.Post) (int, error) {
	sqlStr := "select postLikeNum from user_Post where id = ?"
	err := dB.QueryRow(sqlStr, id).Scan(&user.LikeNum)
	if err != nil {
		return 0, err
	}

	return user.LikeNum, err
}

func LikePost(LikeNum, PostID int, username string) error {
	sqlStr1 := "update user_Post set postLikeNum = ? where id = ?"
	_, err := dB.Exec(sqlStr1, LikeNum, PostID)
	if err != nil {
		return err
	}

	sqlStr2 := "insert into post_Like (username,PostID) values (?,?)"
	_, err = dB.Exec(sqlStr2, username, PostID)
	if err != nil {
		return err
	}

	return err
}
