package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Values struct {
	ContName string
}

// isDirExist 判断目录是否存在
func isDirExist(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}

	return s.IsDir()
}

// GeneratorController 生成控制器
func GeneratorController(controller string) {
	outDir := fmt.Sprintf("internal/controller/%s_controller", strings.ToLower(controller))

	if !isDirExist(outDir) {
		err := os.Mkdir(outDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	buf := new(bytes.Buffer)
	values := Values{controller}
	err := outputTemplate.Execute(buf, values)
	if err != nil {
		panic(err)
	}
	// formattedOutput, err := format.Source(buf.Bytes())
	// if err != nil {
	// 	panic(err)
	// }
	// buf = bytes.NewBuffer(formattedOutput)

	filename := fmt.Sprintf("%s/%s_controller.go", outDir, strings.ToLower(controller))
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0777); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("  └── file : ", fmt.Sprintf("%s", filename), " generator.")
}
