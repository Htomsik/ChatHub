package model

import "encoding/json"

const (
	SystemUser = "System"
)

// ChatMessageType Enum for type message
type ChatMessageType int

const (
	Message ChatMessageType = iota
	UsersList
)

func (msgType ChatMessageType) String() string {
	return [...]string{"Message", "UsersList"}[msgType]
}

func (msgType ChatMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgType.String())
}