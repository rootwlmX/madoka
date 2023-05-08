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

// RegisterRouter 路由注册
func (h *GithubHandler) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/v1/oauth/github")
	group.GET("/redirect", h.GetGithubToken)
}

// GetGithubToken Github重定向handler
func (h *GithubHandler) GetGithubToken(c *gin.Context) {
	code := c.Query("code")
	request := fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s", githubOauthBaseURL, clientID, clientSecret, code)
	response, err := http.Post(request, "application/json", nil)
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	defer response.Body.Close()
	bs, _ := io.ReadAll(response.Body)
	body := string(bs)
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
