package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Model struct {
}

const (
	holidaySuffix = "_holiday.cld"
	workdaySuffix = "_workday.cld"
	tipSuffix = "_tip.txt"
	dir           = "./data/"
)

func (m *Model) GetHolidayList(key string) ([]string, error) {
	holidayFile, err := os.Open(getHolidayFileNmae(key))
	if err != nil {
		return nil, err
	}

	defer holidayFile.Close()

	var ret []string
	reader := bufio.NewReader(holidayFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line);

		if len(line) > 0 {
			ret = append(ret, line)
		}
	}


	return ret, nil
}

func (m *Model) GetReplaceStringMap(key string) (map[string]string, error) {
	workdayFile, err := os.Open(getWorkdayFileNmae(key))
	if err != nil {
		return nil, err
	}

	defer workdayFile.Close()

	ret := make(map[string]string)
	reader := bufio.NewReader(workdayFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}
		
		line = strings.TrimSpace(line);

                if len(line) > 0 {
			words := strings.Split(line, " ")
			if len(words) != 3 {
				return nil, fmt.Errorf("line \"%v\" have no enough field", line)
			}

			ret[words[0]] = words[1]
		}
	}

	return ret, nil
}

func (m *Model) GetReplaceColorMap(key string) (map[string]int, error) {
	workdayFile, err := os.Open(getWorkdayFileNmae(key))
	if err != nil {
		return nil, err
	}

	defer workdayFile.Close()

	ret := make(map[string]int)
	reader := bufio.NewReader(workdayFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line);
                
                if len(line) > 0 {
			words := strings.Split(line, " ")
			if len(words) != 3 {
				return nil, fmt.Errorf("line \"%v\" have no enough field", line)
			}

			color, err := strconv.Atoi(words[2])
			if err != nil {
				return nil, fmt.Errorf("the third field of the line \"%v\" can not convert to int", line)
			}

			ret[words[0]] = color
		}
	}

	return ret, nil
}

func (m *Model) GetTip(key string) (string, error) {
	tipFile, err := os.Open(getTipFileNmae(key))
	if err != nil {
		return "", err
	}

	defer tipFile.Close()

	bytes, err := ioutil.ReadAll(tipFile)
	if err != nil {
		return "",err
	}

	return string(bytes),nil
}

func getHolidayFileNmae(key string) string {
	return dir + key + holidaySuffix
}

func getWorkdayFileNmae(key string) string {
	return dir + key + workdaySuffix
}

func getTipFileNmae(key string) string {
	return dir + key + tipSuffix
}
