package engien

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

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
	style   string
}

// 生成一个Youdao词典引擎
func NewYoudao(query string, style string) *Youdao {
	i := common.RandInt(0, len(YoudaoKeys)-1)
	return &Youdao{
		KeyFrom: YoudaoKeys[i].KeyFrom,
		Key:     YoudaoKeys[i].Key,
		query:   query,
		style:   style,
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
	rs.style = yd.style
	return rs.Format()
}

// 结果集结构体
type YoudaoResult struct {
	Basic       basic    `json:"basic"`
	ErrorCode   int      `json:"errorCode"`
	Query       string   `json:"query"`
	Translation []string `json:"translation"`
	Web         []web    `json:"web"`
	style       string
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

// main format
func (yr *YoudaoResult) Format() string {
	content := fmt.Sprintf("\n%s %s\n\n",
		common.ColorIt("查询:", common.Title, yr.style),
		common.ColorIt(yr.Query, common.Normal))

	// phonetic
	if pho := yr.phoneticFormat(); strings.TrimSpace(pho) != "" {
		content += fmt.Sprintf("%s\n\n", pho)
	}

	// explains
	if exp := yr.explainsFormat(); strings.TrimSpace(exp) != "" {
		content += fmt.Sprintf("%s\n\n%s\n", common.ColorIt("Exps:", common.Title, yr.style), exp)
	}

	// translation
	if tran := yr.transFormat(); strings.TrimSpace(tran) != "" {
		content += fmt.Sprintf("%s\n\n%s\n\n", common.ColorIt("翻译:", common.Title, yr.style), tran)
	}

	// web
	if web := yr.webFormat(); strings.TrimSpace(web) != "" {
		content += fmt.Sprintf("%s\n\n%s\n\n", common.ColorIt("网络释义:", common.Title, yr.style), web)
	}

	return content
}

// format explain
func (yr *YoudaoResult) explainsFormat() string {
	context := ""
	number := 1
	for _, exp := range yr.Basic.Explains {
		exp := strings.Split(exp, ". ")
		if yr.Basic.UkPhonetic == "" && yr.Basic.UsPhonetic == "" {
			numStr := fmt.Sprintf("% 2d", number)
			context += common.ColorIt(numStr, common.Alert, yr.style) + "." + strings.Join(exp, ".\t") + "\n\n"
		} else {
			for i, v := range exp {
				if (i+1)%2 == 0 {
					context += v + "\n"
				} else {
					context += fmt.Sprintf("%5s.  ", v)
				}
			}
		}
		number++
	}
	return context
}

// format phonetic
func (yr *YoudaoResult) phoneticFormat() string {
	context := ""
	if yr.Basic.UkPhonetic == "" && yr.Basic.UsPhonetic == "" && yr.Basic.Phonetic != "" {
		context += common.ColorIt("拼音: ", common.Title, yr.style) + yr.Basic.Phonetic
	} else if yr.Basic.UkPhonetic != "" || yr.Basic.UsPhonetic != "" {
		context += common.ColorIt("英: ", common.Title, yr.style) +
			common.ColorIt(yr.Basic.UkPhonetic, common.Alert, yr.style) +
			strings.Repeat(" ", 4)
		context += common.ColorIt("美: ", common.Title, yr.style) +
			common.ColorIt(yr.Basic.UsPhonetic, common.Alert, yr.style)
	}
	return context
}

// format translation
func (yr *YoudaoResult) transFormat() string {
	content := ""
	content += fmt.Sprintf("  %s", strings.Join(yr.Translation, "\n  "))

	return content
}

// format web
func (yr *YoudaoResult) webFormat() string {
	context := ""
	number := 1
	for _, v := range yr.Web {
		numStr := fmt.Sprintf("% 2d", number)
		context += fmt.Sprintf("  %s.%s:\n    %s\n\n",
			common.ColorIt(numStr, common.Alert, yr.style),
			common.ColorIt(v.Key, common.Alert, yr.style),
			strings.Join(v.Value, ", "))
		number++
	}

	return context
}
