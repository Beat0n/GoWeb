package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func CopeHandle(method, urlVal, data string) {
	client := &http.Client{}
	var req *http.Request

	if data == "" {
		urlArr := strings.Split(urlVal, "?")
		if len(urlArr) == 2 {
			urlVal = urlArr[0] + "?" + getParseParam(urlArr[1])
		}
		req, _ = http.NewRequest(method, urlVal, nil)
	} else {
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}

	cookie := &http.Cookie{
		Name:     "X-Xsrftoken",
		Value:    "dfkaldfadgdfg",
		HttpOnly: true,
		Expires:  time.Now().AddDate(1, 0, 0),
	}
	req.AddCookie(cookie)
	req.Header.Add("X-Xsrftoken", "jhigio456jk43b45")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func getParseParam(param string) string {
	return url.PathEscape(param)
}

func main() {
	CopeHandle("Get", "https://www.baidu.com", "")
}
