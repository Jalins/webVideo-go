package db

import (
	"log"
	"time"
	"webVideo-go/model"
	"webVideo-go/utils"

	_ "github.com/go-sql-driver/mysql"
)

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
