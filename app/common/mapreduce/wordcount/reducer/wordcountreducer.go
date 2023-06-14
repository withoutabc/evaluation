package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	wordCount := make(map[string]int)     // 创建一个map，用于存储单词和对应的出现次数
	scanner := bufio.NewScanner(os.Stdin) // 创建一个从标准输入读取数据的扫描器
	for scanner.Scan() {                  // 循环扫描输入的每一行
		line := scanner.Text()             // 获取扫描器当前扫描的行
		parts := strings.Split(line, "\t") // 将行按制表符分割成单词和出现次数
		if len(parts) != 2 {               // 如果分割出的部分不是两个，则跳过该行
			continue
		}
		word := parts[0]                     // 获取单词部分
		count, err := strconv.Atoi(parts[1]) // 将出现次数部分转换为整数
		if err != nil {                      // 如果转换出错，则跳过该行
			continue
		}
		if _, ok := wordCount[word]; !ok { // 如果单词在map中不存在，则将其初始值设置为0
			wordCount[word] = 0
		}
		wordCount[word] += count // 将出现次数加到单词的计数器上
	}
	for word, count := range wordCount { // 遍历map中的每个单词和对应的计数器
		fmt.Printf("%s\t%d\n", word, count) // 输出单词和对应的出现次数
	}
}
