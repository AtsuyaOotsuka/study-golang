package main

import (
	"fmt"
	"flag"
)

func main() {

	// 引数を受け取る
	var (
		name = flag.String("name", "anonymous", "Your name")
		job  = flag.String("job", "no job", "Your job")
		skill = flag.String("skill", "no skill", "Your skill")
	)

	flag.Parse()

	fmt.Println("==============================")
	fmt.Printf("Name: %s\n", *name)
	fmt.Printf("Job: %s\n", *job)
	fmt.Printf("Skill: %s\n", *skill)
	fmt.Println("==============================")
}
