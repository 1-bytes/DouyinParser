package controllers

import (
	"DouyinParser/app/models/videos"
	"DouyinParser/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type VideosController struct {
	BaseController
}

// ParserURL 解析抖音视频无水印下载地址
func (uc *VideosController) ParserURL(c *gin.Context) {
	appG := uc.GetAppG(c)
	url := c.Query("url")
	shareRe := regexp.MustCompile(`https?://v\.douyin\.com/\S+`)
	shareMatch := shareRe.FindStringSubmatch(url)
	if len(shareMatch) == 0 {
		appG.Response(http.StatusOK, e.InvalidParams, nil)
	}
	videosParser := videos.Parser{}
	videoParserURL, err := videosParser.GetVideoParseURL(shareMatch[0])
	if e.HasError(err) {
		appG.Response(http.StatusOK, e.ErrorDouyinParserUrlFail, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"video_url": videoParserURL,
	})
}
