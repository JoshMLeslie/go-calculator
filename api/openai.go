package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var (
	client *openai.Client
)

var queryDisclaimer string = "Given the following expression, convert it without explanation or otherwise fanfare to reverse polish notation: "

func InitOpenAiAPI(
	key string,
	org string,
	proj string,
) {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println(err)
		panic("key not loaded")
	}

	if key == "" {
		key = os.Getenv("OPEN_AI_KEY")
	}
	if org == "" {
		org = os.Getenv("OPEN_AI_ORG")
	}
	if proj == "" {
		proj = os.Getenv("OPEN_AI_PROJ")
	}

	client = openai.NewClient(
		option.WithMiddleware(logger),
		option.WithAPIKey(key),
		option.WithOrganization(org),
		option.WithProject(proj),
		option.WithMaxRetries(0),
	)

}

func logger(req *http.Request, next option.MiddlewareNext) (res *http.Response, err error) {
	start := time.Now()
	fmt.Printf("REQ %v \n", req)

	res, err = next(req)

	end := time.Now()

	fmt.Printf("RES %v \n ERR %v \n DIFF %v \n", res, err, end.Sub(start))
	return res, err
}

func CallOpenAI(query string) (string, error) {
	query = queryDisclaimer + query

	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(query),
			}),
			Model: openai.F(openai.ChatModelGPT3_5Turbo),
		})
	if err != nil {
		return "", err
	}

	return chatCompletion.JSON.RawJSON(), nil
}
