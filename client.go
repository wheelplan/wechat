package wechat

type Client struct {
	Corpid    string `yaml:"corpid"`
	Appsecret string `yaml:"appsecret"`
	Agentid   string `yaml:"agentid"`
	Touser    string `yaml:"touser"`
	Toparty   string `yaml:"toparty"`
	Totag     string `yaml:"totag"`
}
