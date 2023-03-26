package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dvcrn/gptclassifier/internal/command"
	"github.com/dvcrn/gptclassifier/pkg/openai"
)

type classifyRequest struct {
	Content      string   `json:"content"`
	Options      []string `json:"options"`
	APIKey       string   `json:"apikey"`
	Organization string   `json:"organization"`
}

type nameRequest struct {
	Content      string   `json:"content"`
	APIKey       string   `json:"apikey"`
	Organization string   `json:"organization"`
	ExampleNames []string `json:"exampleNames,omitempty"`
	FileType     string   `json:"fileType,omitempty"`
	NumNames     int      `json:"numNames,omitempty"`
}

type classificationResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Result  string `json:"result,omitempty"`
}

type nameResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error,omitempty"`
	Names   []string `json:"names"`
}

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
var OPENAI_ORG = os.Getenv("OPENAI_ORGANIZATION")

func main() {
	http.HandleFunc("/classify", classifyHandler)
	http.HandleFunc("/name", nameHandler)

	// read port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting server on port", port)
	http.ListenAndServe(":"+port, nil)
}

func classifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	req, err := parseRequest(&classifyRequest{}, w, r)
	if err != nil {
		return
	}

	classifyReq := req.(*classifyRequest)

	apiKey := classifyReq.APIKey
	if apiKey == "" {
		apiKey = OPENAI_API_KEY
	}

	org := classifyReq.Organization
	if org == "" {
		org = OPENAI_ORG
	}

	if apiKey == "" || org == "" {
		writeJSONResponse(w, classificationResponse{Success: false, Error: "API key or organization not provided"})
		return
	}

	if len(classifyReq.Options) == 0 {
		writeJSONResponse(w, classificationResponse{Success: false, Error: "Options must be provided for classification"})
		return
	}

	client := openai.New(apiKey)
	result, err := command.Classify(classifyReq.Content, client, classifyReq.Options)
	if err != nil {
		writeJSONResponse(w, classificationResponse{Success: false, Error: err.Error()})
		return
	}

	writeJSONResponse(w, classificationResponse{Success: true, Result: result})
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	req, err := parseRequest(&nameRequest{}, w, r)
	if err != nil {
		return
	}

	nameReq := req.(*nameRequest)

	apiKey := nameReq.APIKey
	if apiKey == "" {
		apiKey = OPENAI_API_KEY
	}

	org := nameReq.Organization
	if org == "" {
		org = OPENAI_ORG
	}

	if apiKey == "" || org == "" {
		writeJSONResponse(w, classificationResponse{Success: false, Error: "API key or organization not provided"})
		return
	}

	if nameReq.NumNames == 0 {
		nameReq.NumNames = 5
	}

	if nameReq.FileType == "" {
		writeJSONResponse(w, classificationResponse{Success: false, Error: "fileType must be provided for name generation"})
		return
	}

	client := openai.New(apiKey)
	result, err := command.Name(nameReq.Content, client, nameReq.FileType, nameReq.NumNames, nameReq.ExampleNames)
	if err != nil {
		writeJSONResponse(w, nameResponse{Success: false, Error: err.Error()})
		return
	}

	writeJSONResponse(w, nameResponse{Success: true, Names: result})
}

func parseRequest(target interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return nil, err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return nil, err
	}

	return target, nil
}

func writeJSONResponse(w http.ResponseWriter, resp interface{}) {
	respJSON, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}
