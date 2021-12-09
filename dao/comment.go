package dao

import "message-board/model"

func AddComment(commentUser model.Comment) error {
	sqlStr := "insert into userComment (postid,commentName,comment) values (?,?,?)"
	_, err := dB.Exec(sqlStr, commentUser.PostID, commentUser.Username, commentUser.Txt)
	if err != nil {
		return err
	}
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

func SelectComment(username, comment string) (error, bool) {
	var checkName, CheckComment string
	sqlStr := "select commentName,comment from userComment where commentName= ?"
	rows, err := dB.Query(sqlStr, username)
	if err != nil {
		return err, false
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&checkName, &CheckComment)
		if checkName == username && CheckComment == comment {
			return nil, true
		}
		if err != nil {
			return err, false
		}
	}
	return err, false
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
