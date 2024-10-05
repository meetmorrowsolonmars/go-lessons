package card

import domaincard "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/card"

type CreateResponse struct {
}

type DefaultResponse struct {
	domaincard.Card
}
