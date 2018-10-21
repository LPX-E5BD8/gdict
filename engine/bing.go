package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/liipx/gdict/common"
)

const bingUrl = "http://xtk.azurewebsites.net/BingDictService.aspx"

type Bing struct {
	query  string
	style  string
	result *bingResult
}

func NewBing(query string, style string) *Bing {
	b := &Bing{
		query: query,
		style: style,
	}
	return b.getResult()
}

type bingResult struct {
	Query         string `json:"word"`
	Pronunciation struct {
		AmE    string `json:"AmE"`
		AmEmp3 string `json:"AmEmp3"`
		BrE    string `json:"Bre"`
		BrEmp3 string `json:"BrEmp3"`
	}
	Defs []struct {
		Pos string `json:"pos"`
		Def string `json:"def"`
	} `json:"defs"`
	Sams []struct {
		Eng    string `json:"eng"`
		Chn    string `json:"chn"`
		Mp3Url string `json:"mp3Url"`
	}

	style string
}

func (b *Bing) getResult() *Bing {

	b.query = strings.Replace(b.query, " ", "%20", -1)
	translateUrl := fmt.Sprintf("%s?Word=%s", bingUrl, b.query)

	// get bing api
	resp, err := http.Get(translateUrl)
	if err != nil {
		errInfo := fmt.Sprintf("%s :%s", "网络连接异常，请检查网络状态",
			strings.Join(strings.Split(err.Error(), ":")[2:], ":"))
		fmt.Println(common.ColorIt(errInfo, common.Alert, b.style))
	}

	result, _ := ioutil.ReadAll(resp.Body)
	rs := new(bingResult)
	json.Unmarshal(result, rs)

	rs.style = b.style
	b.result = rs

	return b
}

// WorkflowOutput got the output fmt for alfred.
func (b *Bing) WFOutput() string {
	br := b.result
	result := &EngAlfResult{
		Items: make([]*WFItem, 0),
	}

	// generate workflow item
	for _, t := range br.Defs {
		result.Items = append(result.Items, &WFItem{
			Valid:    true,
			Title:    t.Def,
			Subtitle: t.Pos,
			Arg:      t.Def,
		})
	}

	resultByte, err := json.Marshal(result)
	if err != nil {
		os.Exit(1)
	}

	return string(resultByte)

	return ""
}

func (b *Bing) Query() string {
	return b.result.Format()
}

func (br *bingResult) Format() string {
	content := fmt.Sprintf("\n%s %s\n\n",
		common.ColorIt("查询:", common.Title, br.style),
		common.ColorIt(br.Query, common.Normal))

	// phonetic
	if pho := br.phoneticFormat(); strings.TrimSpace(pho) != "" {
		content += fmt.Sprintf("%s\n\n", pho)
	}

	// definition
	if exp := br.defFormat(); strings.TrimSpace(exp) != "" {
		content += fmt.Sprintf("%s\n\n%s\n", common.ColorIt("Exps:", common.Title, br.style), exp)
	}

	// samples
	if sam := br.samFormat(); strings.TrimSpace(sam) != "" {
		content += fmt.Sprintf("%s\n\n%s\n", common.ColorIt("例句:", common.Title, br.style), sam)
	}

	return content
}

// format phonetic
func (br *bingResult) phoneticFormat() string {
	context := ""
	if br.Pronunciation.BrE == "" {
		return context
	}

	context += common.ColorIt("英: ", common.Title, br.style) +
		common.ColorIt(br.Pronunciation.BrE, common.Alert, br.style) +
		strings.Repeat(" ", 4)
	context += common.ColorIt("美: ", common.Title, br.style) +
		common.ColorIt(br.Pronunciation.AmE, common.Alert, br.style)

	return context
}

// format translate
func (br *bingResult) defFormat() string {
	context := ""
	for _, def := range br.Defs {
		context += fmt.Sprintf("%5s  %s;\n", def.Pos, def.Def)
	}
	return context
}

// format samples
func (br *bingResult) samFormat() string {
	context := ""
	number := 0
	for _, sample := range br.Sams {
		number++
		context += fmt.Sprintf("%2d. %s\n%2s %s\n\n", number,
			common.ColorIt(sample.Eng, common.Alert, br.style),
			"",
			common.ColorIt(sample.Chn, common.Normal, br.style))
	}
	return context
}
