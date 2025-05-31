package main

import (
	"fmt"
	"flag"
	"strings"
)

func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	return strings.Repeat(" ", padding) + text
}

func main() {
	var name *string = flag.String("name", "", "Name")
	var job *string = flag.String("job", "", "Job")
	var skill *string = flag.String("skill", "", "Skill list (comma-separated)")
	var wife string
	flag.Parse()

	// 必須チェック
	if *name == "" || *job == "" || *skill == "" {
		fmt.Println("Error: -name, -job, and -skill are required.")
		return
	}

	{
		if strings.Contains(*name, "onichan") {
			wife = "aochan"
		}
	}

	width := 30 // 全体の横幅（好みに応じて調整可）

	fmt.Println(strings.Repeat("=", width))
	fmt.Println(centerText(*name, width))
	fmt.Println(centerText(*job, width))
	fmt.Println(strings.Repeat("-", width))
	fmt.Println(centerText("Skills: "+*skill, width))
	fmt.Println(centerText("MyWife: "+wife, width))
	fmt.Println(strings.Repeat("=", width))
}
