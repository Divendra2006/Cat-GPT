package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type FrontendRequest struct {
	Request string `json:"request"`
}

type APIResponse struct {
	Response string `json:"response"`
}

type APIInteract struct {
	Model  string `json:"model"`
	Query  string `json:"query"`
	Stream bool   `json:"stream"`
}

func Chat(w http.ResponseWriter, r *http.Request) {
	var query FrontendRequest
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := generateResponse(query.Request)
	json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func generateResponse(query string) (string, error) {
	var apiResponse APIResponse

	apiRequest, err := json.Marshal(APIInteract{Model: "llama3", Query: query, Stream: false})
	if err != nil {
		return "res : ", err
	}

	response, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(apiRequest))
	if err != nil {
		return "res: ", err
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&apiResponse)
	return apiResponse.Response, nil
}

// website load -> write request -> post from frontend using axios -> then post to api ->
