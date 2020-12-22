package broker

import "github.com/layer5io/meshsync/internal/model"

var (
	List   ObjectType = "list"
	Single ObjectType = "single"
)

type ObjectType string

type Message struct {
	Type   ObjectType
	Object model.Object
}

type PublishInterface interface {
	Publish(string, *Message) error
	PublishWithCallback(string, string, *Message) error
}

type SubscribeInterface interface {
	Subscribe(string, string, *Message) error
	SubscribeWithHandler(string, string, *Message) error
	SubscribeWithChannel(string, string, chan *Message) error
}

type Handler interface {
	PublishInterface
	SubscribeInterface
}
