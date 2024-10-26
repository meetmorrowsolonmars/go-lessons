package balance

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/meetmorrowsolonmars/go-lessons/lesson-3/internal/pb/api/v1/balance"
)

func (i *Implementation) GetOptions(_ context.Context, request *balance.GetOptionsRequest) (*balance.GetOptionsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %s", err)
	}

	switch request.GetOptionType() {
	case balance.OptionType_ACCOUNT, balance.OptionType_TRANSFER:
		return &balance.GetOptionsResponse{
			Options: []string{"create", "read"},
		}, nil
	default:
	}

	return &balance.GetOptionsResponse{
		Options: []string{"create", "read", "update", "delete"},
	}, nil
}
