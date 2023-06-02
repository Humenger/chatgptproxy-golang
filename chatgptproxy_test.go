package chatgptproxy

import (
	"log"
	"testing"
)

func TestChatGptProxy_Completions(t *testing.T) {
	chat := new(ChatGptProxy)
	completions, err := chat.Completions("Hello")
	if err != nil {
		t.Error(err)
	} else {
		log.Println(completions)
	}
	completions, err = chat.Completions("golang generates a random [0-9a-z] string of length 16")
	if err != nil {
		t.Error(err)
	} else {
		log.Println(completions)
	}
	completions, err = chat.Completions("Generate another")
	if err != nil {
		t.Error(err)
	} else {
		log.Println(completions)
	}

}
