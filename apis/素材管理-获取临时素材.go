package apis

import (
	"encoding/json"
	"net/url"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetMedia 获取临时素材请求
// 文档：https://developer.work.weixin.qq.com/document/path/90254#获取临时素材
type ReqGetMedia struct {
	// MediaID 媒体文件id，见<a href="#10112" rel="nofollow">上传临时素材</a>，以及<a href="#42044" rel="nofollow">异步上传临时素材</a>，必填
	MediaID string `json:"media_id"`
}

var _ urlValuer = ReqGetMedia{}

func (x ReqGetMedia) intoURLValues() url.Values {
	var vals map[string]interface{}
	jsonBytes, _ := json.Marshal(x)
	_ = json.Unmarshal(jsonBytes, &vals)

	var ret url.Values = make(map[string][]string)
	for k, v := range vals {
		ret.Add(k, StrVal(v))
	}
	return ret
}

// RespGetMedia 获取临时素材响应
// 文档：https://developer.work.weixin.qq.com/document/path/90254#获取临时素材
type RespGetMedia struct{}

var _ bodyer = RespGetMedia{}

func (x RespGetMedia) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetMedia 获取临时素材
// 文档：https://developer.work.weixin.qq.com/document/path/90254#获取临时素材
func (c *ApiClient) ExecGetMedia(req ReqGetMedia) (RespGetMedia, error) {
	var resp RespGetMedia
	err := c.executeWXApiGet("/cgi-bin/media/get", req, &resp, true)
	if err != nil {
		return RespGetMedia{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetMedia{}, bizErr
	}
	return resp, nil
}
