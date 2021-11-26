package main

import (
	"eicesoft/web-demo/cmd/gencont/pkg"
	"flag"
	"fmt"
	"os"
)

var (
	controller string
)

func init() {
	flagController := flag.String("c", "", "[Required] The name of the controller name\n")

	if !flag.Parsed() {
		flag.Parse()
	}

	if *flagController == "" {
		flag.Usage()
		os.Exit(1)
	}

	controller = *flagController
}

func main() {
	ui := `
██████╗  ███████╗███╗   ██╗ ██████╗ ██████╗ ███╗   ██╗████████╗
██╔════╝ ██╔════╝████╗  ██║██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝
██║  ███╗█████╗  ██╔██╗ ██║██║     ██║   ██║██╔██╗ ██║   ██║   
██║   ██║██╔══╝  ██║╚██╗██║██║     ██║   ██║██║╚██╗██║   ██║   
╚██████╔╝███████╗██║ ╚████║╚██████╗╚██████╔╝██║ ╚████║   ██║   
 ╚═════╝ ╚══════╝╚═╝  ╚═══╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝`
	fmt.Println(ui)

	pkg.GeneratorController(controller)
}
