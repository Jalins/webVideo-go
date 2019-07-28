package db

import (
	"log"
	"time"
	"webVideo-go/model"
	"webVideo-go/utils"

	_ "github.com/go-sql-driver/mysql"
)

// 增加用户
func AddUserCradential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUE (?, ?)")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		log.Panic(err)
	}
	defer stmtIns.Close()

	return nil
}

// 获取用户凭证
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")

	if err != nil {
		log.Panic(err)
		return "", err
	}

	var pwd string
	// 读取一行，然后将匹配的信息赋值到变量中
	stmtOut.QueryRow(loginName).Scan(&pwd)

	stmtOut.Close()
	return pwd, err
}

// 删除user
func DelectUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")

	if err != nil {
		log.Panic(err)
		return err
	}

	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()

	return nil
}

// 删除vedio
func AddNewVedio(aid int, name string) (*model.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()

	if err != nil {
		log.Panic(err)
	}
	t := time.Now()

	ctime := t.Format("Jan 02 2006, 15:04:05") //M D y, HH:MM:SS

	stmt, err := dbConn.Prepare("INSERT INTO video_info (id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)")

	if err != nil {
		log.Panic(err)
	}

	_, err = stmt.Exec(vid, aid, name, ctime)
	if err != nil {
		log.Panic(err)
	}

	result := &model.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmt.Close()

	return result, nil
}

// 删除vedio
func DelVedioInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM vedio_info WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)

	if err != nil {
		return err
	}

	defer stmtDel.Close()

	return nil

}

func AddNewComments(vid string, aid int, content string) error {
	uuid, err := utils.NewUUID()

	if err != nil {
		log.Panic(err)
	}

	stmtIns, err := dbConn.Prepare("INNER INTO  comments (id, video_id, author_id,content) VALUES (?, ?, ?, ?)")

	if err != nil {
		log.Panic(err)
	}

	_, err = stmtIns.Exec(uuid, vid, aid, content)
	if err != nil {
		log.Panic(err)
	}

	defer stmtIns.Close()

	return nil
}

func ListComments(vid string, from, to int) ([]*model.Comment, error) {

	stmtOut, err := dbConn.Prepare(`SELECT  comments.id, users.Login_name, comments.content FROM comments
						INNER JOIN users ON comments.author_id = users.id
						WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)
	if err != nil {
		log.Panic(err)
	}
	var res []*model.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		var id, name, content string

		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, nil
		}

		c := &model.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil

}
