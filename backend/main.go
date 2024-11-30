package main

import (
	"backend/greenapi"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/rs/cors"
)

type RequestData struct {
	ChatId           string `json:"chatId"`
	Message          string `json:"message"`
	UrlFile          string `json:"urlFile"`
	FileName         string `json:"fileName"`
	Caption          string `json:"caption"`
	IdInstance       string `json:"idInstance"`
	ApiTokenInstance string `json:"apiTokenInstance"`
}

var api *greenapi.GreenAPI

type HandlerFunc func(w http.ResponseWriter, r *http.Request, api *greenapi.GreenAPI)

func createHandler(action string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, api *greenapi.GreenAPI) {
		var requestData RequestData
		fmt.Println("Получен запрос:", r.Method, r.URL)

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Данные запроса:", requestData)

		apiInstance := greenapi.NewGreenAPI("https://1103.api.green-api.com", requestData.IdInstance, requestData.ApiTokenInstance)

		var resultChan <-chan greenapi.AsyncResult


		switch action {
		case "sendMessage":
			fmt.Println("Отправляем запрос на API (SendMessage):", requestData.ChatId, requestData.Message)
			resultChan = apiInstance.SendMessageAsync(requestData.ChatId, requestData.Message)
		case "sendFile":
			fmt.Println("Отправляем запрос на API (SendFile):", requestData.ChatId, requestData.UrlFile, requestData.FileName)
			resultChan = apiInstance.SendFileByUrlAsync(requestData.ChatId, requestData.UrlFile, requestData.FileName)
		case "getSettings":
			fmt.Println("Запрашиваем настройки экземпляра")
			resultChan = apiInstance.GetSettingsAsync()
		case "getStateInstance":
			fmt.Println("Запрашиваем состояние экземпляра")
			resultChan = apiInstance.GetStateInstanceAsync()
		default:
			http.Error(w, "Unknown action", http.StatusBadRequest)
			return
		}

		result := <-resultChan

		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Ответ от API:", result.Response)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result.Response)
	}
}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	http.HandleFunc("/send-message", func(w http.ResponseWriter, r *http.Request) {
		createHandler("sendMessage")(w, r, api)
	})
	http.HandleFunc("/send-file", func(w http.ResponseWriter, r *http.Request) {
		createHandler("sendFile")(w, r, api)
	})
	http.HandleFunc("/getSettings", func(w http.ResponseWriter, r *http.Request) {
		createHandler("getSettings")(w, r, api)
	})
	http.HandleFunc("/getStateInstance", func(w http.ResponseWriter, r *http.Request) {
		createHandler("getStateInstance")(w, r, api)
	})

	handler := c.Handler(http.DefaultServeMux)

	fmt.Println("API сервер запущен на порту 8881")
	log.Fatal(http.ListenAndServe(":8881", handler))
}



// func main() {
// 	//данные из личного кабинета, нужны для всех методовЫ
// 	baseURL := "https://1103.api.green-api.com"
// 	idInstance := "1103157166"
// 	apiTokenInstance := "cecd74e7d35849efa82c5a46f1fc543d618525b63cb84abda2"
// 	//нужно для 3 метода
// 	api := NewGreenAPI(baseURL, idInstance, apiTokenInstance)
// 	chatId := "79687019003@c.us"
// 	message := "I use Green-API to send this message to you!"
// 	//нужно для 4 метода
// 	urlFile := "https://sun9-37.userapi.com/impg/rWmGseJlnzKZjgbL9qNstDQyYpw7lo5S80Scgg/saRRNI0Crho.jpg?size=1170x1143&quality=95&sign=1b5e80e430c6338c90c9a4e06d2775b7&type=album"
// 	fileName := "saRRNI0Crho.jpg"
// 	caption := "Here's the file!"
// 	///////////////////4 метод тест Отправка файлов
// 	file, err := api.SendFileByUrl(chatId, urlFile, fileName, caption)
// 	if err != nil {
// 		log.Fatalf("Ошибка при отправке файла: %v", err)
// 	}

// 	fmt.Println("\nОтвет от отправки файла:")
// 	for key, value := range file {
// 		fmt.Printf("%s: %v\n", key, value)
// 	}
// 	///////////////////2 метод тест Отправка сообщений
// 	response, err := api.SendMessage(chatId, message)
// 	if err != nil {
// 		log.Fatalf("Ошибка при отправке сообщения: %v", err)
// 	}

// 	fmt.Println("\nОтвет от отправки сообщения:")
// 	for key, value := range response {
// 		fmt.Printf("%s: %v\n", key, value)
// 	}
// 	///////////////////2 метод тест Получения статуса Инстанса
// 	state, err := api.GetStateInstance()
// 	if err != nil {
// 		log.Fatalf("Ошибка при вызове GetStateInstance: %v", err)
// 	}
// 	fmt.Println("\nСостояние инстанса:")
// 	for key, value := range state {
// 		fmt.Printf("%s: %v\n", key, value)
// 	}
// 	/////////////////1 метод тест Получение настроек инстанса
// 	settings, err := api.GetSettings()
// 	if err != nil {
// 		log.Fatalf("Ошибка при вызове GetSettings: %v", err)
// 	}

// 	fmt.Println("Настройки инстанса:")
// 	for key, value := range settings {
// 		fmt.Printf("%s: %v\n", key, value)
// 	}
// }
