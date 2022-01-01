# Work WeChat
[Official api documentation](https://work.weixin.qq.com/api/doc/90000/90135/90664)


## Using the code samples
***

### Send text message
```
// Configure the client and push target information
var client Client = Client{
	Corpid:    "wwd327d0ea50c*****",
	Appsecret: "1kJw7t38aUxb3doG7IaLGXiX6DGobUllCLvoMf*****",
	Agentid:   "10*****",
	Touser:    "@all",
}

// Send text message
err = client.SendText(context)
if err != nil {
    return
}
```

### Send card message
```
// Configure the client and push target information

var client := Client{
    Corpid:    "wwd327d0ea50c*****",
    Appsecret: "1kJw7t38aUxb3doG7IaLGXiX6DGobUllCLvoMf*****",
    Agentid:   "10*****",
    Touser:    "@all",
}

// Send card message

card := Card{
    Title: "XXXXX",
    Description: "<div>XXXXX<div class=\\\"highlight\\\">                           ☀️ I hope everything goes well .</div>",
    Url: "https://XXXXX.com",
    //Btntxt: "more",
}

err = client.SendCard(card)
if err != nil {
    return
}
```