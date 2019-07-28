package db

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"webVideo-go/model"
)

// 写入session
func InserSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmt, err := dbConn.Prepare("INSERT INTO sessions (sessions_id, TTl, login_name) VALUES (?, ?, ?)")
	if err != nil {
		log.Panic(err)
	}

	_, err = stmt.Exec(sid, ttlstr, uname)
	if err != nil {
		log.Panic(err)
	}

	defer stmt.Close()
	return nil
}

// 获取session

func RetrieveSession(sid string) (*model.SimpleSession, error) {
	session := &model.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE sessions_id=?")
	if err != nil {
		log.Panic(err)
	}

	var ttl string
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		session.TTL = res
		session.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()

	return session, nil

}

// 获取所有的session
func RetrieveAllSession() (*sync.Map, error) {
	var (
		err     error
		stmtOut *sql.Stmt
		rows    *sql.Rows
		ttl     int64
	)
	m := &sync.Map{}
	stmtOut, err = dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Panic(err)
	}

	rows, err = stmtOut.Query()
	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		var (
			id         string
			ttlstr     string
			login_name string
		)

		err = rows.Scan(&id, &ttlstr, &login_name)
		if err != nil {
			log.Printf("retrieve sessions error: %s", err)
			break
		}

		ttl, err = strconv.ParseInt(ttlstr, 10, 64)
		if err == nil {
			session := &model.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, session)
			log.Printf("session id: %s, ttl: %s", id, session.TTL)
		}

	}

	return m, nil

}

// 删除session

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE sessions_id=? ")
	if err != nil {
		log.Panic(err)
	}

	_, err = stmtOut.Query(sid)

	if err != nil {
		log.Panic(err)
	}
	return nil

}
