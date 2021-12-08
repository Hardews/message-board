package dao

func PostComment(username, userPost, comment string) error {
	sqlStr := "insert into userComment (username,userPosr,comment) values (?,?,?)"
	_, err := dB.Exec(sqlStr, username, userPost, comment)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(username, comment string) error {
	sqlStr := "delete comment from userComment where username = ? and comment = ?"
	_, err := dB.Exec(sqlStr, username, comment)
	if err != nil {
		return err
	}
	return nil
}

func SelectComment(username, comment string) (error, bool) {
	var checkName, CheckComment string
	sqlStr := "select username,comment from userComment where username = ? and comment = ?"
	rows, err := dB.Query(sqlStr, username, comment)
	if err != nil {
		return err, false
	}
	for rows.Next() {
		err := rows.Scan(&checkName, &CheckComment)
		if checkName == username && CheckComment == comment {
			return nil, true
		}
		if err != nil {
			return err, false
		}
	}
	return err, true
}

func ChangeComment(username, oldComment, newComment string) error {
	sqlStr := "update userComment set comment = ? where username = ? and comment = ?"
	_, err := dB.Exec(sqlStr, newComment, username, oldComment)
	if err != nil {
		return err
	}
	return nil
}
