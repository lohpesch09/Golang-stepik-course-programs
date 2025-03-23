package main

import (
	"fmt"
	"os"
)

func main_1() {
	os.Chdir("testdata")
	// dirs, _ := os.ReadDir("testdata")
	stat, _ := os.Stat("zzfile.txt")
	fmt.Println(stat.Size())
	// for _, dir := range dirs {
	// 	stat, _ := os.Stat(dir.Name())
	// 	fmt.Println(dir.Name())

	// }
	// a := make([]int, 0)
	// a = append(a, 1, 2, 3, 4)
}
