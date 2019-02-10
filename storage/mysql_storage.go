package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlStorage struct {
	host   string
	user   string
	dbName string
	pass   string
}

func NewMysqlStorage(host string, user string, pass string, dbName string) *MysqlStorage {
	s := &MysqlStorage{}
	s.host = host
	s.user = user
	s.pass = pass
	s.dbName = dbName

	return s
}

func (s *MysqlStorage) SaveMessage(chatId int, userId int, text string) error {

	db := s.connection()
	defer db.Close()

	_, err := db.Exec(
		"insert into messages(chat_id, from_user, message)	 values (?, ?, ?)",
		chatId, userId, text)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *MysqlStorage) GetMessages(chatId int, messageFrom int, messageCount int) MessageList {

	var query string
	var placeholdersValues []interface{}

	//пока не нашел квери билдера сделаю так :(
	if messageFrom == 0 {
		query = "select * from messages where chat_id = ? order by id limit ?"
		placeholdersValues = append(placeholdersValues, chatId, messageCount)
	} else {
		query = "select * from messages where chat_id = ? and id < ? order by id limit ?"
		placeholdersValues = append(placeholdersValues, chatId, messageFrom, messageCount)
	}

	db := s.connection()
	defer db.Close()

	rsMessages, err := db.Query(query, placeholdersValues...)
	if err != nil {
		panic(err)
	}

	messages := MessageList{}
	for rsMessages.Next() {
		m := Message{}

		err := rsMessages.Scan(&m.ID, &m.Chat, &m.User, &m.Text, &m.CreatedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		messages = append(messages, m)
	}

	return messages
}

func (s *MysqlStorage) GetChats() ChatList {
	db := s.connection()
	defer db.Close()

	rsChats, err := db.Query("select * from chats")
	if err != nil {
		panic(err)
	}

	chats := ChatList{}
	for rsChats.Next() {
		c := Chat{}

		err := rsChats.Scan(&c.ID, &c.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		chats[c.ID] = c
	}

	return chats
}

func (s *MysqlStorage) connection() *sql.DB {
	db, err := sql.Open("mysql", s.user+":"+s.pass+"@/"+s.dbName)
	if err != nil {
		panic(err)
	}
	return db
}
