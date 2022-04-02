package videos

import (
	"DouyinParser/pkg/e"
	"DouyinParser/pkg/fetcher"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type RespJson struct {
	StatusCode int `json:"status_code"`
	ItemList   []struct {
		IsLiveReplay bool `json:"is_live_replay"`
		AnchorInfo   struct {
			Id   string `json:"id"`
			Name string `json:"name"`
			Type int    `json:"type"`
		} `json:"anchor_info"`
		ForwardId string `json:"forward_id"`
		Video     struct {
			Vid         string `json:"vid"`
			IsLongVideo int    `json:"is_long_video"`
			Cover       struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover"`
			Height       int `json:"height"`
			DynamicCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"dynamic_cover"`
			OriginCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"origin_cover"`
			Ratio    string      `json:"ratio"`
			BitRate  interface{} `json:"bit_rate"`
			PlayAddr struct {
				UrlList []string `json:"url_list"`
				Uri     string   `json:"uri"`
			} `json:"play_addr"`
			Width        int  `json:"width"`
			HasWatermark bool `json:"has_watermark"`
			Duration     int  `json:"duration"`
		} `json:"video"`
		ShareInfo struct {
			ShareDesc      string `json:"share_desc"`
			ShareTitle     string `json:"share_title"`
			ShareWeiboDesc string `json:"share_weibo_desc"`
		} `json:"share_info"`
		LongVideo interface{} `json:"long_video"`
		Desc      string      `json:"desc"`
		Images    interface{} `json:"images"`
		Category  int         `json:"category"`
		ChaList   []struct {
			Desc         string      `json:"desc"`
			UserCount    int         `json:"user_count"`
			Type         int         `json:"type"`
			IsCommerce   bool        `json:"is_commerce"`
			Cid          string      `json:"cid"`
			ChaName      string      `json:"cha_name"`
			ConnectMusic interface{} `json:"connect_music"`
			CoverItem    struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover_item"`
			ViewCount      int    `json:"view_count"`
			HashTagProfile string `json:"hash_tag_profile"`
		} `json:"cha_list"`
		Statistics struct {
			AwemeId      string `json:"aweme_id"`
			CommentCount int    `json:"comment_count"`
			DiggCount    int    `json:"digg_count"`
			PlayCount    int    `json:"play_count"`
			ShareCount   int    `json:"share_count"`
		} `json:"statistics"`
		TextExtra []struct {
			HashtagId   int64  `json:"hashtag_id"`
			Start       int    `json:"start"`
			End         int    `json:"end"`
			Type        int    `json:"type"`
			HashtagName string `json:"hashtag_name"`
		} `json:"text_extra"`
		AwemeId      string      `json:"aweme_id"`
		GroupId      int64       `json:"group_id"`
		AuthorUserId int64       `json:"author_user_id"`
		Promotions   interface{} `json:"promotions"`
		IsPreview    int         `json:"is_preview"`
		VideoLabels  interface{} `json:"video_labels"`
		LabelTopText interface{} `json:"label_top_text"`
		MixInfo      struct {
			NextInfo struct {
				MixName  string `json:"mix_name"`
				Desc     string `json:"desc"`
				CoverUrl struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
				} `json:"cover_url"`
			} `json:"next_info"`
			MixId  string `json:"mix_id"`
			Statis struct {
				PlayVv           int `json:"play_vv"`
				CollectVv        int `json:"collect_vv"`
				CurrentEpisode   int `json:"current_episode"`
				UpdatedToEpisode int `json:"updated_to_episode"`
			} `json:"statis"`
			Status struct {
				Status      int `json:"status"`
				IsCollected int `json:"is_collected"`
			} `json:"status"`
			Desc       string `json:"desc"`
			Extra      string `json:"extra"`
			CreateTime int    `json:"create_time"`
			MixName    string `json:"mix_name"`
			CoverUrl   struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover_url"`
		} `json:"mix_info"`
		Geofencing  interface{} `json:"geofencing"`
		FromXigua   bool        `json:"from_xigua"`
		CreateTime  int         `json:"create_time"`
		Duration    int         `json:"duration"`
		CommentList interface{} `json:"comment_list"`
		ImageInfos  interface{} `json:"image_infos"`
		RiskInfos   struct {
			Warn             bool   `json:"warn"`
			Type             int    `json:"type"`
			Content          string `json:"content"`
			ReflowUnplayable int    `json:"reflow_unplayable"`
		} `json:"risk_infos"`
		VideoText  interface{} `json:"video_text"`
		GroupIdStr string      `json:"group_id_str"`
		Author     struct {
			UniqueId     string      `json:"unique_id"`
			TypeLabel    interface{} `json:"type_label"`
			CardEntries  interface{} `json:"card_entries"`
			ShortId      string      `json:"short_id"`
			Nickname     string      `json:"nickname"`
			AvatarMedium struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"avatar_medium"`
			PolicyVersion interface{} `json:"policy_version"`
			Signature     string      `json:"signature"`
			AvatarLarger  struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"avatar_larger"`
			FollowersDetail interface{} `json:"followers_detail"`
			Uid             string      `json:"uid"`
			AvatarThumb     struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"avatar_thumb"`
			PlatformSyncInfo interface{} `json:"platform_sync_info"`
			Geofencing       interface{} `json:"geofencing"`
			MixInfo          interface{} `json:"mix_info"`
		} `json:"author"`
		ShareUrl  string `json:"share_url"`
		AwemeType int    `json:"aweme_type"`
	} `json:"item_list"`
	FilterList []interface{} `json:"filter_list"`
	Extra      struct {
		Logid string `json:"logid"`
		Now   int64  `json:"now"`
	} `json:"extra"`
}

type Parser struct {
}

// GetVideoParseURL 解析抖音视频无水印下载地址
func (p *Parser) GetVideoParseURL(url string) (string, error) {
	key, err := GetVideoKey(url)
	if e.HasError(err) {
		return "", err
	}
	vid, err := GetVideoVID(key)
	if e.HasError(err) {
		return "", err
	}
	fmt.Printf("vid:%s\n", vid)
	return fmt.Sprintf("https://aweme.snssdk.com/aweme/v1/play/?video_id=%s", vid), nil
}

// GetVideoVID 根据视频的 key 获取视频的 VID
func GetVideoVID(key string) (string, error) {
	api := "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=" + strings.TrimSpace(key)
	resp, err := fetcher.JsonRequest(http.MethodGet, api, "", &http.Header{})
	if e.HasError(err) {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	respJson := &RespJson{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &respJson)
	if e.HasError(err) {
		return "", err
	}
	return respJson.ItemList[0].Video.Vid, nil
}

// GetVideoKey 获取视频的key
func GetVideoKey(url string) (string, error) {
	resp, err := http.Get(url)
	if e.HasError(err) {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	keyRe := regexp.MustCompile(`/video/(\d+)\D?`)
	key := keyRe.FindStringSubmatch(resp.Request.URL.String())
	if len(key) < 2 {
		return "", fmt.Errorf("can not find video key")
	}
	return key[1], nil
}
