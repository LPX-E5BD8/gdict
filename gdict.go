package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/liipx/gdict/engien"
)

func main() {
	query := flag.String("q", "", "Queries.")
	style := flag.String("S", "dark", "Color Style. [dark,light]")
	say := flag.Bool("s", false, "Only used on MacOS.")
	e := flag.String("-engine", "youdao", "Query engine, only youdao yet.")

	flag.Parse()

	if *query == "" {
		fmt.Println("Useage: gdict -q 'lips'")
		os.Exit(1)
	}

	// TODO more engines supported
	switch *e {
	default:
		if len(os.Args) > 0 {
			engien.NewYoudao(*query, *style).Query()
		}
	}

	if *say && runtime.GOOS == "darwin" {
		sayPath, err := exec.LookPath("say")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		cmd := exec.Command(sayPath, *query)
		cmd.Run()
	}

}
