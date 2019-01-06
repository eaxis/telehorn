TeleHorn
===

TeleHorn is simple and flexible tool to make newsletters in Telegram. 
You can use it for sending bulk messages to your users. 
It has API/GUI/CLI and Go-bindings. Check **[online-demo](http://199.247.25.75:2019/)** to try it yourself. 

- [Overview](#overview)
- [Requirements](#requirements)
- [Download](#download)
- [API/GUI](#apigui)
- [CLI](#cli)
- [In your code](#in-your-code)
- [Compiling](#compiling)
- [Limits](#limits)

## Overview
![TeleHorn](https://raw.githubusercontent.com/Narrator69/telehorn/master/promo.png)

## Requirements
The idea of TeleHorn implies you've:
1) Telegram Bot-API token. [Read more](https://core.telegram.org/bots/api).
2) The list of chats (Chat ID's) you want to send your message. [Read more](https://core.telegram.org/bots/api#chat).

## Download
You are welcome download TeleHorn from directory `redist` or [compile](#compiling) it yourself.

## API/GUI
Just one command to start API/GUI web-server:
```
$ telehorn web
```
It is! So simple. You should see:
```
Starting up HTTP-Server on localhost:2019...
```
You may also specify custom port, user and password to protect TeleHorn with basic auth:
```
$ telehorn web --port=1337 --user=admin --pass=supersonic
```
At this step GUI is ready to use. Check it. If you want to use API, you should send requests like this:
```
POST /submit HTTP/1.1
Host: localhost:1337 <-- Your port
Content-Length: 45
Content-Type: application/json; charset=UTF-8

{
  "token":"a",
  "chats":[2,3],
  "message":"b"
}
```
TeleHorn will answer you:
```
{
  "success": false,
  "description": "Incorrect token / Timeout."
}
```
Or:
```
{
  "success": true,
  "description": "We are sending your messages."
}
```


## CLI
Start with cli by typing:
```
$ telehorn cli --token=a --chat=1 --chat=2 --message=b
```
Or:
```
$ telehorn cli --token=a --chats=1,2 --message=b
```
Or:
```
$ telehorn cli --file=file.json

# Content of file.json:
# {
#   "token":"a",
#   "chats":[2,3],
#   "message":"b"
# }
```
Anyway, you will get something like this:
```
Incorrect token / Timeout.
```
Or this in case of success:
```
Done! Messages sent: 2, Failed: [], Successful: [1 2]
```
## In your code
At first you need to get TeleHorn as a dependency:
```
go get github.com/Narrator69/telehorn
```
Or:
```
dep ensure --add github.com/Narrator69/telehorn
```
Now you can use TeleHorn in your code:
```
package main

import (
	"fmt"
	"github.com/narrator69/telehorn"
)

func main() {
	app := cli.NewApp()
	
	service, _ := telehorn.NewTeleHorn("my-awesome-token")
	results := service.Send([]int{1, 2}, "Hello!")
	
	fmt.Println(results)
}
```

## Compiling
- You need Golang-language 1.6 or above
- You need [Dep](github.com/golang/dep)
- Place TeleHorn at `home/user/go/src/TeleHorn/`
- Execute `export GOPATH=/home/user/go/ && cd home/user/go/src/TeleHorn/ && dep ensure`
- And compile with `go build`

## Limits
TeleHorn sends about 30 messages per second and then sleep for 1 second to avoid limits of Telegram.
