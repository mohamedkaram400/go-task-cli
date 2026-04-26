package model

type Command struct {
	Action string
	Payload interface{}
	Result chan interface{}
}