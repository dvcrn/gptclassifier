package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dvcrn/gptclassifier/internal/command"
	"github.com/dvcrn/gptclassifier/internal/utils"
	"github.com/dvcrn/gptclassifier/pkg/openai"
)

var (
	content      string
	filepath     string
	apiKey       string
	organization string
	action       string
	options      string
	numNames     int
	fileType     string
	fileName     string
)

func init() {
	flag.StringVar(&content, "content", "", "Content to process")
	flag.StringVar(&filepath, "filepath", "", "Path to the file with content")
	flag.StringVar(&apiKey, "apikey", os.Getenv("OPENAI_API_KEY"), "OpenAI API key")
	flag.StringVar(&organization, "organization", os.Getenv("OPENAI_API_SECRET"), "OpenAI organization")
	flag.StringVar(&action, "action", "", "Action to perform (classify or name)")
	flag.StringVar(&options, "options", "", "Comma-separated list of options for classification")
	flag.IntVar(&numNames, "num_names", 5, "Number of names to generate")
	flag.StringVar(&fileType, "file_type", "", "File type")
	flag.StringVar(&fileName, "example_name", "", "Example of file name")
}

func main() {
	flag.Parse()

	if apiKey == "" || organization == "" {
		fmt.Println("Error: API key or organization not provided.")
		return
	}

	if content == "" && filepath == "" {
		fmt.Println("Error: Either content or filepath must be provided.")
		return
	}

	if filepath != "" {
		var err error
		content, err = utils.ReadFileContent(filepath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}

	client := openai.New(apiKey)

	switch action {
	case "classify":
		if options == "" {
			fmt.Println("Error: Options must be provided for classification.")
			return
		}
		optionSlice := strings.Split(options, ",")
		result, err := command.Classify(content, client, optionSlice)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println(result)
	case "name":
		if fileType == "" {
			fmt.Println("Error: File type must be provided for name generation.")
			return
		}

		result, err := command.Name(content, client, fileType, numNames, fileName)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(strings.Join(result, ","))
	default:
		fmt.Println("Usage: ./your-cli [-content OR -filepath] [-apikey OR use env variable OPENAI_API_KEY] [-organization OR use env variable OPENAI_API_SECRET] -action [classify OR name] [-options (comma-separated) for classify]")
	}
}
