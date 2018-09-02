package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/liipx/gdict/engine"
)

func main() {
	// 朗读开关
	say := flag.Bool("say", false, "使用系统发音软件朗读查询结果（MacOS only）")
	shortSay := flag.Bool("s", false, "")
	// 配色方案
	dark := flag.Bool("dark", true, "配色方案 dark")
	light := flag.Bool("light", false, "配色方案 light")
	// engine
	engStr := flag.String("e", "youdao", "指定字典引擎")
	// --help and -h
	help := flag.Bool("help", false, "输出帮助信息")
	shortHelp := flag.Bool("h", false, "")

	// alfred 输出格式
	workflow := flag.Bool("w", false, "是否使用Alfred输出模式")

	// auto parse
	// 将查询的词语参数分离，让参数的使用更灵活
	queryList := make([]string, 0)
	argList := []string{os.Args[0]}

	lastArg := ""
	for _, arg := range os.Args[1:] {

		if strings.Index(arg, "-") == 0 {
			argList = append(argList, arg)
			lastArg = arg
			continue
		}

		switch lastArg {
		case "-e":
			argList = append(argList, arg)
			lastArg = arg
			continue
		}

		queryList = append(queryList, arg)
		lastArg = arg
	}

	os.Args = argList
	flag.Parse()

	// 检查
	if len(queryList) < 1 {
		fmt.Println("未输入参数")
		flag.Usage()
		return
	}

	// 输出help
	if *help || *shortHelp {
		fmt.Printf("Options:\n%7s: -dark, -light"+
			"\n%7s: -s, --say  (MacOS only)"+
			"\n%7s: -youdao"+
			"\n%7s: -h, --help \n", "Style", "Read", "Engine", "Help")
		return
	}

	var style string
	if *light {
		style = "light"
	} else if *dark {
		style = "dark"
	}

	query := strings.Join(queryList, " ")

	// create engine
	var eng engine.Engine
	switch strings.TrimSpace(strings.ToLower(*engStr)) {
	case "iciba":
		// engine power by 'iciba'
		eng = engine.NewIciba(query, style)
	case "bing":
		// engine power by 'bing'
		eng = engine.NewBing(query, style)
	case "youdao":
		fallthrough
	default:
		// default is 'youdao'
		eng = engine.NewYoudao(query, style)
	}

	// 输出
	var result string
	if *workflow {
		result = eng.WFOutput()
	} else {
		result = eng.Query()
	}
	fmt.Println(result)

	// 发声
	if (*say || *shortSay) && runtime.GOOS == "darwin" {
		sayPath, err := exec.LookPath("say")
		if err != nil {
			log.Println(err)
		}

		cmd := exec.Command(sayPath, query)
		cmd.Run()
	}

}
