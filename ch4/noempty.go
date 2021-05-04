package main

import "fmt"

func main ()  {
	data := []string{"one", "", "three"}
	fmt.Println(data)
	//fmt.Println(noempty(data))
	data = noempty(data)
	fmt.Println(data)
}
func noempty(strings []string) []string  {
	i := 0
	for _, s := range strings{
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
