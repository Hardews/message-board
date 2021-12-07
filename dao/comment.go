package dao

func PostComment(username, userPost, comment string) error {
	sqlStr := "insert into comment (username,userPosr,comment) values (?,?,?)"
	_, err := dB.Exec(sqlStr, username, userPost, comment)
	if err != nil {
		return err
	}
	return nil
}
