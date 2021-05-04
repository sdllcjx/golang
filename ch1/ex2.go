package main

import (
	"fmt"
	"os"
)

func main()  {
	/*for i:=0;i<len(os.Args);i++ {
		fmt.Println(i, os.Args[i])
	}*/
	for k, arg:=range os.Args[0:] {
		fmt.Println(k, arg)
	}
}
