package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"greenapi/greenapi"
	"github.com/rs/cors"
	"bytes"
	"io"
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

func createHandler(action string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestData RequestData
		log.Printf("Получен запрос: %s %s", r.Method, r.URL.Path)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Ошибка при чтении тела запроса: %v", err)
			http.Error(w, "Ошибка при чтении тела запроса", http.StatusInternalServerError)
			return
		}

		log.Printf("Тело запроса: %s", string(body))

		r.Body = io.NopCloser(bytes.NewReader(body))

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			log.Printf("Ошибка при парсинге JSON: %v", err)
			http.Error(w, "Ошибка при парсинге JSON", http.StatusBadRequest)
			return
		}

		log.Printf("Данные запроса: %+v", requestData)

		apiInstance := greenapi.NewGreenAPI("https://1103.api.green-api.com", requestData.IdInstance, requestData.ApiTokenInstance)

		var resultChan <-chan greenapi.AsyncResult
		switch action {
		case "sendMessage":
			resultChan = apiInstance.SendMessageAsync(requestData.ChatId, requestData.Message)
		case "sendFile":
			resultChan = apiInstance.SendFileByUrlAsync(requestData.ChatId, requestData.UrlFile, requestData.FileName)
		case "getSettings":
			resultChan = apiInstance.GetSettingsAsync()
		case "getStateInstance":
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

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/send-message", createHandler("sendMessage"))
	apiMux.HandleFunc("/send-file", createHandler("sendFile"))
	apiMux.HandleFunc("/getSettings", createHandler("getSettings"))
	apiMux.HandleFunc("/getStateInstance", createHandler("getStateInstance"))

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		handler := c.Handler(apiMux)
		apiPort := "8881"
		log.Printf("API запущено на порту %s", apiPort)
		if err := http.ListenAndServe(":"+apiPort, handler); err != nil {
			log.Fatalf("Ошибка при запуске API-сервера: %v", err)
		}
	}()

	wg.Wait()
}


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

////////////////////////////////////запросы для будушего теста апи
// 1. Запрос getSettings:
// bash
// Copy code
// curl -X POST http://localhost:8881/getSettings \
//   -H "Content-Type: application/json" \
//   -d '{"idInstance": "1103157166", "apiTokenInstance": "cecd74e7d35849efa82c5a46f1fc543d618525b63cb84abda2"}'
// 2. Запрос getStateInstance:
// bash
// Copy code
// curl -X POST http://localhost:8881/getStateInstance \
//   -H "Content-Type: application/json" \
//   -d '{"idInstance": "1103157166", "apiTokenInstance": "cecd74e7d35849efa82c5a46f1fc543d618525b63cb84abda2"}'
// 3. Запрос на отправку сообщения (sendMessage):
// bash
// Copy code
// curl -X POST http://localhost:8881/send-message \
//   -H "Content-Type: application/json" \
//   -d '{
//     "chatId": "79687019003",
//     "message": "test with message",
//     "idInstance": "1103157166",
//     "apiTokenInstance": "cecd74e7d35849efa82c5a46f1fc543d618525b63cb84abda2"
//   }'
// 4. Запрос на отправку файла (sendFileByUrl):
// bash
// Copy code
// curl -X POST http://localhost:8881/send-file \
//   -H "Content-Type: application/json" \
//   -d '{
//     "chatId": "79687019003",
//     "urlFile": "https://sun9-37.userapi.com/impg/rWmGseJlnzKZjgbL9qNstDQyYpw7lo5S80Scgg/saRRNI0Crho.jpg?size=1170x1143&quality=95&sign=1b5e80e430c6338c90c9a4e06d2775b7&type=album",
//     "fileName": "saRRNI0Crho.jpg",
//     "idInstance": "1103157166",
//     "apiTokenInstance": "cecd74e7d35849efa82c5a46f1fc543d618525b63cb84abda2"
//   }'