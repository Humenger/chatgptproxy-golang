# chatgptproxy-golang

golang version of chatgptproxy.me simulates request wrapping

# demo

```go
package main

import "github.com/Humenger/chatgptproxy-golang/ChatGptProxy"

func main() {
	chat := new(ChatGptProxy)
	completions, err := chat.Completions("你好")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(completions)
	}
	completions, err = chat.Completions("golang 随机生成一个[0-9a-z]的长度为16的字符串")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(completions)
	}
	completions, err = chat.Completions("再生成一个")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(completions)
	}

}
```
