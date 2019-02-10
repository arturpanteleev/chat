package delivery

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type CentrifugoDelivery struct {
	url string
}

func NewCentrifugoDelivery(apiUrl string) *CentrifugoDelivery {
	//return &CentrifugoDelivery{"http://localhost:8000/api"}
	return &CentrifugoDelivery{apiUrl}
}

func (cfDelivery *CentrifugoDelivery) SendMessage(chatId uint, userId uint, messageText string) {

	message := CentrifugoMessage{
		"publish",
		CentrifugoMessageParams{
			"chat" + strconv.Itoa(int(chatId)),
			CentrifugoMessagePayload{
				chatId,
				userId,
				messageText,
			},
		},
	}

	jsonMsgStr, _ := json.Marshal(message)

	req, err := http.NewRequest("POST", cfDelivery.url, bytes.NewBuffer(jsonMsgStr))

	//это очень хуево и не по христиански, но пока оставлю так, вернусь когда буду "прилизывать проект"
	req.Header.Set("Authorization", "apikey 5064b860-c24a-4b2e-b56c-72ac5b392ef6")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	println(err)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
