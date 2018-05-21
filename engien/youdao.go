package engien

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/kr/pretty"
	"github.com/liipx/gdict/common"
)

const ydAPIOld = "http://fanyi.youdao.com/openapi.do"

var YoudaoKeys = []Youdao{
	{
		KeyFrom: "CoderVar",
		Key:     "802458398",
	},
	{
		KeyFrom: "whatMean",
		Key:     "1933652137",
	},
	{
		KeyFrom: "chinacache",
		Key:     "1247577973",
	},
	{
		KeyFrom: "huipblog",
		Key:     "439918742",
	},
	{
		KeyFrom: "chinacache",
		Key:     "1247577973",
	},
	{
		KeyFrom: "fanyi-node",
		Key:     "593554388",
	},
	{
		KeyFrom: "wbinglee",
		Key:     "1127870837",
	},
	{
		KeyFrom: "forum3",
		Key:     "1268771022",
	},
	{
		KeyFrom: "node-translator",
		Key:     "2058911035",
	},
	{
		KeyFrom: "kaiyao-robot",
		Key:     "2016811247",
	},
	{
		KeyFrom: "stone2083",
		Key:     "1576383390",
	},
	{
		KeyFrom: "myWebsite",
		Key:     "423366321",
	},
	{
		KeyFrom: "leecade",
		Key:     "54015339",
	},
	{
		KeyFrom: "github-wdict",
		Key:     "619541059",
	},
	{
		KeyFrom: "lanyuejin",
		Key:     "2033774719",
	},
}

type Youdao struct {
	KeyFrom string
	Key     string
	query   string
}

// 生成一个Youdao词典引擎
func NewYoudao(query string) *Youdao {
	i := common.RandInt(0, len(YoudaoKeys)-1)
	return &Youdao{
		KeyFrom: YoudaoKeys[i].KeyFrom,
		Key:     YoudaoKeys[i].Key,
		query:   query,
	}
}

// 获取老版本API的RUL
func (yd Youdao) getUrlOldVer() string {
	values := &url.Values{}
	values.Set("keyfrom", yd.KeyFrom)
	values.Set("key", yd.Key)
	values.Set("type", "data")
	values.Set("doctype", "json")
	values.Set("version", "1.1")
	values.Set("q", yd.query)

	return fmt.Sprintf("%s?%s", ydAPIOld, values.Encode())
}

// 获取新版本API的RUL
func (yd Youdao) getUrlNewVer() string {
	return fmt.Sprintf("%s?keyfrom=%s&key=%s&type=data&doctype=json&version=1.1&q=%s",
		ydAPIOld, yd.KeyFrom, yd.Key, yd.query)
}

// 查询
func (yd Youdao) Query() string {
	urlStr := yd.getUrlOldVer()
	resp, err := http.Get(urlStr)

	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	rs := new(YoudaoResult)
	json.Unmarshal(result, rs)
	return rs.Format()
}

// 结果集结构体
type YoudaoResult struct {
	Basic       basic    `json:"basic"`
	ErrorCode   int      `json:"errorCode"`
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
	Web         []web    `json:"web"`
}

type basic struct {
	Explains   []string `json:"explains"`
	Phonetic   string   `json:"phonetic"`
	UkPhonetic string   `json:"uk-phonetic"`
	UsPhonetic string   `json:"us-phonetic"`
}

type web struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

func (yr *YoudaoResult) Format() string {
	pretty.Println(yr)
	return ""
}
