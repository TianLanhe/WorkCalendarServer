package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var year = flag.Int("y", 2019, "year")
var month = flag.Int("m", 0, "month")
var key = flag.String("k", "jianglijie", "whose calendar you want to generate")
var dir = flag.String("d", "./data/", "output directory")

const (
	Red    = 0xFF0000
	Blue   = 0x0000FF
	Orange = 0xFFA500
)

func main() {
	// 解析命令行
	flag.Parse()

	// 月参数检查
	if *month < 1 || *month > 12 {
		fmt.Println("Illegal month(", *month, ")")
		return
	}

	// 创建数据目录
	if isExist, err := isDirExists(*dir); err != nil {
		fmt.Printf("Check directory exist \"%v\" error:%v\n", *dir, err)
		return
	} else if !isExist {
		err = os.MkdirAll(*dir, 0766)
		if err != nil {
			fmt.Printf("Check directory exist \"%v\" error:%v\n", *dir, err)
			return
		}
	}

	// 生成假期文件和工作文件
	holidayFileName := *dir + *key + "_holiday.cld"
	holidayFile, err := os.OpenFile(holidayFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Create file \"%v\" error:%v\n", holidayFileName, err)
		return
	}

	workdayFilename := *dir + *key + "_workday.cld"
	workdayFile, err := os.OpenFile(workdayFilename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Create file \"%v\" error:%v\n", workdayFilename, err)
		return
	}

	holidayWriter := bufio.NewWriter(holidayFile)
	workdayWriter := bufio.NewWriter(workdayFile)

	// 从标准输入读取上班与假期数据
	var str string
	if _, err = fmt.Scan(&str); err != nil {
		fmt.Printf("Read Calendar error:%v\n", err)
		return
	}

	for i, ch := range []rune(str) {
		if ch == '休' {
			dealHoliday(i, holidayWriter)
		} else {
			dealWorkday(i, ch, workdayWriter)
		}
	}

	holidayWriter.Flush()
	workdayWriter.Flush()
	holidayFile.Close()
	workdayFile.Close()
}

func isDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func dealHoliday(i int, writer *bufio.Writer) {
	record := fmt.Sprintf("%4d-%02d-%02d", *year, *month, i+1)

	fmt.Println("休：" + record)
	writer.WriteString(record + "\n")
}

func dealWorkday(i int, ch rune, writer *bufio.Writer) {
	if ch >= 'a' && ch <= 'z' {
		ch -= 32
	}

	color := 0
	if ch == 'A' {
		color = Red
	} else if ch == 'P' {
		color = Blue
	} else if ch == 'H' {
		color = Orange
	}

	record := fmt.Sprintf("%4d-%02d-%02d %c %d", *year, *month, i+1, ch, color)

	fmt.Println(record)
	writer.WriteString(record + "\n")
}
