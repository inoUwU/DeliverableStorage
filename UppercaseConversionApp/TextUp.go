package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	println("This app converts string to uppercase")
	targetText := GetInputTxt()
	targetText = strings.ToUpper(targetText)
	println(targetText)
	WriteClipboard(targetText)
}

func GetInputTxt() string {
	//スキャナーの定義
	scanner := bufio.NewScanner(os.Stdin)
	//スキャン開始
	fmt.Print("Enter Text >")
	scanner.Scan()
	return scanner.Text()
}

func WriteClipboard(text string) {

	//スキャナーの定義
	scanner := bufio.NewScanner(os.Stdin)

	//local func
	Write := func(text string) {
		clipboard.WriteAll(text)
		println("copied!")
	}

	//無限ループ
	for {
		fmt.Print("Write Clipboard? y/n >")
		scanner.Scan()
		inputTxt := scanner.Text()

		switch inputTxt {
		case "y":
			Write(text)
			return
		case "n":
			return
		default:
		}
	}
}
