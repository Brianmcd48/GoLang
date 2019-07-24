package main

//echo prints command line arguments 

import(
    "fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	


    for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		
        sep = " "
	}
	fmt.Println(s)
	fmt.Println("%.2fs elasped\n", time.Since(start).Seconds())
}