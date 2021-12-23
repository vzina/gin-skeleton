package errno

import (
	"errors"
	"strconv"
)

var codes = map[int]string{}

type Message struct {
	Code      int    `json:"code"`
	Msg       string `json:"message"`
	Data      interface{} `json:"data"`
	RequestId string `json:"request_id"`
}

func NewError(code int, msg string) error {
	codes[code] = msg
	return errors.New(strconv.Itoa(code))
}

func GetMessageByErr(err error) Message {
	return GetMessage(err.Error())
}

func GetMessage(code string) Message {
	cv, e := strconv.Atoi(code)
	if e != nil {
		return Message{Msg: code, Data: []string{}}
	}

	if v, ok := codes[cv]; ok {
		return Message{Code: cv, Msg: v, Data: []string{}}
	}

	return Message{}
}

func (m *Message) WithData(data interface{}) *Message {
	m.Data = data
	return m
}

func (m *Message) WithRequestId(requestId string) *Message {
	m.RequestId = requestId
	return m
}