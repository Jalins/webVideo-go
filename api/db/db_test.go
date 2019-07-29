package db

import (
	"log"
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {

	clearTables()
	m.Run()
	//clearTables()

}

func TestUserWorkFlow(t *testing.T) {

	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeletUser)
	t.Run("ReGet", testReGetUser)

}

func testAddUser(t *testing.T) {
	err := AddUserCradential("jalins", "123")

	if err != nil {
		log.Panic(err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("jalins")

	if err != nil {
		log.Panic(err)
	}

	log.Printf("the passworld is : %s\n", pwd)

}

func testDeletUser(t *testing.T) {
	err := DelectUser("jalins", "123")

	if err != nil {
		log.Panic(err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("jalins")

	if err != nil {
		log.Panic(err)
	}

	if pwd == "" {
		log.Println("the user is empty")
	}

}

func TestVideoInfoWorkFlow(t *testing.T) {

}

func TestCommontsWorkFlow(t *testing.T) {

}

func TestSessionsWorkFlow(t *testing.T) {

}
