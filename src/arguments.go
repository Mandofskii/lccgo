package src
import (
	"fmt"
	"os"
	"regexp"
	"strings"
)
func ArgsHandler(cmd string){
	var directoryRegex = regexp.MustCompile("^(.*/)([^/]*)$")
	if len(os.Args) <= 1 {
		fmt.Println("Need at least one argument")
		fmt.Println("Use --help")
		os.Exit(0)
	} else {
		cmd = os.Args[1]
	}
	if strings.Contains(cmd, "--help") {
		fmt.Println("Usage:")
		fmt.Println("	lccgo ./file.txt")
		fmt.Println("	lccgo src/")
		os.Exit(0)
	}
	if directoryRegex.MatchString(cmd) {
		if strings.HasSuffix(cmd, "/") {
			var counts, lines = DirectoryCharacterAndLineCounter(cmd)
			fmt.Println("Directory: ",cmd)
			fmt.Println("Total Characters: ",counts,"\nTotal Lines: ",lines)
			os.Exit(0)
		}
		if _, err := os.Stat(cmd); err == nil {
			var counts, lines = CharacterAndLineCounter(cmd)
			fmt.Println("File: ",cmd)
			fmt.Println("Characters: ",counts,"\nLines: ",lines)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
}
