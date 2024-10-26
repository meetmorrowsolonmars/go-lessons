package main

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/meetmorrowsolonmars/go-lessons/lesson-3/internal/pb/api/v1/balance"
)

func main() {
	acc := balance.Account{
		Id:        "account_id_1",
		UserId:    "user_id_1",
		Email:     "test@gmail.com",
		CreatedAt: timestamppb.New(time.Now()),
	}

	message, _ := protojson.MarshalOptions{Multiline: true, Indent: "\t"}.Marshal(&acc)

	fmt.Println(string(message))
}
