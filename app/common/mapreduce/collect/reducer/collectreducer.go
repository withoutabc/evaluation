package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type WordCount struct { //定义一个结构体，包含两个字段：单词和单词出现的次数
	Word  string
	Count int
}

// 实现sort.Interface接口

type CountDesc []WordCount //定义一个结构体切片，用于排序

func (a CountDesc) Len() int           { return len(a) }                  //返回切片长度
func (a CountDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }        //交换元素位置
func (a CountDesc) Less(i, j int) bool { return a[i].Count > a[j].Count } //比较元素大小

func main() {
	wordlist := make([]string, 0)     //要删除的单词列表
	wordCount := make(map[string]int) //定义一个map，用于记录单词出现的次数
	scanner := bufio.NewScanner(os.Stdin)
	//循环读取标准输入的每一行数据
	for scanner.Scan() {
		line := scanner.Text()
		//如果没有数字则表示这个单词要删除，直接continue
		parts := strings.Split(line, "\t") //将当前行按照制表符分割成两部分
		if len(parts) != 2 {
			if len(parts) == 1 {
				//如果是1，那就说明要删除，把这个单词写入
				wordlist = append(wordlist, parts[0])
			}
			continue
		}
		word := parts[0]                     //获取单词
		count, err := strconv.Atoi(parts[1]) //将单词出现的次数转化为整型
		if err != nil {                      //如果转化出错，说明格式不正确，跳过此行
			continue
		}
		if _, ok := wordCount[word]; !ok { //如果单词不存在于map中，初始化其出现次数为0
			wordCount[word] = 0
		}
		wordCount[word] += count //将单词出现的次数加入到map中
	}
	//删除
	for _, word := range wordlist {
		delete(wordCount, word)
	}
	var wc []WordCount            //定义一个WordCount结构体切片，用于存放map中的数据
	for w, c := range wordCount { //遍历map
		wc = append(wc, WordCount{w, c}) //将单词和单词出现的次数封装成WordCount结构体，添加到切片中
	}
	sort.Sort(CountDesc(wc)) //将切片按照出现次数降序排序
	for _, w := range wc {   //遍历切片，输出每个单词和其出现次数
		fmt.Printf("%s\t%d\n", w.Word, w.Count)
	}
}
