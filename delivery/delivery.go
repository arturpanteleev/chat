package delivery

type RealTimeDelivery interface {
	SendMessage(chatId uint, userId uint, messageText string)
}
