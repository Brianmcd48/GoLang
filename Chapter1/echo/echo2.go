package main 

import (
    "fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
    s, sep := "", ""
    for i, arg := range os.Args[0:] {
		s += sep + arg 
		sep = " "
		fmt.Println(s, i)
		s= ""
	}
	fmt.Println("%.2fs elasped\n", time.Since(start).Seconds())
}