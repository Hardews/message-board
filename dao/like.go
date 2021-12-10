package dao

import (
	"message-board/model"
)

func SelectCommentLikeNum(id int, user model.Comment) (int, error) {
	sqlStr := "select likeNum from userComment where id = ?"
	err := dB.QueryRow(sqlStr, id).Scan(&user.LikeNum)
	if err != nil {
		return 0, err
	}

	return user.LikeNum, err
}

func LikeComment(LikeNum int, info model.Comment, username string) error {
	sqlStr1 := "update userComment set likeNum = ? where id = ?"
	_, err := dB.Exec(sqlStr1, LikeNum, info.CommentId)
	if err != nil {
		return err
	}

	sqlStr2 := "insert into commentLike (username,PostID,CommentID) values (?,?,?)"
	_, err = dB.Exec(sqlStr2, username, info.PostID, info.CommentId)
	if err != nil {
		return err
	}
	return err
}

func SelectPostLikeNum(id int, user model.Post) (int, error) {
	sqlStr := "select likeNum from userPost where id = ?"
	err := dB.QueryRow(sqlStr, id).Scan(&user.LikeNum)
	if err != nil {
		return 0, err
	}

	return user.LikeNum, err
}

func LikePost(LikeNum, PostID int, username string) error {
	sqlStr1 := "update userPost set likeNum = ? where id = ?"
	_, err := dB.Exec(sqlStr1, LikeNum, PostID)
	if err != nil {
		return err
	}

	sqlStr2 := "insert into postLike (username,PostID) values (?,?)"
	_, err = dB.Exec(sqlStr2, username, PostID)
	if err != nil {
		return err
	}
	return err
}
