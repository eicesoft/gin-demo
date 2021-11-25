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

	controller = *flagController
	if controller == "" {
		// panic(errors.New("-c 参数必须指定"))
		flag.PrintDefaults()
		os.Exit(1)
	}
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
