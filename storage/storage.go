package storage

type Storage interface {
	SaveMessage(chatId int, userId int, text string) error
	GetMessages(chatId int, messageFrom int, messageCount int) MessageList
	GetChats() ChatList
}
