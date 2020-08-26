package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func SendDingTalk(UserName, DownloadUrl string) {
	access_token := "xxx"
	url := "https://oapi.dingtalk.com/robot/send?access_token="+access_token
	method := "POST"
	//	 变量注入到str
	StringPayload := fmt.Sprintf("{\"msgtype\": \"markdown\",\"markdown\": {\"title\":\"【下载通知】\",\"text\": \"【下载人】：%s \n【下载链接】：%s\n \"},\"at\": {\"isAtAll\": false}}", UserName, DownloadUrl)
	payload := strings.NewReader(StringPayload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func main() {
	UserName := "lijinghua"
	DownloadUrl := "https://lijinghua.club"
	SendDingTalk(UserName, DownloadUrl)
}
