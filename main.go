package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/outputparser"
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

	// customer_style := `American English
	// in a calm and respectful tone`
	// customer_email := `
	// Arrr, I be fuming that me blender lid
	// flew off and splattered me kitchen walls
	// with smoothie! And to make matters worse,
	// the warranty don't cover the cost of
	// cleaning up me kitchen. I need yer help
	// right now, matey!
	// `

	// promptTemplate := prompts.NewPromptTemplate(
	// 	`Translate the text
	// 	that is delimited by triple quotes
	// 	into a style that is {{.style}}.
	// 	text: """{{.text}}"""
	// 	`,
	// 	[]string{"style", "text"},
	// )

	// customerPrompt, _ := promptTemplate.Format(map[string]any{
	// 	"style": customer_style,
	// 	"text":  customer_email,
	// })

	// completion, err := GetCompletion(ctx, llm, customerPrompt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v", completion)

	// service_reply := `Hey there customer,
	// the warranty does not cover
	// cleaning expenses for your kitchen
	// because it's your fault that
	// you misused your blender
	// by forgetting to put the lid on before
	// starting the blender.
	// Tough luck! See ya!
	// `
	// service_style_pirate := `
	// a polite tone
	// that speaks in English Pirate
	// `

	// servicePrompt, _ := promptTemplate.Format(map[string]any{
	// 	"style": service_style_pirate,
	// 	"text":  service_reply,
	// })

	// completion, err = GetCompletion(ctx, llm, servicePrompt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v", completion)

	customer_review := `
	This leaf blower is pretty amazing.  It has four settings:
	candle blower, gentle breeze, windy city, and tornado.
	It arrived in two days, just in time for my wife's
	anniversary present.
	I think my wife liked it so much she was speechless.
	So far I've been the only one using it, and I've been
	using it every other morning to clear the leaves on our lawn.
	It's slightly more expensive than the other leaf blowers
	out there, but I think it's worth it for the extra features.
	`

	reviewTemplate := `
	For the following text, extract the following information:

	gift: Was the item purchased as a gift for someone else? 
	Answer True if yes, False if not or unknown.

	delivery_days: How many days did it take for the product 
	to arrive? If this information is not found, output -1.

	price_value: Extract any sentences about the value or price,
	and output them as a comma separated Python list.

	text: {{.text}}

	{{.format_instructions}}
	`

	promptTemplate := prompts.NewPromptTemplate(
		reviewTemplate,
		[]string{"text", "format_instructions"},
	)

	gift_schema := outputparser.ResponseSchema{
		Name: "gift",
		Description: `Was the item purchased
		as a gift for someone else?
		Answer True if yes,
		False if not or unknown.`,
	}

	delivery_days_schema := outputparser.ResponseSchema{
		Name: "delivery_days",
		Description: `How many days
		did it take for the product
		to arrive? If this
		information is not found,
		output -1.`,
	}

	price_value_schema := outputparser.ResponseSchema{
		Name: "price_value",
		Description: `Extract any
		sentences about the value or
		price, and output them as a
		comma separated go slice.`,
	}

	outputParser := outputparser.NewStructured([]outputparser.ResponseSchema{gift_schema, delivery_days_schema, price_value_schema})

	formatInstructions := outputParser.GetFormatInstructions()

	reviewPrompt, _ := promptTemplate.Format(map[string]any{
		"text":                customer_review,
		"format_instructions": formatInstructions,
	})

	completion, err := GetCompletion(ctx, llm, reviewPrompt)
	if err != nil {
		log.Fatal(err)
	}

	result, err := outputParser.Parse(completion)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", result)
}
