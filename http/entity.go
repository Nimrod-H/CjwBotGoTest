package http

type Request struct {
	Content string `json:"content"`
	CE      CE     `json:"message_reference"`
	MsgID   string `json:"msg_id"`
}

type CE struct {
	MessageID             string `json:"message_id"`
	IgnoreGetMessageError bool   `json:"ignore_get_message_error"`
}
