package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
    "flag"
)

var sep = flag.String("s","•", "char indicating this line begins with a fortune")


func main() {    
    flag.Parse()
	s := bufio.NewScanner(os.Stdin)
	first := true
	fmt.Printf(`package main
import "fmt"
import "math/rand"
import "time"
var text []string = []string`)
	fmt.Printf("{")
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, *sep) {
			if !first {
				fmt.Printf(",\n")
			}
			fmt.Printf("`%s`", line)
			first = false
		}
	}
	fmt.Println("}")
	fmt.Println(`func main() {
r := rand.New(rand.NewSource(time.Now().UnixNano()))
fmt.Println(text[r.Intn(len(text))])`)
	fmt.Println("}")

}
