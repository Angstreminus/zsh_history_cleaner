package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Limit_persent int `yaml:"limit_persent"`
}

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
	}

	outStr := string(out)

	resString := strings.ReplaceAll(outStr, "\n", "")

	totalCnt, err := strconv.Atoi(resString)
	if err != nil {
		return 0, err
	}

	return totalCnt, nil

}

func getHistSize() (int, error) {
	histFile := os.Getenv("HOME") + "/.zsh_history"

	file, err := os.Open(histFile)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	histFileReader := bufio.NewReader(file)

	var cnt int

	for cnt = 0; ; cnt++ {
		_, err := histFileReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return cnt, nil
			} else {
				return 0, err
			}
		}
	}
}

func getStat(curComCnt int, totalComCnt int) (float64, error) {
	if totalComCnt == 0 {
		return 0, errors.New("Total commands count is 0")
	}

	var res float64 = (float64(curComCnt) / float64(totalComCnt)) * 100

	return res, nil
}

func ReadConfig() (string, error) {

	file, err := os.Open("config.yaml")
	if err != nil {
		return "", err
	}

	defer file.Close()

	cfgScanner := bufio.NewScanner(file)

	var confCont string

	for cfgScanner.Scan() {
		confCont = cfgScanner.Text()
	}

	if err := cfgScanner.Err(); err != nil {
		return "", err
	}

	return confCont, nil
}

func ParseConfig(cfgCont string) (*Config, error) {
	var cfg Config

	err := yaml.Unmarshal([]byte(cfgCont), &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {
	cfgContent, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	config, err := ParseConfig(cfgContent)
	if err != nil {
		log.Fatal(err)
	}

	//inf loop to run app
	for {
		totalComCnt, err := getTotalHistSize()
		if err != nil {
			fmt.Println(err)
			break
		}

		curHistSize, err := getHistSize()
		if err != nil {
			fmt.Println(err)
			break
		}

		stat, err := getStat(curHistSize, totalComCnt)
		if err != nil {
			fmt.Println(err)
			break
		}
		if stat > float64(config.Limit_persent) {
			clearHistFile()
		}
	}
}
