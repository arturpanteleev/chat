package api

import (
	"chat/delivery"
	"chat/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

const DEFAULT_MESSAGES_COUNT = 100
const CURRENT_USER = 1

/*
в этот объект инкапсулируем всю работу с api и вызов разных инфраструктурных штук
по сути будет выполнять роль контроллера и application сервиса
*/
type API struct {
	port     string
	echo     *echo.Echo
	storage  storage.Storage
	delivery delivery.RealTimeDelivery
}

func New(port string, echo *echo.Echo, storage storage.Storage, delivery delivery.RealTimeDelivery) *API {
	api := &API{}
	api.port = port
	api.echo = echo
	api.storage = storage
	api.delivery = delivery

	apiGroup := api.echo.Group("/api")

	// пока так, потом настрою
	apiGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	//пока захуярю все гетами чтобы тетсить через браузер
	apiGroup.GET("/read/", api.readMessages)
	apiGroup.GET("/sendMessage/", api.sendMessage)
	apiGroup.GET("/messages/", api.getMessages)
	apiGroup.GET("/chats/", api.getChats)

	return api
}

func (api *API) Start() {
	err := api.echo.Start(api.port)
	if err != nil {
		panic(err)
	}
}

/*
Получааем сообщения пачками
Идея такая что пользователь будет передавать ID самого старого сообщения чата, которое у него хранится
и получать следующую пачку
*/
func (api *API) getMessages(c echo.Context) error {
	chatId, err := strconv.Atoi(c.QueryParam("chat_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &DefaultResponse{
			Success: false,
			Message: "Не корректно указан chat_id",
		})
	}

	fromMessage, err := strconv.Atoi(c.QueryParam("oldest_message"))
	if err != nil {
		fromMessage = 0
	}

	messageCount, _ := strconv.Atoi(c.QueryParam("message_count"))
	if messageCount <= 0 {
		messageCount = DEFAULT_MESSAGES_COUNT
	}

	messages := api.storage.GetMessages(chatId, fromMessage, messageCount)
	return c.JSON(http.StatusOK, messages)
}

func (api *API) sendMessage(c echo.Context) error {

	chatId, err := strconv.Atoi(c.QueryParam("chat_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &DefaultResponse{
			Success: false,
			Message: "Не корректно указан chat_id",
		})
	}
	if chatId < 1 {
		return c.JSON(http.StatusBadRequest, &DefaultResponse{
			Success: false,
			Message: "Не корректно указан chat_id",
		})
	}

	message := c.QueryParam("message")
	if len(message) < 1 {
		return c.JSON(http.StatusBadRequest, &DefaultResponse{
			Success: false,
			Message: "Сообщение не может быть пустым",
		})
	}

	err = api.storage.SaveMessage(chatId, CURRENT_USER, message)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &DefaultResponse{
			Success: false,
			Message: "Ошибка сохранения сообщения",
		})
	}

	//тут не фатально если ошибка
	api.delivery.SendMessage(uint(chatId), uint(CURRENT_USER), message)
	return c.JSON(http.StatusOK, &DefaultResponse{
		Success: true,
		Message: "Сообщение добавлено",
	})
}

func (api *API) getChats(c echo.Context) error {
	chats := api.storage.GetChats()
	return c.JSON(http.StatusOK, chats)
}

/**
тут будем сдвигать курсор прочитанных сообщений
*/
func (api *API) readMessages(c echo.Context) error {
	return nil
}
