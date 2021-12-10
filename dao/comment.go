package dao

import (
	"message-board/model"
	"strconv"
)

func AddComment(commentUser model.Comment) error {
	sqlStr := "insert into userComment (postid,commentName,comment) values (?,?,?)"
	_, err := dB.Exec(sqlStr, commentUser.PostID, commentUser.Username, commentUser.Txt)
	if err != nil {
		return err
	}

	strNum := "post" + strconv.Itoa(commentUser.PostID)
	sqlStrM := "insert into " + strNum + " (username,txt,likeNum) values (?,?,?)"
	sqlStr = sqlStrM
	_, err = dB.Exec(sqlStr, commentUser.Username, commentUser.Txt, 0)
	return nil
}

func DeleteComment(id, PostId int) error {
	sqlStr := "delete comment from userComment where id = ? and PostID = ?"
	_, err := dB.Exec(sqlStr, id, PostId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment2(PostID, id int) error {
	strNum := "post" + strconv.Itoa(PostID)
	sqlStrM := "delete txt from " + strNum + " where id = ?"
	sqlStr := sqlStrM

	_, err := dB.Exec(sqlStr, PostID, id)
	if err != nil {
		return err
	}
	return nil
}

func ChangeComment(newComment string, commentID, id int) error {
	sqlStr := "update userComment set comment = ? where id = ? "
	_, err := dB.Exec(sqlStr, newComment, commentID)
	if err != nil {
		return err
	}

	strNum := "post" + strconv.Itoa(id)
	sqlStrM := "update " + strNum + " set comment = ? where id = ?"
	sqlStr = sqlStrM

	_, err = dB.Exec(sqlStr, newComment, id)
	return nil
}

func SelectByCommentId(cUser model.Comment) (error, int) {
	var id int
	sqlStr := "select id from userComment where commentName = ? and comment = ? and postID = ?"
	err := dB.QueryRow(sqlStr, cUser.Username, cUser.Txt, cUser.PostID).Scan(&id)
	if err != nil {
		return err, id
	}
	return nil, id
}

func SelectCommentsSectionID(cUser model.Comment) (error, int) {
	var id int

	strNum := "post" + strconv.Itoa(cUser.PostID)
	sqlStrM := "select id from " + strNum + " where username = ? and txt = ?"
	sqlStr := sqlStrM

	err := dB.QueryRow(sqlStr, cUser.Username, cUser.Txt).Scan(&id)
	if err != nil {
		return err, id
	}
	return nil, id
}
