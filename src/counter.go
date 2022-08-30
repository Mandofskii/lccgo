package src

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)
func CharacterAndLineCounter(cmd string) (int,int){
	var counts int
	lines := 0
	file, err := os.Open(cmd)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lines ++
		counts += utf8.RuneCountInString(text)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return counts,lines
}

func DirectoryCharacterAndLineCounter(cmd string) (int,int){
	var counts int
	lines := 0
	files, err := ioutil.ReadDir(cmd)
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
		if f.IsDir(){
			files2, err := ioutil.ReadDir(cmd+f.Name()+"/")
			if err != nil {
				log.Fatal(err)
			}
			for _, f2 := range files2 {
				file,err := os.Open(cmd+f.Name()+"/"+f2.Name())
				if err != nil{
					log.Fatal(err)
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan(){
					text := scanner.Text()
					lines ++
					counts += utf8.RuneCountInString(text)
				}
				if err := scanner.Err(); err != nil{
					log.Fatal(err)
				}
			}
			continue
		}
		file,err := os.Open(cmd+f.Name())
		if err != nil{
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			text := scanner.Text()
			lines ++
			counts += utf8.RuneCountInString(text)
		}
		if err := scanner.Err(); err != nil{
			log.Fatal(err)
		}
    }
	return counts,lines
}
