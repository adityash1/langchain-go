package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
)

func GetCompletion(ctx context.Context, llm *ollama.LLM, prompt string) (string, error) {
	completion, err := llm.Call(ctx, prompt,
		llms.WithTemperature(0.8),
		// llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		// 	fmt.Print(string(chunk))
		// 	return nil
		// }),
	)

	return completion, err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model := os.Getenv("MODEL")

	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	customer_style := "American English in a calm and respectful tone"
	customer_email := "Arrr, I be fuming that me blender lid flew off and splattered me kitchen walls with smoothie! And to make matters worse, the warranty don't cover the cost of cleaning up me kitchen. I need yer help right now, matey!"

	promptTemplate := prompts.NewPromptTemplate(
		"Translate the text that is delimited by triple backticks into a style that is {{.style}}?\n```text: {{.text}}```",
		[]string{"style", "text"},
	)

	customerPrompt, _ := promptTemplate.Format(map[string]any{
		"style": customer_style,
		"text":  customer_email,
	})

	completion, err := GetCompletion(ctx, llm, customerPrompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", completion)

	service_reply := "Hey there customer, the warranty does not cover cleaning expenses for your kitchen because it's your fault that you misused your blender by forgetting to put the lid on before starting the blender. Tough luck! See ya!"
	service_style_pirate := "a polite tone that speaks in English Pirate"

	servicePrompt, _ := promptTemplate.Format(map[string]any{
		"style": service_style_pirate,
		"text":  service_reply,
	})

	completion, err = GetCompletion(ctx, llm, servicePrompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", completion)
}
