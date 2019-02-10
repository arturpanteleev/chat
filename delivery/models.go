package delivery

type CentrifugoMessage struct {
	Method string                  `json:"method"`
	Params CentrifugoMessageParams `json:"params"`
}

type CentrifugoMessageParams struct {
	Channel string                   `json:"channel"`
	Data    CentrifugoMessagePayload `json:"data"`
}

type CentrifugoMessagePayload struct {
	ChatId uint   `json:"chat_id"`
	UserId uint   `json:"user_id"`
	Text   string `json:"text"`
}
