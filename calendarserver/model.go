package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Model struct {
}

const (
	holidaySuffix = "_holiday.cld"
	workdaySuffix = "_workday.cld"
	dir           = "./data/"
)

func (m *Model) GetHolidayList(key string) ([]string, error) {
	holidayFile, err := os.Open(getHolidayFileNmae(key))
	if err != nil {
		return nil, err
	}

	var ret []string
	reader := bufio.NewReader(holidayFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}

		ret = append(ret, strings.TrimSpace(line))
	}

	return ret, nil
}

func (m *Model) GetReplaceStringMap(key string) (map[string]string, error) {
	workdayFile, err := os.Open(getWorkdayFileNmae(key))
	if err != nil {
		return nil, err
	}

	ret := make(map[string]string)
	reader := bufio.NewReader(workdayFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}

		words := strings.Split(strings.TrimSpace(line), " ")
		if len(words) != 3 {
			return nil, fmt.Errorf("line \"%v\" have no enough field", line)
		}

		ret[words[0]] = words[1]
	}

	return ret, nil
}

func (m *Model) GetReplaceColorMap(key string) (map[string]int, error) {
	workdayFile, err := os.Open(getWorkdayFileNmae(key))
	if err != nil {
		return nil, err
	}

	ret := make(map[string]int)
	reader := bufio.NewReader(workdayFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}

		words := strings.Split(strings.TrimSpace(line), " ")
		if len(words) != 3 {
			return nil, fmt.Errorf("line \"%v\" have no enough field", line)
		}

		color, err := strconv.Atoi(words[2])
		if err != nil {
			return nil, fmt.Errorf("the third field of the line \"%v\" can not convert to int", line)
		}

		ret[words[0]] = color
	}

	return ret, nil
}

func getHolidayFileNmae(key string) string {
	return dir + key + holidaySuffix
}

func getWorkdayFileNmae(key string) string {
	return dir + key + workdaySuffix
}
