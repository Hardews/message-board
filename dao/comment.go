package dao

func PostComment(username, comment string, postId int) error {
	sqlStr := "insert into userComment (postid,username,comment) values (?,?,?)"
	_, err := dB.Exec(sqlStr, postId, username, comment)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(id int) error {
	sqlStr := "delete comment from userComment where id = ?"
	_, err := dB.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

func SelectComment(username, comment string) (error, bool) {
	var checkName, CheckComment string
	sqlStr := "select username,comment from userComment where username = ?"
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

func ChangeComment(oldComment, newComment string) error {
	sqlStr := "update userComment set comment = ? where comment = ? "
	_, err := dB.Exec(sqlStr, newComment, oldComment)
	if err != nil {
		return err
	}
	return nil
}

func SelectByCommentId(username, comment string, postID int) (error, int) {
	var id int
	sqlStr := "select id from userComment where username = ? and comment = ? and postID = ?"
	err := dB.QueryRow(sqlStr, username, comment, postID).Scan(&id)
	if err != nil {
		return err, id
	}
	return nil, id
}
