package balance

import desc "github.com/meetmorrowsolonmars/go-lessons/lesson-3/internal/pb/api/v1/balance"

type Implementation struct {
	desc.UnimplementedBalanceServiceServer
}

func NewBalanceServerImplementation() *Implementation {
	return &Implementation{}
}
