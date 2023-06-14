package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // 创建一个从标准输入读取数据的扫描器
	for scanner.Scan() {                  // 循环扫描输入的每一行
		line := scanner.Text()        // 获取扫描器当前扫描的行
		words := strings.Fields(line) // 将行按空格分割成单词
		for _, word := range words {  // 遍历每个单词
			fmt.Printf("%s\t%d\n", word, 1) // 输出单词和1，作为mapreduce的输入
		}
	}
}
