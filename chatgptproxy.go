package chatgptproxy

import (
	"encoding/json"
	"fmt"
	"github.com/rapid7/go-get-proxied/proxy"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ChatGptProxy struct {
	SessionId string
	ChatId    string
	inited    bool
	client    *http.Client
	baseUrl   string
	proxy     string
}

func (that *ChatGptProxy) _init() {
	if that.inited {
		return
	}
	that.SessionId = that.genSessionId()
	that.ChatId = "0"
	that.baseUrl = "https://chatgptproxy.me"
	that.client = new(http.Client)
	proxyStr := strings.TrimSpace(that.proxy)
	url0, err := url.Parse(proxyStr)
	if proxyStr == "" || err != nil {
		proxy0 := proxy.NewProvider("").GetProxy("http", that.baseUrl)
		if proxy0 != nil {
			url0 = proxy0.URL()
		}
	}

	if url0 != nil && url0.String() != "" {
		log.Println("url0:", url0)
		that.client.Transport = &http.Transport{Proxy: http.ProxyURL(url0)}
	}
	that.inited = true
}
func (that *ChatGptProxy) SetProxy(proxy string) {
	that.proxy = proxy
}
func (that *ChatGptProxy) genSessionId() string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
func (that *ChatGptProxy) Completions(text string) (string, error) {
	if !that.inited {
		that._init()
	}
	err := that.heart()
	if err != nil {
		return "", err
	}
	response, err := that.conversation(that.ChatId, text)
	if err != nil {
		return "", err
	}
	if response.Code != 200 {
		return "", fmt.Errorf("conversation error:%s", response.CodeMsg)
	}
	that.ChatId = response.RespData.ChatId
	for {
		resultResponse, err := that.result()
		if err != nil {
			return "", err
		}
		if resultResponse.Code != 200 {
			return "", fmt.Errorf("result error:%s", resultResponse.CodeMsg)
		}
		if resultResponse.RespData.Status == 3 {
			return resultResponse.RespData.Answer, nil
		}
		time.Sleep(1 * time.Second)
	}
}

func (that *ChatGptProxy) heart() error {
	if !that.inited {
		that._init()
	}
	request := new(_ConversationRequest)
	request.Data = new(_ConversationRequestData)
	request.Data.SessionId = that.SessionId
	configByte, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", that.baseUrl+"/api/v1/chat/heart", strings.NewReader(string(configByte)))

	req.Header.Set("Content-Type", "application/json")

	resp, err := that.client.Do(req)
	if err != nil {
		return fmt.Errorf("ChatGptProxy response failed, Details of the error: \n%w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	data := new(_ConversationResponse)
	if err = json.Unmarshal(body, data); err != nil {
		return err
	}
	if data.Code != 200 {
		return fmt.Errorf("heart error: %s", data.CodeMsg)
	}
	return nil
}

func (that *ChatGptProxy) conversation(parentId, question string) (*_ConversationResponse, error) {
	if !that.inited {
		that._init()
	}
	//POST https://chatgptproxy.me/api/v1/chat/conversation
	//{"data":{"parent_id":"0","session_id":"2fa7mdomcy9ukutn","question":"你好","user_fake_id":"tnjjvkr1qqa1xbld"}}
	conversationRequest := &_ConversationRequest{}
	conversationRequest.Data = new(_ConversationRequestData)
	conversationRequest.Data.ParentId = parentId
	conversationRequest.Data.Question = question
	conversationRequest.Data.SessionId = that.SessionId
	configByte, err := json.Marshal(conversationRequest)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", that.baseUrl+"/api/v1/chat/conversation", strings.NewReader(string(configByte)))

	req.Header.Set("Content-Type", "application/json")

	resp, err := that.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ChatGptProxy response failed, Details of the error: \n%w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(_ConversationResponse)
	if err = json.Unmarshal(body, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (that *ChatGptProxy) result() (*_ResultResponse, error) {
	if !that.inited {
		that._init()
	}
	conversationRequest := &_ResultRequest{}
	conversationRequest.Data = new(_ResultRequestData)
	conversationRequest.Data.ChatId = that.ChatId
	conversationRequest.Data.SessionId = that.SessionId
	configByte, err := json.Marshal(conversationRequest)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", that.baseUrl+"/api/v1/chat/result", strings.NewReader(string(configByte)))

	req.Header.Set("Content-Type", "application/json")

	resp, err := that.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ChatGptProxy response failed, Details of the error: \n%w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(_ResultResponse)
	if err = json.Unmarshal(body, data); err != nil {
		return nil, err
	}
	return data, nil
}
