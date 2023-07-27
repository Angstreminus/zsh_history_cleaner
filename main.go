package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func clearHistFile() {
	histFilePath := os.Getenv("HOME") + "/.zsh_history"

	if err := os.Truncate(histFilePath, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
}

func getTotalHistSize() (int, error) {
	out, err := exec.Command("echo", os.Getenv("HISTSIZE")).Output()
	if err != nil {
		return 0, err
	} else {
		outStr := string(out)
		resString := strings.ReplaceAll(outStr, "\n", "")

		totalCnt, err := strconv.Atoi(resString)
		if err != nil {
			return 0, err
		}
		return totalCnt, nil
	}
}

func getCurrentSize() (int, error) {
	histFile := os.Getenv("HOME") + "/.zsh_history"

	//get current content of zsh history
	histContent, err := ioutil.ReadFile(histFile)
	if err != nil {
		return 0, err
	}

	var histValStr string

	for _, elem := range histContent {
		histValStr += string(elem)
	}
	res := strings.Split(histValStr, "\n")

	return len(res), nil
}

func getStat(curComCnt int, totalComCnt int) error {
	if totalComCnt == 0 {
		return errors.New("Total commands count is 0")
	} else {
		fmt.Println("Max limit zsh history file size is:", totalComCnt)

		fmt.Println("Current hist file size is:", curComCnt)

		var res float64 = (float64(curComCnt) / float64(totalComCnt)) * 100
		fmt.Printf("Fill of history file is %.2f", res)

		return nil
	}
}

func main() {
	clearHistFile()

	curComCnt, err := getCurrentSize()
	if err != nil {
		fmt.Println(err)
	}

	totalComCnt, err := getTotalHistSize()
	if err != nil {
		fmt.Println(err)
	}

	getStat(curComCnt, totalComCnt)

}
