package session

import (
	"api/db"
	"api/model"
	"api/utils"
	"log"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSeesionFromDB() {
	r, err := db.RetrieveAllSession()
	if err != nil {
		log.Panic(err)
	}

	r.Range(func(key, value interface{}) bool {
		session := value.(*model.SimpleSession)
		sessionMap.Store(key, session)
		return true
	})

}

func GenerateNewSeesionId(un string) string {

	uuid, err := utils.NewUUID()
	if err != nil {
		log.Panic(err)
	}

	ct := nowInMilli()
	ttl := ct * 30 * 60 * 1000 //session 超过30分钟

	ss := &model.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(uuid, ss)

	db.InserSession(uuid, ttl, un)

	return uuid

}

func IsSeesionExpired(sid string) (string, bool) {

	ss, ok := sessionMap.Load(sid)

	if ok {
		ct := nowInMilli()
		if ss.(*model.SimpleSession).TTL < ct {
			deleteExpireSession(sid)
			return "", true
		}

		return ss.(*model.SimpleSession).Username, false
	}
	return "", true

}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 100000
}

func deleteExpireSession(sid string) {
	sessionMap.Delete(sid)
	db.DeleteSession(sid)
}
