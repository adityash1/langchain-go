package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
)

func GetCompletion(ctx context.Context, llm *ollama.LLM, prompt string) (string, error) {
	completion, err := llm.Call(ctx, prompt,
		llms.WithTemperature(0.9),
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

	// ctx := context.Background()

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

	// customer_review := `
	// This leaf blower is pretty amazing.  It has four settings:
	// candle blower, gentle breeze, windy city, and tornado.
	// It arrived in two days, just in time for my wife's
	// anniversary present.
	// I think my wife liked it so much she was speechless.
	// So far I've been the only one using it, and I've been
	// using it every other morning to clear the leaves on our lawn.
	// It's slightly more expensive than the other leaf blowers
	// out there, but I think it's worth it for the extra features.
	// `

	// reviewTemplate := `
	// For the following text, extract the following information:

	// gift: Was the item purchased as a gift for someone else?
	// Answer True if yes, False if not or unknown.

	// delivery_days: How many days did it take for the product
	// to arrive? If this information is not found, output -1.

	// price_value: Extract any sentences about the value or price,
	// and output them as a comma separated Python list.

	// text: {{.text}}

	// {{.format_instructions}}
	// `

	// promptTemplate := prompts.NewPromptTemplate(
	// 	reviewTemplate,
	// 	[]string{"text", "format_instructions"},
	// )

	// gift_schema := outputparser.ResponseSchema{
	// 	Name: "gift",
	// 	Description: `Was the item purchased
	// 	as a gift for someone else?
	// 	Answer True if yes,
	// 	False if not or unknown.`,
	// }

	// delivery_days_schema := outputparser.ResponseSchema{
	// 	Name: "delivery_days",
	// 	Description: `How many days
	// 	did it take for the product
	// 	to arrive? If this
	// 	information is not found,
	// 	output -1.`,
	// }

	// price_value_schema := outputparser.ResponseSchema{
	// 	Name: "price_value",
	// 	Description: `Extract any
	// 	sentences about the value or
	// 	price, and output them as a
	// 	comma separated go slice.`,
	// }

	// outputParser := outputparser.NewStructured([]outputparser.ResponseSchema{gift_schema, delivery_days_schema, price_value_schema})

	// formatInstructions := outputParser.GetFormatInstructions()

	// reviewPrompt, _ := promptTemplate.Format(map[string]any{
	// 	"text":                customer_review,
	// 	"format_instructions": formatInstructions,
	// })

	// completion, err := GetCompletion(ctx, llm, reviewPrompt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result, err := outputParser.Parse(completion)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v", result)

	// bufferMemory := memory.NewConversationBuffer()

	// conversationChain := chains.NewConversation(llm, bufferMemory)

	// _, err = chains.Run(ctx, conversationChain, "Hi, my name is Andrew")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = chains.Run(ctx, conversationChain, "What is 1+1?")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = chains.Run(ctx, conversationChain, "What is my name?")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// messages, _ := bufferMemory.ChatHistory.Messages(ctx)

	// fmt.Printf("%+v", messages)

	// windowBufferMemory := memory.NewConversationWindowBuffer(1)

	// windowBufferMemory.SaveContext(ctx, map[string]any{
	// 	"input": "Hi",
	// }, map[string]any{
	// 	"output": "What's up",
	// })

	// windowBufferMemory.SaveContext(ctx, map[string]any{
	// 	"input": "Not much, just hanging",
	// }, map[string]any{
	// 	"output": "Cool",
	// })

	// conversationChain := chains.NewConversation(llm, windowBufferMemory)

	// _, err = chains.Run(ctx, conversationChain, "Hi, my name is Andrew")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = chains.Run(ctx, conversationChain, "What is 1+1?")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = chains.Run(ctx, conversationChain, "What is my name?")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v", windowBufferMemory.ChatHistory)

	// tokenBufferMemory := memory.NewConversationTokenBuffer(llm, 30)

	// tokenBufferMemory.SaveContext(ctx, map[string]any{
	// 	"input": "AI is what?!",
	// }, map[string]any{
	// 	"output": "Amazing!",
	// })

	// tokenBufferMemory.SaveContext(ctx, map[string]any{
	// 	"input": "Backpropagation is what?",
	// }, map[string]any{
	// 	"output": "Beautiful!",
	// })

	// tokenBufferMemory.SaveContext(ctx, map[string]any{
	// 	"input": "Chatbots are what?",
	// }, map[string]any{
	// 	"output": "Charming!",
	// })

	// fmt.Printf("%+v", tokenBufferMemory.ChatHistory)

	// prompt := prompts.NewPromptTemplate(
	// 	`What is the best name to describe \
	// a company that makes {{.product}}?
	// 	`,
	// 	[]string{"product"},
	// )

	// chain := chains.NewLLMChain(llm, prompt)

	// product := "Queen Size Sheet Set"
	// res, err := chains.Run(ctx, chain, product)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v", res)

	// firstPrompt := prompts.NewPromptTemplate(
	// 	`What is the best name to describe
	// 	a company that makes {{.product}}?. give only one name`,
	// 	[]string{"product"},
	// )

	// chainOne := chains.NewLLMChain(llm, firstPrompt)

	// secondPrompt := prompts.NewPromptTemplate(
	// 	`Write a 20 words description for the following
	// 	company:{{.company_name}}`,
	// 	[]string{"company_name"},
	// )

	// chainTwo := chains.NewLLMChain(llm, secondPrompt)

	// chainOne.OutputKey = "company_name"

	// simpleChain, err := chains.NewSequentialChain([]chains.Chain{chainOne, chainTwo}, []string{"product"}, []string{"text"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// product := "Queen Size Sheet Set"

	// res, err := chains.Run(ctx, simpleChain, product)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v", res)

	// type Data struct {
	// 	Product string `csv:"Product"`
	// 	Review  string `csv:"Review"`
	// }

	// file, err := os.Open("Data.csv")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer file.Close()

	// var data []Data
	// if err := gocsv.UnmarshalFile(file, &data); err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// firstPrompt := prompts.NewPromptTemplate(
	// 	`Translate the following review to english:

	// 	{{.Review}}`,
	// 	[]string{"Review"},
	// )
	// chainOne := chains.NewLLMChain(llm, firstPrompt)
	// chainOne.OutputKey = "English_Review"

	// secondPrompt := prompts.NewPromptTemplate(
	// 	`Can you summarize the following review in 1 sentence:

	// 	{{.English_Review}}`,
	// 	[]string{"English_Review"},
	// )
	// chainTwo := chains.NewLLMChain(llm, secondPrompt)
	// chainTwo.OutputKey = "summary"

	// thirdPrompt := prompts.NewPromptTemplate(
	// 	`What language is the following review:

	// 	{{.Review}}`,
	// 	[]string{"Review"},
	// )
	// chainThird := chains.NewLLMChain(llm, thirdPrompt)
	// chainThird.OutputKey = "language"

	// fourthPrompt := prompts.NewPromptTemplate(
	// 	`Write a follow up response to the following
	// 	summary in the specified language:

	// 	Summary: {{.summary}}`,

	// Language: {{.language}}`,
	// []string{"summary", "language"},
	// 	[]string{"summary"},
	// )
	// chainFourth := chains.NewLLMChain(llm, fourthPrompt)
	// chainFourth.OutputKey = "followup_message"

	// seqChain, err := chains.NewSequentialChain([]chains.Chain{chainOne, chainTwo, chainFourth}, []string{"Review"}, []string{"followup_message"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := chains.Predict(ctx, seqChain, map[string]any{
	// 	"Review": data[5].Review,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v", res)

	physics_template := `You are a very smart physics professor.
	You are great at answering questions about physics in a concise
	and easy to understand manner.
	When you don't know the answer to a question you admit
	that you don't know.
	
	Here is a question:
	{{.input}}`

	math_template := `You are a very good mathematician.
	You are great at answering math questions.
	You are so good because you are able to break down
	hard problems into their component parts, 
	answer the component parts, and then put them together
	to answer the broader question.
	
	Here is a question:
	{{.input}}`

	history_template := `You are a very good historian.
	You have an excellent knowledge of and understanding of people,
	events and contexts from a range of historical periods.
	You have the ability to think, reflect, debate, discuss and
	evaluate the past. You have a respect for historical evidence
	and the ability to make use of it to support your explanations
	and judgements.
	
	Here is a question:
	{{.input}}`

	computerscience_template := `You are a successful computer scientist.
	You have a passion for creativity, collaboration,
	forward-thinking, confidence, strong problem-solving capabilities,
	understanding of theories and algorithms, and excellent communication
	skills. You are great at answering coding questions.
	You are so good because you know how to solve a problem by
	describing the solution in imperative steps
	that a machine can easily interpret and you know how to
	choose a solution that has a good balance between
	time complexity and space complexity.
	
	Here is a question:
	{{.input}}`

	promptInfos := []map[string]string{
		{
			"name":            "physics",
			"description":     "Good for answering questions about physics",
			"prompt_template": physics_template,
		},
		{
			"name":            "math",
			"description":     "Good for answering math questions",
			"prompt_template": math_template,
		},
		{
			"name":            "History",
			"description":     "Good for answering history questions",
			"prompt_template": history_template,
		},
		{
			"name":            "computer science",
			"description":     "Good for answering computer science questions",
			"prompt_template": computerscience_template,
		},
	}

	destinationChains := make(map[string]any)
	var destinations []string
	var destinationsStr string

	for _, pInfo := range promptInfos {
		prompt := prompts.NewPromptTemplate(
			pInfo["prompt_template"],
			[]string{"input"},
		)
		chain := chains.NewLLMChain(llm, prompt)

		destinationChains[pInfo["name"]] = chain

		destinations = append(destinations, fmt.Sprintf("%s: %s", pInfo["name"], pInfo["description"]))
	}

	destinationsStr = strings.Join(destinations, "\n")

	fmt.Println(destinationsStr)

	defaultPrompt := prompts.NewPromptTemplate(
		"",
		[]string{"input"},
	)
	defaultChain := chains.NewLLMChain(llm, defaultPrompt)

	const MULTI_PROMPT_ROUTER_TEMPLATE = `Given a raw text input to a 
	language model select the model prompt best suited for the input. 
	You will be given the names of the available prompts and a 
	description of what the prompt is best suited for. 
	You may also revise the original input if you think that revising
	it will ultimately lead to a better response from the language model.

	<< FORMATTING >>
	Return a markdown code snippet with a JSON object formatted to look like:
	` + "```json" + `
	{
		"destination": string // "DEFAULT" or name of the prompt to use in {{.destinations}}
		"next_inputs": string // a potentially modified version of the original input
	}
	` + "```" + `

	REMEMBER: The value of "destination" MUST match one of 
	the candidate prompts listed below.
	If "destination" does not fit any of the specified prompts, set it to "DEFAULT."
	REMEMBER: "next_inputs" can just be the original input 
	if you don't think any modifications are needed.

	<< CANDIDATE PROMPTS >>
	{{.destinations}}

	<< INPUT >>
	{{.input}}

	<< OUTPUT (remember to include the ` + "```json" + `)>>`

	// chains.NewConstitutional(llm, destinationChains, )
}
