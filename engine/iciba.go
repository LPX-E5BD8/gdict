package engine

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/liipx/gdict/common"
)

const icibaKey = "C8DDDA9B1360645BA7C3888DF6F54702"
const icibaUrl = "http://dict-co.iciba.com/api/dictionary.php"

type Iciba struct {
	query  string
	style  string
	result *IcibaResult
}

func NewIciba(query, style string) *Iciba {
	if query == "" {
		return nil
	}

	i := &Iciba{
		query: query,
		style: style,
	}
	i.result = i.getResult()
	return i
}

type IcibaResult struct {
	Key         string   `xml:"key"`
	Ps          []string `xml:"ps"`
	Pos         []string `xml:"pos"`
	Acceptation []string `xml:"acceptation"`
	Sent        []struct {
		Orig  string `xml:"orig"`
		Trans string `xml:"trans"`
	} `xml:"sent"`

	style string
}

// get the result from iciba api
func (i *Iciba) getResult() *IcibaResult {
	translateUrl := fmt.Sprintf("%s?key=%s&w=%s", icibaUrl, icibaKey, i.query)
	// get iciba api
	resp, err := http.Get(translateUrl)
	if err != nil {
		errInfo := fmt.Sprintf("%s :%s", "网络连接异常，请检查网络状态",
			strings.Join(strings.Split(err.Error(), ":")[2:], ":"))
		fmt.Println(common.ColorIt(errInfo, common.Alert, i.style))
	}

	result, _ := ioutil.ReadAll(resp.Body)
	rs := new(IcibaResult)
	rs.style = i.style

	xml.Unmarshal(result, rs)

	return rs
}

// WorkflowOutput got the output fmt for alfred.
func (i *Iciba) WFOutput() string {
	ir := i.result
	result := &EngAlfResult{
		Items: make([]*WFItem, 0),
	}

	// generate workflow item
	for _, t := range ir.Sent {
		result.Items = append(result.Items, &WFItem{
			Valid:    true,
			Title:    strings.TrimSpace(t.Trans),
			Subtitle: strings.TrimSpace(t.Orig),
			Arg:      strings.TrimSpace(t.Trans),
		})
	}

	resultByte, err := json.Marshal(result)
	if err != nil {
		os.Exit(1)
	}

	return string(resultByte)

	return ""
}

func (i *Iciba) Query() string {
	return i.result.Format()
}

// format the result
func (ir *IcibaResult) Format() string {
	content := fmt.Sprintf("\n%s %s\n\n",
		common.ColorIt("查询:", common.Title, ir.style),
		common.ColorIt(ir.Key, common.Normal))

	// phonetic
	if pho := ir.phoneticFormat(); strings.TrimSpace(pho) != "" {
		content += fmt.Sprintf("%s\n\n", pho)
	}

	// definition
	if exp := ir.acceptationFormat(); strings.TrimSpace(exp) != "" {
		content += fmt.Sprintf("%s\n\n%s\n", common.ColorIt("Exps:", common.Title, ir.style), exp)
	}

	// samples
	if sam := ir.sentFormat(); strings.TrimSpace(sam) != "" {
		content += fmt.Sprintf("%s\n\n%s\n", common.ColorIt("例句:", common.Title, ir.style), sam)
	}

	return content
}

// format phoneticFormat
func (ir *IcibaResult) phoneticFormat() string {
	context := ""

	if len(ir.Ps) == 0 || ir.Ps[0] == "" {
		return context
	}

	context += common.ColorIt("英: ", common.Title, ir.style) +
		common.ColorIt(ir.Ps[0], common.Alert, ir.style) +
		strings.Repeat(" ", 4)
	context += common.ColorIt("美: ", common.Title, ir.style) +
		common.ColorIt(ir.Ps[1], common.Alert, ir.style)

	return context
}

// format exp
func (ir *IcibaResult) acceptationFormat() string {
	context := ""

	for i, pos := range ir.Pos {
		if pos == "" {
			continue
		}

		context += fmt.Sprintf("%5s  %s\n", pos, strings.TrimSpace(ir.Acceptation[i]))
	}

	return context
}

// format sent
func (ir *IcibaResult) sentFormat() string {
	context := ""
	number := 0

	if len(ir.Sent) == 0 {
		return context
	}

	for _, sample := range ir.Sent {
		number++
		context += fmt.Sprintf("%2d. %s\n%2s  %s\n\n", number,
			common.ColorIt(strings.TrimSpace(sample.Orig), common.Alert, ir.style),
			"",
			common.ColorIt(strings.TrimSpace(sample.Trans), common.Normal, ir.style))
	}

	return context
}
