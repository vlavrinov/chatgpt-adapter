package lmsys

import (
	"chatgpt-adapter/core/common"
	"chatgpt-adapter/core/gin/inter"
	"chatgpt-adapter/core/gin/model"
	"chatgpt-adapter/core/gin/response"
	"chatgpt-adapter/core/logger"
	"github.com/gin-gonic/gin"
	"github.com/iocgo/sdk/env"
)

var (
	Model = "lmsys"

	/*
		// lmsys 模型导出代码
		const lis = $0.querySelectorAll('li')
		let result = ''
		for (let index = 0, len = lis.length; index < len; index ++) {
			result += `"${lis[index].getAttribute('aria-label')}",\n`
		}
		console.log(`[${result}]`)
	*/
	modelSlice = []string{
		
  "chatgpt-4o-latest-20250326",
  "o3-2025-04-16",
  "o4-mini-2025-04-16",
  "gpt-4.1-2025-04-14",
  "gemini-2.5-pro-exp-03-25",
  "llama-4-maverick-03-26-experimental",
  "grok-3-preview-02-24",
  "claude-3-7-sonnet-20250219",
  "claude-3-7-sonnet-20250219-thinking-32k",
  "deepseek-v3-0324",
  "llama-4-maverick-17b-128e-instruct",
  "gpt-4.1-mini-2025-04-14",
  "gpt-4.1-nano-2025-04-14",
  "gemini-2.0-flash-thinking-exp-01-21",
  "gemini-2.0-flash-001",
  "gemini-2.0-flash-lite-preview-02-05",
  "gemma-3-27b-it",
  "gemma-3-12b-it",
  "gemma-3-4b-it",
  "deepseek-r1",
  "claude-3-5-sonnet-20241022",
  "o3-mini",
  "llama-3.3-70b-instruct",
  "gpt-4o-mini-2024-07-18",
  "gpt-4o-2024-11-20",
  "gpt-4o-2024-08-06",
  "gpt-4o-2024-05-13",
  "command-a-03-2025",
  "qwq-32b",
  "p2l-router-7b",
  "claude-3-5-haiku-20241022",
  "claude-3-5-sonnet-20240620",
  "doubao-1.5-pro-32k-250115",
  "doubao-1.5-vision-pro-32k-250115",
  "mistral-small-24b-instruct-2501",
  "phi-4",
  "amazon-nova-pro-v1.0",
  "amazon-nova-lite-v1.0",
  "amazon-nova-micro-v1.0",
  "cobalt-exp-beta-v3",
  "cobalt-exp-beta-v4",
  "qwen-max-2025-01-25",
  "qwen-plus-0125-exp",
  "qwen2.5-vl-32b-instruct",
  "qwen2.5-vl-72b-instruct",
  "gemini-1.5-pro-002",
  "gemini-1.5-flash-002",
  "gemini-1.5-flash-8b-001",
  "gemini-1.5-pro-001",
  "gemini-1.5-flash-001",
  "llama-3.1-405b-instruct-bf16",
  "llama-3.3-nemotron-49b-super-v1",
  "llama-3.1-nemotron-ultra-253b-v1",
  "llama-3.1-nemotron-70b-instruct",
  "llama-3.1-70b-instruct",
  "llama-3.1-8b-instruct",
  "hunyuan-standard-2025-02-10",
  "hunyuan-large-2025-02-10",
  "hunyuan-standard-vision-2024-12-31",
  "hunyuan-turbo-0110",
  "hunyuan-turbos-20250226",
  "mistral-large-2411",
  "pixtral-large-2411",
  "mistral-large-2407",
  "llama-3.1-nemotron-51b-instruct",
  "granite-3.1-8b-instruct",
  "granite-3.1-2b-instruct",
  "step-2-16k-exp-202412",
  "step-2-16k-202502",
  "step-1o-vision-32k-highres",
  "yi-lightning",
  "glm-4-plus",
  "glm-4-plus-0111",
  "jamba-1.5-large",
  "jamba-1.5-mini",
  "gemma-2-27b-it",
  "gemma-2-9b-it",
  "gemma-2-2b-it",
  "eureka-chatbot",
  "claude-3-haiku-20240307",
  "claude-3-sonnet-20240229",
  "claude-3-opus-20240229",
  "nemotron-4-340b",
  "llama-3-70b-instruct",
  "llama-3-8b-instruct",
  "qwen2.5-plus-1127",
  "qwen2.5-coder-32b-instruct",
  "qwen2.5-72b-instruct",
  "qwen-max-0919",
  "qwen-vl-max-1119",
  "qwen-vl-max-0809",
  "llama-3.1-tulu-3-70b",
  "olmo-2-0325-32b-instruct",
  "gpt-3.5-turbo-0125",
  "reka-core-20240904",
  "reka-flash-20240904",
  "c4ai-aya-expanse-32b",
  "c4ai-aya-expanse-8b",
  "c4ai-aya-vision-32b",
  "command-r-plus-08-2024",
  "command-r-08-2024",
  "codestral-2405",
  "mixtral-8x22b-instruct-v0.1",
  "mixtral-8x7b-instruct-v0.1",
  "pixtral-12b-2409",
  "ministral-8b-2410",

	}
)

type api struct {
	inter.BaseAdapter

	env *env.Environment
}

func (api *api) Match(ctx *gin.Context, model string) (ok bool, err error) {
	token := ctx.GetString("token")
	if len(model) <= 6 || model[:6] != Model+"/" {
		return
	}

	slice := api.env.GetStringSlice("lmsys.model")
	for _, mod := range append(slice, modelSlice...) {
		if model[6:] != mod {
			continue
		}

		password := api.env.GetString("server.password")
		if password != "" && password != token {
			err = response.UnauthorizedError
			return
		}

		ok = true
	}
	return
}

func (api *api) Models() (result []model.Model) {
	slice := api.env.GetStringSlice("lmsys.model")
	for _, mod := range append(slice, modelSlice...) {
		result = append(result, model.Model{
			Id:      "lmsys/" + mod,
			Object:  "model",
			Created: 1686935002,
			By:      "lmsys-adapter",
		})
	}
	return
}

func (api *api) ToolChoice(ctx *gin.Context) (ok bool, err error) {
	var (
		proxied    = api.env.GetString("server.proxied")
		completion = common.GetGinCompletion(ctx)
	)

	if toolChoice(ctx, api.env, proxied, completion) {
		ok = true
	}
	return
}

func (api *api) Completion(ctx *gin.Context) (err error) {
	var (
		proxied    = api.env.GetString("server.proxied")
		completion = common.GetGinCompletion(ctx)
	)

	completion.Model = completion.Model[6:]
	newMessages, err := mergeMessages(ctx, completion)
	if err != nil {
		response.Error(ctx, -1, err)
		return
	}
	ctx.Set(ginTokens, response.CalcTokens(newMessages))
	ch, err := fetch(ctx.Request.Context(), api.env, proxied, newMessages,
		options{
			model:       completion.Model,
			temperature: completion.Temperature,
			topP:        completion.TopP,
			maxTokens:   completion.MaxTokens,
		})
	if err != nil {
		logger.Error(err)
		return
	}

	content := waitResponse(ctx, ch, completion.Stream)
	if content == "" && response.NotResponse(ctx) {
		response.Error(ctx, -1, "EMPTY RESPONSE")
	}
	return
}
