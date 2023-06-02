package chatgptproxy

import (
	"log"
	"testing"
)

func TestChatGptProxy_Completions(t *testing.T) {
	chat := new(ChatGptProxy)
	completions, err := chat.Completions("你好")
	if err != nil {
		t.Error(err)
	} else {
		log.Println(completions)
	}
	completions, err = chat.Completions("golang 随机生成一个[0-9a-z]的长度为16的字符串")
	if err != nil {
		t.Error(err)
	} else {
		log.Println(completions)
	}
	completions, err = chat.Completions("再生成一个")
	if err != nil {
		t.Error(err)
	} else {
		log.Println(completions)
	}

}
