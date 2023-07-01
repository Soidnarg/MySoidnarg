package openai_local

import (
	"context"
	"fmt"
	"log"

	"github.com/Grandioso99/adminBot/utils"
	"github.com/sashabaranov/go-openai"
)

const (
	LENGHT_API_KEY = 51
)

var (
	TempoKey string
)

func NewOpenAI(key string) (n int, client *openai.Client) {
	n = 0
	if key == "" {
		key = utils.GetToken("openai")
	}
	if TempoKey != "" || len(key) < LENGHT_API_KEY {
		client = openai.NewClient(TempoKey)
	} else {
		client = openai.NewClient(key[:LENGHT_API_KEY])
		n = LENGHT_API_KEY
	}
	return
}

func AskToChatGPT(client *openai.Client, question string) (answer string, err error) {
	if client == nil {
		return "", fmt.Errorf("no client")
	}

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	}

	answer, err = superChatMode(client, req)

	return
}

func AskToGPT4(client *openai.Client, question string) (answer string, err error) {
	if client == nil {
		return "", fmt.Errorf("no client")
	}

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	}

	answer, err = superChatMode(client, req)

	/*
		req := openai.CompletionRequest{
			Model:  openai.GPT4,
			Prompt: question,
		}

		answer, err = normalChatMode(client, req)
	*/
	return
}

func AskToDaVinci(client *openai.Client, question string) (answer string, err error) {
	if client == nil {
		return "", fmt.Errorf("no client")
	}

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 500,
		Prompt:    question,
	}

	answer, err = normalChatMode(client, req)

	return
}

func AskToDumbAda(client *openai.Client, question string) (answer string, err error) {
	if client == nil {
		return "", fmt.Errorf("no client")
	}

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextAda001,
		MaxTokens: 500,
		Prompt:    question,
	}

	answer, err = normalChatMode(client, req)

	return
}

func AskToImageGPT(client *openai.Client, question string) (answer string, err error) {
	if client == nil {
		return "", fmt.Errorf("no client")
	}

	resp, err := client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         question,
			Size:           openai.CreateImageSize1024x1024,
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              1,
		},
	)
	return resp.Data[0].URL, err
}

func AskToAudioGPT(client *openai.Client, question string) (answer string, err error) {

	return
}

/*
	func streamChatMode(client *openai.Client, req openai.CompletionRequest) (answer string, err error) {
		stream, err := client.CreateCompletionStream(context.Background(), req)
		if err != nil {
			return "", fmt.Errorf("error streaming")
		}
		defer stream.Close()

		bench := time.Now()
		for time.Since(bench) < 10*time.Second {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				answer += "\nStream finished"
				return answer, err
			}

			if err != nil {
				return answer, err
			}

			answer += response.Choices[0].Text
			log.Println(response)
		}
		err = fmt.Errorf("timeout")
		return
	}
*/

func normalChatMode(client *openai.Client, req openai.CompletionRequest) (answer string, err error) {
	response, err := client.CreateCompletion(context.Background(), req)
	if err != nil {
		log.Panic(err)
		return "", fmt.Errorf("error streaming: %s", err.Error())
	}

	answer += response.Choices[0].Text
	log.Println(response)

	return
}

func superChatMode(client *openai.Client, req openai.ChatCompletionRequest) (answer string, err error) {
	response, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Panic(err)
		return "", fmt.Errorf("error superchat-gpt: %s", err.Error())
	}

	for _, choice := range response.Choices {
		answer += choice.Message.Content
	}
	log.Println(response)

	return
}
