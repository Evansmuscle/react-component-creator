package main

import (
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func makeFolder(folderName string) {
	err := os.Mkdir(folderName, 0755)
	checkErr(err)
}

func makeFile(fileName string, content string) {
	file, err := os.Create(fileName)
	checkErr(err)
	
	_, err = file.WriteString(content)
	checkErr(err)
}

func main() {
	args := make(map[int]string, len(os.Args[1:]))

	for i := 0; i < len(os.Args[1:]); i++ {
		args[i] = os.Args[1:][i]	
	}
	
	if componentName, ok := args[0]; ok {
		tsxFile := "import './" + componentName + "';\n\nimport React from 'react';\n\ninterface " + componentName + "Props { }\n\n const " + componentName + " = () => {\n\treturn <div>Hello World</div>\n}\n\n export default " + componentName
		exportFile := "import " + componentName + " from './" + componentName + "'\n\n export default " + componentName

		makeFolder(componentName)

		err := os.Chdir("./" + componentName)
		checkErr(err)
		
		makeFile(componentName + ".tsx", tsxFile)
		makeFile(componentName + ".scss", "")
		makeFile("index.ts", exportFile)
	} else {
		fmt.Printf("You forgot to input a component name!")
	}
}