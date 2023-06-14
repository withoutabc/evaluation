package file

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func DataClear(file *os.File) {
	// 打开文件
	defer file.Close()
	// 创建正则表达式，用于匹配标点、汉字、特殊符号和's'
	reg, err := regexp.Compile("[\\p{Han}]|[^a-zA-Z0-9\\s']+|'s")
	if err != nil {
		fmt.Println("正则表达式创建失败：", err)
		return
	}
	// 创建输出文件
	outFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("文件创建失败：", err)
		return
	}
	defer outFile.Close()
	// 使用bufio读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 对每一行应用正则表达式，替换掉匹配到的内容
		line := reg.ReplaceAllString(scanner.Text(), "")
		// 将所有单词转换为小写，并使用空格分隔
		words := strings.ToLower(line)
		// 将处理后的数据写入输出文件
		fmt.Fprintln(outFile, words)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("文件读取错误：", err)
		return
	}
	fmt.Println("处理完成！")
}
