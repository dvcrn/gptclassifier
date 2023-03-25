# GPT File Classifier

Tool to 

- Classify: Classify the given content into one of the given options
- Name: Find a matching file name for the given content

GPT Classifier is a classifier tool written in Golang that can be used as both a Command-Line Interface (CLI) and an HTTP server. 

Fun fact: This entire repository was generated by GPT4

## Using the CLI Tool

To use the CLI tool, pass it the content to be classified and the options for classification as command line arguments.

### Classify: 

```
./gpt-classifier -organization "openai org" -apikey "openai api key" -action classify -content "this is regarding your credit card bill" -options "credit card,debit card"

credit card
```

### Name: 

```
./gpt-classifier -organization "openai org" -apikey "openai api key" -action name -content "this is regarding your credit card bill" -file_type "pdf" -example_name "2023-01-bill.pdf" -num_names 5

2022-05-credit-card-bill.pdf,credit-card-bill-05-2022.pdf,cc-bill-may-2022.pdf,2022-05-credit-card-statement.pdf,may-2022-credit-card-invoice.pdf
```

## Using the HTTP Server

### Classify:

To use the HTTP server, send a `POST` request to the `/classify` endpoint with a JSON payload containing the following fields:
- `content`: the content to be classified
- `options`: an array of options for classification
- `apikey`: the OpenAI API key
- `organization`: the OpenAI organization

```
curl -X POST -H "Content-Type: application/json" -d '{"content":"This is my credit card statement","options":["credit card","debit card","utility bill","phone bill"],"apikey":"xxx","organization":"xxx"}' http://localhost:8080/classify

{"success":true,"result":"credit card"}⏎
```

### Name:

To use the HTTP server, send a `POST` request to the `/name` endpoint with a JSON payload containing the following fields:
- `content`: the content to be classified
- `fileType`: file type of the content (pdf, txt, etc)
- `exampleName`: example how the file name should look like
- `numNames`: how many names to generate
- `apikey`: the OpenAI API key
- `organization`: the OpenAI organization

```
curl -X POST -H "Content-Type: application/json" -d '{"content":"This is my credit card statement","fileType": "pdf","exampleName":"2023-03-water-bill.pdf","apikey":"xxx","organization":"xxx"}' http://localhost:8080/name | json_pp

{
   "names" : [
      "credit-card-statement.pdf",
      "finance-summary.pdf",
      "monthly-finances.pdf",
      "account-summary.pdf",
      "billing-statement.pdf"
   ],
   "success" : true
}
```

## License

MIT