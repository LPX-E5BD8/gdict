package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/liipx/gdict/engien"
)

func main() {
	// 朗读开关
	say := flag.Bool("say", false, "使用系统发音软件朗读查询结果（MacOS only）")
	shortSay := flag.Bool("s", false, "")
	// 配色方案
	dark := flag.Bool("dark", true, "配色方案 dark")
	light := flag.Bool("light", false, "配色方案 light")
	// engine
	eng := flag.String("e", "youdao", "指定字典引擎")
	// --help and -h
	help := flag.Bool("help", false, "输出帮助信息")
	shortHelp := flag.Bool("h", false, "")
	flag.Parse()

	// 检查
	if flag.NArg() < 1 {
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

	// 目前认为多个连续的单词为单一句子，使用空格组合后查询
	// TODO: 增加参数，使程序把单词分开单独查询
	query := strings.Join(flag.Args(), " ")

	switch *eng {
	default:
		// default is 'youdao'
		// TODO: More engines
		engien.NewYoudao(query, style).Query()
	}

	if (*say || *shortSay) && runtime.GOOS == "darwin" {
		sayPath, err := exec.LookPath("say")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		cmd := exec.Command(sayPath, query)
		cmd.Run()
	}

}
