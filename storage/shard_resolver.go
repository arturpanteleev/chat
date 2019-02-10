package storage

/* к сожалению руки до шардинга так и не дошли */
type ChatDb struct {
	serverName string
	dbName     string
	tableName  string
}

//тут будет сложная логика определения бд чата по его id пока заглушка
func findChatDb(chatId int) ChatDb {
	if chatId > 0 {
		return ChatDb{"localhost", "chat", "messages"}
	}

	return ChatDb{}
}

//функция будет выыватья при содании чата, и определять его virtualBacket, server, db , table
// и прочие параметры в зависимость от того как организуем шардинг
func defineChatDb( /*int chatId*/ ) {

}
