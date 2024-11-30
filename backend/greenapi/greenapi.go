package greenapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GreenAPI struct {
	BaseURL          string
	IDInstance       string
	APITokenInstance string
}

func NewGreenAPI(baseURL, idInstance, apiToken string) *GreenAPI {
	return &GreenAPI{
		BaseURL:          baseURL,
		IDInstance:       idInstance,
		APITokenInstance: apiToken,
	}
}

type AsyncResult struct {
	Response map[string]interface{}
	Error    error
}

func (api *GreenAPI) GetSettingsAsync() <-chan AsyncResult {
	resultChan := make(chan AsyncResult)
	go func() {
		defer close(resultChan)
		url := fmt.Sprintf("%s/waInstance%s/getSettings/%s", api.BaseURL, api.IDInstance, api.APITokenInstance)

		resp, err := http.Get(url)
		if err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)}
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)}
			return
		}

		resultChan <- AsyncResult{result, nil}
	}()
	return resultChan
}

func (api *GreenAPI) GetStateInstanceAsync() <-chan AsyncResult {
	resultChan := make(chan AsyncResult)
	go func() {
		defer close(resultChan)
		url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", api.BaseURL, api.IDInstance, api.APITokenInstance)

		resp, err := http.Get(url)
		if err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)}
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)}
			return
		}

		resultChan <- AsyncResult{result, nil}
	}()
	return resultChan
}

func (api *GreenAPI) SendMessageAsync(chatId, message string) <-chan AsyncResult {
	resultChan := make(chan AsyncResult)
	go func() {
		defer close(resultChan)
		url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", api.BaseURL, api.IDInstance, api.APITokenInstance)

		body := map[string]interface{}{
			"chatId":  chatId,
			"message": message,
		}

		jsonData, err := json.Marshal(body)
		if err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при маршаллинге данных: %w", err)}
			return
		}

		resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
		if err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)}
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)}
			return
		}

		resultChan <- AsyncResult{result, nil}
	}()
	return resultChan
}

func (api *GreenAPI) SendFileByUrlAsync(chatId, urlFile, fileName string) <-chan AsyncResult {
	resultChan := make(chan AsyncResult)
	go func() {
		defer close(resultChan)
		url := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", api.BaseURL, api.IDInstance, api.APITokenInstance)

		body := map[string]interface{}{
			"chatId":   chatId,
			"urlFile":  urlFile,
			"fileName": fileName,
		}

		jsonData, err := json.Marshal(body)
		if err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при маршаллинге данных: %w", err)}
			return
		}

		resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
		if err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)}
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resultChan <- AsyncResult{nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)}
			return
		}

		resultChan <- AsyncResult{result, nil}
	}()
	return resultChan
}
//sync methods
func (api *GreenAPI) GetSettings() (map[string]interface{}, error) {
    url := fmt.Sprintf("%s/waInstance%s/getSettings/%s", api.BaseURL, api.IDInstance, api.APITokenInstance)

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)
    }

    return result, nil
}

func (api *GreenAPI) GetStateInstance() (map[string]interface{}, error) {
    url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", api.BaseURL, api.IDInstance, api.APITokenInstance)

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)
    }

    return result, nil
}

func (api *GreenAPI) SendMessage(chatId, message string) (map[string]interface{}, error) {
    url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", api.BaseURL, api.IDInstance, api.APITokenInstance)

    body := map[string]interface{}{
        "chatId":  chatId,
        "message": message,
    }

    jsonData, err := json.Marshal(body)
    if err != nil {
        return nil, fmt.Errorf("ошибка при маршаллинге данных: %w", err)
    }

    resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
    if err != nil {
        return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)
    }

    return result, nil
}

func (api *GreenAPI) SendFileByUrl(chatId, urlFile, fileName string) (map[string]interface{}, error) {
    url := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", api.BaseURL, api.IDInstance, api.APITokenInstance)

    body := map[string]interface{}{
        "chatId":   chatId,
        "urlFile":  urlFile,
        "fileName": fileName,
    }

    jsonData, err := json.Marshal(body)
    if err != nil {
        return nil, fmt.Errorf("ошибка при маршаллинге данных: %w", err)
    }

    resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
    if err != nil {
        return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("ошибка при декодировании ответа: %w", err)
    }

    return result, nil
}