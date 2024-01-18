package query_agent

import (
	"context"

	"google.golang.org/grpc"
)

type UseCase struct {
	QueryAgentClient PromptServiceClient
}

// GetQuery implements PromptServiceClient.
func (u *UseCase) GetQuery(ctx context.Context, in *Prompt, opts ...grpc.CallOption) (*Response, error) {
	response, err := u.QueryAgentClient.GetQuery(ctx, in)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewUseCase(queryAgentClient PromptServiceClient) *UseCase {
	return &UseCase{
		QueryAgentClient: queryAgentClient,
	}
}
