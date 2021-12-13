package dao

import (
	"message-board/model"
)

func AddComment(commentUser model.Comment) error {
	sqlStr := "insert into userComment (postid,commentName,comment) values (?,?,?)"
	_, err := dB.Exec(sqlStr, commentUser.PostID, commentUser.Username, commentUser.Txt)
	if err != nil {
		return err
	}

	return nil
}

func DeleteComment(id, PostId int) error {
	sqlStr := "delete from userComment where id = ? and PostID = ?"
	_, err := dB.Exec(sqlStr, id, PostId)
	if err != nil {
		return err
	}
	return nil
}

func ChangeComment(newComment string, commentID int) error {
	sqlStr := "update userComment set comment = ? where id = ? "
	_, err := dB.Exec(sqlStr, newComment, commentID)
	if err != nil {
		return err
	}

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

func SelectByPostID(postName, userPost string) (int, error) {
	var u model.Comment
	sqlStr := "select id from userPost where username= ? and userPost= ?"
	err := dB.QueryRow(sqlStr, postName, userPost).Scan(&u.PostID)
	if err != nil {
		return u.PostID, err
	}
	return u.PostID, err
}
