package storage

/**
пока оставим мапппинг на json здесь, потом конечно же переделаем, т.к в реальном проекте
не допустимо использовать одну и ту же дто для уровня представления и хранилища данных
*/
type Message struct {
	ID        uint   `json:"id"`
	Text      string `json:"text"`
	Chat      uint   `json:"chat"`
	User      uint   `json:"user"`
	CreatedAt string `json:"created_at"`
}
type MessageList []Message

type Chat struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type ChatList map[uint]Chat
