package message

import (
	"encoding/json"
	"log"
)

type SendMsg struct {
	Name       string
	TargetName string
	Body       []byte
}

type ReceiveMsg struct {
	Name       string
	SourceName string
	Body       []byte
}

func (s *SendMsg) Decode(buf []byte) {
	json.Unmarshal(buf, s)
}

func (s *SendMsg) Encode() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (r *ReceiveMsg) Decode(buf []byte) {
	json.Unmarshal(buf, r)
}

func (r *ReceiveMsg) Encode() []byte {
	data, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
