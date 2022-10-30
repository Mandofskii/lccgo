package src

import (
	"bufio"
	"log"
	"os"
	"unicode/utf8"
	"fmt"
	"strings"
)
func CharacterAndLineCounter(cmd string){
	counts := 0
	lines := 0
	isPrinted := false
	file, err := os.Open(cmd)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.IsDir() {
		files, err := file.Readdir(0)
		if err != nil {
			log.Fatal(err)
		}
		if !strings.HasSuffix(cmd, "/") {
			cmd += "/"
		}
		for _, v := range files {
			if v.IsDir() {
				fmt.Println("Directory: ", cmd + v.Name())
				CharacterAndLineCounter(cmd + v.Name())	
			} else {
				CharacterAndLineCounter(cmd + v.Name())
			}
			
		}		
	} else {
		fmt.Println("File: ", cmd)
		isPrinted = true
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			lines ++
			counts += utf8.RuneCountInString(text)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	if isPrinted {
		fmt.Println("Lines: ", lines)
		fmt.Println("Characters: ", counts)
	}
}
