package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// 创建正则表达式，用于匹配标点、汉字、特殊符号和's'\单个字符，数字
	reg, err := regexp.Compile("[\\p{Han}]|[^a-zA-Z0-9\\s']+|([a-zA-Z]'s)|([0-9]+)|([a-zA-Z]{1})")
	if err != nil {
		fmt.Println("正则表达式创建失败：", err)
		return
	}
	scanner := bufio.NewScanner(os.Stdin) // 创建一个从标准输入读取数据的扫描器
	for scanner.Scan() {
		// 对每一行应用正则表达式，替换掉匹配到的内容
		line := reg.ReplaceAllString(scanner.Text(), "")
		// 将所有单词转换为小写
		words := strings.ToLower(line)
		// 将转换后的单词按空格分割成数组
		wordArr := strings.Fields(words)
		for _, word := range wordArr { // 遍历每个单词
			fmt.Printf("%s\t%d\n", word, 1) // 输出单词和1，作为mapreduce的输入
		}
	}
}
