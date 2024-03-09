package client

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/misikdmytro/granny-grams/model"
)

type OpenAIClient interface {
	GenerateImage(ctx context.Context, prompt string, n uint8, size string) (model.ImageGenerationResponse, error)
}

type openAIClient struct {
	*resty.Client
}

type fakeOpenAIClient struct{}

// GenerateImage implements OpenAIClient.
func (f *fakeOpenAIClient) GenerateImage(context.Context, string, uint8, string) (model.ImageGenerationResponse, error) {
	// This is a fake implementation that returns a single image URL.
	return model.ImageGenerationResponse{
		Data: []model.ImageGenerationData{
			{
				URL: "https://oaidalleapiprodscus.blob.core.windows.net/private/org-fUuDr1cMLkCLclVZqq8JDgBZ/user-3eQQri1lXrRJq8yXf7L1EGPw/img-KLUaGtUl44ar2BmwtujmtOMP.png?st=2024-03-09T11%3A07%3A35Z&se=2024-03-09T13%3A07%3A35Z&sp=r&sv=2021-08-06&sr=b&rscd=inline&rsct=image/png&skoid=6aaadede-4fb3-4698-a8f6-684d7786b067&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2024-03-08T18%3A22%3A24Z&ske=2024-03-09T18%3A22%3A24Z&sks=b&skv=2021-08-06&sig=xF7uOxAUcG1Tgdx%2BizBZb%2Bx4rYz8zyc8dKsg3lTjVHM%3D",
			},
		},
	}, nil
}

// GenerateImage implements OpenAIClient.
func (o *openAIClient) GenerateImage(ctx context.Context, prompt string, n uint8, size string) (model.ImageGenerationResponse, error) {
	var success model.ImageGenerationResponse
	var fail model.OpenAIErrorResponse

	resp, err := o.R().
		SetContext(ctx).
		SetBody(model.ImageGenerationRequest{
			Model:  "dall-e-3",
			Prompt: prompt,
			N:      n,
			Size:   size,
		}).
		SetResult(&success).
		SetError(&fail).
		Post("/v1/images/generations")

	if err != nil {
		return model.ImageGenerationResponse{}, err
	}

	if resp.IsError() {
		return model.ImageGenerationResponse{}, &model.RESTClientError{
			Message:  "OpenAI error",
			Response: fail,
		}
	}

	return success, nil
}

func NewOpenAIClient(baseURL string, token string) OpenAIClient {
	client := resty.New()
	client.SetBaseURL(baseURL)
	client.SetAuthScheme("Bearer")
	client.SetAuthToken(token)

	return &openAIClient{client}
}

func NewFakeOpenAIClient() OpenAIClient {
	return &fakeOpenAIClient{}
}
