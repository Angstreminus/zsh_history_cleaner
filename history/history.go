package history

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"zshcleaner/config"
)

func ClearHistFile() error {
	histFilePath := os.Getenv("HOME") + "/.zsh_history"

	return os.Truncate(histFilePath, 0)
}

func GetTotalHistSize() (int, error) {
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

func GetHistSize() (int, error) {
	histFile := os.Getenv("HOME") + "/.zsh_history"

	file, err := os.Open(histFile)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	histFileReader := bufio.NewReader(file)

	for cnt := 0; ; cnt++ {
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

func GetStat(curComCnt int, totalComCnt int) (float64, error) {
	if totalComCnt == 0 {
		return 0, errors.New("Total commands count is 0")
	}

	return (float64(curComCnt) / float64(totalComCnt)) * 100, nil
}

func AnalyzeHistory(config *config.Config) error {
	totalComCnt, err := GetTotalHistSize()
	if err != nil {
		return err
	}

	curHistSize, err := GetHistSize()
	if err != nil {
		return err
	}

	stat, err := GetStat(curHistSize, totalComCnt)
	if err != nil {
		return err
	}

	if stat > float64(config.LimitPersent) {
		ClearHistFile()
	}
	return nil
}
