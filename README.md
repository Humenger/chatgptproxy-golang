# chatgptproxy-golang

golang version of chatgptproxy.me simulates request wrapping

# demo

```go
package main

import "github.com/Humenger/chatgptproxy-golang/ChatGptProxy"

func main() {
	chat := new(ChatGptProxy)
	completions, err := chat.Completions("Hello")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(completions)
	}
	completions, err = chat.Completions("golang generates a random [0-9a-z] string of length 16")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(completions)
	}
	completions, err = chat.Completions("Generate another")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(completions)
	}

}
```
