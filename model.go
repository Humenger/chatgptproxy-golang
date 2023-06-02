package chatgptproxy

type _ConversationRequest struct {
	Data *_ConversationRequestData `json:"data"`
}
type _ConversationRequestData struct {
	ParentId   string `json:"parent_id"`
	SessionId  string `json:"session_id"`
	Question   string `json:"question"`
	UserFakeId string `json:"user_fake_id"`
}
type _ConversationResponse struct {
	Code     int                            `json:"code"`
	CodeMsg  string                         `json:"code_msg"`
	TraceId  string                         `json:"trace_id"`
	RespData *_ConversationResponseRespData `json:"resp_data"`
}
type _ConversationResponseRespData struct {
	ChatId string `json:"chat_id"`
}

type _ResultRequest struct {
	Data *_ResultRequestData `json:"data"`
}
type _ResultRequestData struct {
	ChatId     string `json:"chat_id"`
	UserFakeId string `json:"user_fake_id"`
	SessionId  string `json:"session_id"`
}
type _ResultResponse struct {
	Code     int                      `json:"code"`
	CodeMsg  string                   `json:"code_msg"`
	TraceId  string                   `json:"trace_id"`
	RespData *_ResultResponseRespData `json:"resp_data"`
}
type _ResultResponseRespData struct {
	Answer       string `json:"answer"`
	Status       int    `json:"status"`
	ShareCode    string `json:"share_code"`
	ShareCodeAll string `json:"share_code_all"`
	ShareHost    string `json:"share_host"`
}
