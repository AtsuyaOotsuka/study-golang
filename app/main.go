package main

import (
	"fmt"
	"flag"
	"strings"
	"encoding/json"
	"io"
	"os"
)

type BusinessCard struct {
	Name   string   `json:"name"`
	Job    string   `json:"job"`
	Skills []string `json:"skills"`
	Wife   string   `json:"wife,omitempty"` // "none" のとき非表示にできるよ
}

func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	return strings.Repeat(" ", padding) + text
}

func getWife(name string) string {
	if strings.Contains(name, "onichan") {
		return "aochan"
	}
	return "none"
}

func printer(w io.Writer, isFileOutput bool, text string) {
	if isFileOutput {
		fmt.Fprintln(w, text)
	} else {
		fmt.Println(text)
	}

}

func printNormal(w io.Writer, isFileOutput bool, name string, job string, skill string, wife string, width int) {
	printer(w, isFileOutput, strings.Repeat("=", width))
	printer(w, isFileOutput, centerText(name, width))
	printer(w, isFileOutput, centerText(job, width))
	printer(w, isFileOutput, strings.Repeat("-", width))
	printer(w, isFileOutput, centerText("--- Skills ---", width))
	for _, s := range strings.Split(skill, ",") {
		printer(w, isFileOutput, centerText("• "+strings.TrimSpace(s), width))
	}
	if wife != "none" {
		printer(w, isFileOutput, strings.Repeat("-", width))
		printer(w, isFileOutput, centerText("MyWife: "+wife, width))
	}
	printer(w, isFileOutput, strings.Repeat("=", width))
}

func splitAndTrim(input string) []string {
	parts := strings.Split(input, ",")
	var trimmed []string
	for _, p := range parts {
		trimmed = append(trimmed, strings.TrimSpace(p))
	}
	return trimmed
}

func printJSON(w io.Writer, isFileOutput bool, name string, job string, skill string, wife string) {

	card := BusinessCard{
		Name:   name,
		Job:    job,
		Skills: splitAndTrim(skill),
		Wife:   wife,
	}

	jsonBytes, err := json.MarshalIndent(card, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return
	}

	if isFileOutput {
		fmt.Fprintln(w, string(jsonBytes))
	} else {
		fmt.Println(string(jsonBytes))
	}
	return // ← JSON出力したら処理終了
}

func main() {
	var name *string = flag.String("name", "", "Name")
	var job *string = flag.String("job", "", "Job")
	var skill *string = flag.String("skill", "", "Skill list (comma-separated)")
	var wife string
	var output *string = flag.String("output", "text", "Output format: text or json")
	var outputFile *string = flag.String("output-file", "", "Output file path (if specified, output will be written to this file)")
	flag.Parse()

	// 必須チェック
	if *name == "" || *job == "" || *skill == "" {
		fmt.Println("Error: -name, -job, and -skill are required.")
		return
	}

	wife = getWife(*name)

	width := 30 // 全体の横幅（好みに応じて調整可）

	var w io.Writer = os.Stdout

	var isFileOutput bool = false
	var isJson bool = false

	if *outputFile != "" {
		isFileOutput = true
		if strings.Contains(*outputFile, "json") {
			isJson = true
		}

		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		w = file
	}

	if isJson || *output == "json" {
		printJSON(w, isFileOutput, *name, *job, *skill, wife)
	} else {
		printNormal(w, isFileOutput, *name, *job, *skill, wife, width)
	}
}
