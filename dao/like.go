package dao

import (
	"message-board/model"
	"strconv"
)

func SelectCommentLikeNum(id int, user model.Comment) (int, error) {
	sqlStr := "select commentLikeNum from userComment where id = ?"
	err := dB.QueryRow(sqlStr, id).Scan(&user.LikeNum)
	if err != nil {
		return 0, err
	}

	return user.LikeNum, err
}

func LikeComment(LikeNum, id int, info model.Comment, username string) error {
	sqlStr1 := "update userComment set commentLikeNum = ? where id = ?"
	_, err := dB.Exec(sqlStr1, LikeNum, info.CommentId)
	if err != nil {
		return err
	}

	sqlStr2 := "insert into commentLike (username,PostID,CommentID) values (?,?,?)"
	_, err = dB.Exec(sqlStr2, username, info.PostID, info.CommentId)
	if err != nil {
		return err
	}

	strNum := "post" + strconv.Itoa(info.PostID)
	sqlStrM := "update " + strNum + " set LikeNum = ? where id = ?"
	sqlStr3 := sqlStrM
	_, err = dB.Exec(sqlStr3, LikeNum, id)
	if err != nil {
		return err
	}
	return err
}

func SelectPostLikeNum(id int, user model.Post) (int, error) {
	sqlStr := "select postLikeNum from userPost where id = ?"
	err := dB.QueryRow(sqlStr, id).Scan(&user.LikeNum)
	if err != nil {
		return 0, err
	}

	return user.LikeNum, err
}

func LikePost(LikeNum, PostID int, username string) error {
	sqlStr1 := "update userPost set postLikeNum = ? where id = ?"
	_, err := dB.Exec(sqlStr1, LikeNum, PostID)
	if err != nil {
		return err
	}

	sqlStr2 := "insert into postLike (username,PostID) values (?,?)"
	_, err = dB.Exec(sqlStr2, username, PostID)
	if err != nil {
		return err
	}

	strNum := "post" + strconv.Itoa(PostID)
	sqlStrM := "update " + strNum + " set LikeNum = ? where id = 1"
	sqlStr3 := sqlStrM
	_, err = dB.Exec(sqlStr3, LikeNum)
	if err != nil {
		return err
	}
	return err
}
