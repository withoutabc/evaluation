package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // 创建一个从标准输入读取数据的扫描器
	for scanner.Scan() {
		//只需要全部打印即可
		fmt.Printf("%s\n", scanner.Text())
	}
}
