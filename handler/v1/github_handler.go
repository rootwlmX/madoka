package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"madoka/common"
	"madoka/util"
	"net/http"
)

// GithubHandler Github三方登录handler
type GithubHandler struct {
}

const clientID string = "d36422ab56e8ee83ebd0"
const clientSecret string = "915bbef9ce4a9ba04a4304fdcb5b707cd37bff41"
const githubOauthBaseURL string = "https://github.com/login/oauth/access_token"
const redirectURL string = "http://madokami.top/v1/oauth/github/redirect"

// RegisterRouter 路由注册
func (h *GithubHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/oauth/github")
	group.GET("/redirect", h.GetGithubToken)
}

// GetGithubToken Github重定向handler
func (h *GithubHandler) GetGithubToken(c *gin.Context) {
	code := c.Query("code")
	client := &http.Client{}
	loginURL := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s", githubOauthBaseURL, clientID, clientSecret, code, redirectURL)
	request, err := http.NewRequest("POST", loginURL, nil)
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	request.Header.Add("Accept", "application/json")
	response, _ := client.Do(request)
	defer response.Body.Close()
	bs, _ := io.ReadAll(response.Body)
	body := string(bs)
	fmt.Println(body)
	resultMap := util.JSONToMap(body)
	accessToken := resultMap["access_token"].(string)
	getGithubUserMessage(accessToken, c)
}

func getGithubUserMessage(accessToken string, c *gin.Context) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", "token "+accessToken)
	response, _ := client.Do(request)
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	defer response.Body.Close()
	bs, _ := io.ReadAll(response.Body)
	body := string(bs)
	common.Success(c, body, "ok")
}
