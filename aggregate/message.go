package message

import (
	"fmt"

	"github.com/pix303/eventstore/event"
	"github.com/pix303/localemgmtES-commons/model"
)

const ()

type MessageAggregate struct {
	State []model.Message
}

func NewMessageAggregate() *MessageAggregate {
	ma := MessageAggregate{
		State: []model.Message{},
	}

	return &ma
}

func (ma *MessageAggregate) Apply(event event.StoreEvent) {
	fmt.Println("Apply something")
	fmt.Println("Ha Ha")
	fmt.Printf("dai %s", "dai")
}
func (ma *MessageAggregate) Reduce() {}
