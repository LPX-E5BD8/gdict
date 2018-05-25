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

	// 检查
	if len(os.Args) < 1 {
		fmt.Println("未输入参数")
		return
	}

	// 传入参数
	query := os.Args[1]

	// 动态参数设置
	startSetting := false
	if len(os.Args) > 2 {
		for _, arg := range os.Args[2:] {
			switch strings.ToLower(arg) {
			case "-h", "--help":
				fmt.Printf("Options:\n%7s: -dark, -light"+
					"\n%7s: -s, --say  (MacOS only)"+
					"\n%7s: -youdao"+
					"\n%7s: -h, --help \n", "Style", "Read", "Engine", "Help")
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
