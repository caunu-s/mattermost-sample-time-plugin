package main

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

type Response struct {
	Datetime string `json:"datetime"`
}

func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	if strings.Index(post.Message, "hoge") == -1 {
		return post, ""
	}

	p.API.LogDebug("Request WorldtimeAPI")

	url_time := "http://worldtimeapi.org/api/timezone/Asia/Tokyo"
	res_time, err_time := http.Get(url_time)
	if err_time != nil {
		fmt.Println(err_time)
    }
	defer res_time.Body.Close()
	rbody, _ := ioutil.ReadAll(res_time.Body)
	var response Response
	json.Unmarshal(rbody, &response)

	post.Message = fmt.Sprintf("%s\n%s", post.Message, response.Datetime)
	return post, ""
}