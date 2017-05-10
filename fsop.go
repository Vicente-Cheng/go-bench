package main

import (
	"fmt"

	"github.com/Vicente-Cheng/go-bench/fs"
)

func main() {
	err, t_dur := fs.Dir_Create("./test_folder")
	if err != nil {
		fmt.Println("Create failure with error: ", err)
	}
	fmt.Printf("creating time: %v\n", t_dur)
	err, t_dur = fs.Dir_Remove("./test_folder")
	if err != nil {
		fmt.Println("Remove failure with error: ", err)
	}
	fmt.Printf("removing time: %v\n", t_dur)
}
