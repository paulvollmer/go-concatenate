package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/paulvollmer/go-concatenate"
)

var version = "0.1.0"

func usage() {
	fmt.Println("Usage: concat [-flags] || out [in]")
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println("\nIf you've found a bug please open an issue at github.com/paulvollmer/go-concatenate/issues")
}

func main() {
	flagVersion := flag.Bool("v", false, "print the version and exit")
	flagConfig := flag.String("c", "", "read a json config")
	flag.Usage = usage
	flag.Parse()
	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}
	if *flagConfig != "" {
		m := concatenate.NewManager()
		err := m.ReadConfig(*flagConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(127)
		}
		fmt.Println("config", m.Config)
		os.Exit(0)
	}

	totalArgs := len(os.Args)

	if totalArgs == 1 {
		usage()
		os.Exit(127)
	}
	if totalArgs == 2 {
		fmt.Println("Missing input sources")
		os.Exit(127)
	}
	if totalArgs > 2 {
		out := os.Args[1]
		sources := os.Args[2:]

		fmt.Println("out", out)
		fmt.Println("sources", sources)

	}
}
