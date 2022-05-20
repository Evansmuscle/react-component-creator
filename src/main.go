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
	
	noTs := false
	useJest := false

	for i := 0; i < len(os.Args[1:]); i++ {
		args[i] = os.Args[1:][i]	

		if args[i] == "--no-ts" {
			noTs = true
		}
		
		if args[i] == "--use-jest" {
			useJest = true
		}

		fmt.Println(args[i])
	}
	
	if componentName, ok := args[0]; ok {
		if noTs {
			tsxFile := "import './" + componentName + ".scss';\n\nimport React from 'react';\n\nconst " + componentName + " = () => {\n\treturn <div>Hello World</div>\n}\n\n export default " + componentName
			exportFile := "import " + componentName + " from './" + componentName + "'\n\nexport default " + componentName
			testFile := ""
			if useJest {
				testFile = `import renderer from 'react-test-renderer';
import ` + componentName + ` from './` + componentName +`';

it('Renders components', () => {
    const component = renderer.create(
   	<` + componentName + ` />
    )
    
    let tree = component.toJSON();
    
    // Your test code here..

    expect(tree).toMatchSnapshot();
});`
			} else {
				testFile = `import React from 'react';
import { mount } from '@cypress/react';
import ` + componentName + ` from './` + componentName +`';

it('Renders components', () => {
   mount(<` + componentName + ` />);

   // Your test code here..
});`
			}

			makeFolder(componentName)

			err := os.Chdir("./" + componentName)
			checkErr(err)
			
			makeFile(componentName + ".scss", "")
			makeFile("index.js", exportFile)
			makeFile(componentName + ".jsx", tsxFile)
			makeFile(componentName + ".spec.jsx", testFile)
			
			return
		}

		tsxFile := "import './" + componentName + ".scss';\n\nimport React from 'react';\n\ninterface " + componentName + "Props { }\n\nconst " + componentName + " = () => {\n\treturn <div>Hello World</div>\n}\n\n export default " + componentName
		exportFile := "import " + componentName + " from './" + componentName + "'\n\nexport default " + componentName
		testFile := ""
		if useJest {
			testFile = `import renderer from 'react-test-renderer';
import ` + componentName + ` from './` + componentName +`';

it('Renders components', () => {
    const component = renderer.create(
   	<` + componentName + ` />
    )
    
    let tree = component.toJSON();
    
    // Your test code here..

    expect(tree).toMatchSnapshot();
});`
		} else {
			testFile = `import React from 'react';
import { mount } from '@cypress/react';
import ` + componentName + ` from './` + componentName +`';

it('Renders components', () => {
   mount(<` + componentName + ` />);

   // Your test code here..
});`
		}

		makeFolder(componentName)

		err := os.Chdir("./" + componentName)
		checkErr(err)
		
		makeFile(componentName + ".scss", "")
		makeFile("index.ts", exportFile)
		makeFile(componentName + ".tsx", tsxFile)
		makeFile(componentName + ".spec.tsx", testFile)
	} else {
		fmt.Printf("You forgot to input a component name!")
	}
}