package main

import (
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
	say := false
	style := "dark"
	e := "youdao"

	// 传入参数
	query := os.Args[1]

	// 动态参数设置
	startSetting := false
	for _, arg := range os.Args[1:] {
		switch strings.ToLower(arg) {
		case "-h", "--help":
			fmt.Printf("Options：\n\t配色类: -dark, -light" +
				"\n\t语音朗读(限MacOS): -s, --say" +
				"\n\t翻译引擎: -youdao" +
				"\n\t帮助: -h, --help")
			return
		case "-black", "-light":
			startSetting = true
			style = strings.Replace(arg, "-", "", -1)
		case "-s", "--say":
			startSetting = true
			say = true
		case "-youdao":
			startSetting = true
			style = strings.Replace(arg, "-", "", -1)
		default:
			if !startSetting {
				query += arg
			} else {
				fmt.Printf("\nUsage: \n\tgdict words ... [args ...]\n\t请将参数置于查询词之后")
				return
			}
		}
	}

	if query == "" {
		fmt.Println("Useage: gdict 'lips'")
		os.Exit(1)
	}

	switch e {
	default:
		// TODO: More engines
		engien.NewYoudao(query, style).Query()
	}

	if say && runtime.GOOS == "darwin" {
		sayPath, err := exec.LookPath("say")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		cmd := exec.Command(sayPath, query)
		cmd.Run()
	}

}
