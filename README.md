# Evaluation/英语试卷单词统计
- evaluation for CET4/6

- 整体性上有些许破碎

## 功能

### 用户相关

- 注册
- 登录
- 下载试卷
- 获取试卷列表
- 查看历年试卷单词出现次数（可选择删除各等级词汇）

### 管理员相关

- 上传试卷
- 删除试卷
- 提交wordcount作业
- 提交collect作业
- 查看作业列表
- 聚合作业数据

## 接口

具体参考每个服务的`api`文件（已粘贴至末尾）

### 用户相关

#### 注册

![register](.\images\register.png)

#### 登录

![login](.\images\login.png)

#### 刷新token

- 基于`jwt`的双token认证模式

![refresh](.\images\refresh.png)

### 文件相关

所有操作先指向`mysql`，文件操作在路径正确的情况下基本不会出错

#### 上传文件

- `mysql`保存信息
- 可上传所有格式，但仅在上传pdf时创建信息，需提前上传处理后的txt供mapreduce使用

#### 下载文件

- 查询`mysql`数据，验证是否存在(传参设计略有纰漏)

#### 查看文件列表

- 查询`mysql`数据

#### 删除文件

### 作业相关

#### 提交wordcount作业

- 由于内存和集群机器数量原因，处理结果只能划分成2个部分
- 对试卷对应的txt文件进行mapreduce

#### 提交collect作业

- 由于内存和集群机器数量原因，处理结果只能划分成2个部分
- 遍历存在的符合要求的所有文件进行整合

#### 查看作业列表

- 使用现成api查询

#### 聚合作业数据

- 在hdfs系统上对输出路径的所有文件进行拼接，聚合至一个对应路径，用type:1/2区分聚合哪个作业的数据

#### 查看历年试卷单词出现次数

## Hadoop

### hdfs分布式文件系统

- 高容错性：数据自动保存多个副本。它通过增加副本的形式，提高容错性；某一个副本丢失以后，它可以自动恢复。
- 用途：基于此实现了上传文件、下载文件、删除文件、数据聚合接口
- 给出文件系统的路径和结构：https://z55s3unkae.feishu.cn/mindnotes/RHiLbtgfxm1aIUnT9bAcEbZinhc?from=from_copylink

### Hadoop Streaming

- 允许使用各种语言自定义map和reduce
- 允许提交至hadoop进行分布式运算
- 用途：基于此实现了提交wordcount作业、提交collect作业接口

## Casbin权限管理（RBAC模型）

管理员拥有所有权限，用户没有管理员权限

方式：把用户id添加进管理员组，赋予权限

（身份比较少，不太好操作）

## Mapreduce实现

### 单个试卷的单词统计（wordcount）

#### mapper

- 数据清洗：正则表达式匹配标点、汉字、特殊符号和's'\单个字符，数字，再全部转换成小写
- 打印：按空格分隔，依次打印每个单词和1

```go
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
	reg := regexp.MustCompile(`[^\w\s]|[\p{Han}]|[\d]|[']s`)
	scanner := bufio.NewScanner(os.Stdin) // 创建一个从标准输入读取数据的扫描器
	for scanner.Scan() {
		// 对每一行应用正则表达式，替换掉匹配到的内容
		line := reg.ReplaceAllString(scanner.Text(), "")
		// 将所有单词转换为小写
		words := strings.ToLower(line)
		// 将转换后的单词按空格分割成数组
		wordArr := strings.Fields(words)
		for _, word := range wordArr { 
			fmt.Printf("%s\t%d\n", word, 1) // 输出单词和1，作为mapreduce的输入
		}
	}
}
```

#### reducer

- 排序：实现sort.Interface接口，根据出现次数排序
- 统计：将map阶段的标准输出进行累加

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type WordCount struct { 
	Word  string
	Count int
}

// 实现sort.Interface接口
type CountDesc []WordCount 

func (a CountDesc) Len() int           { return len(a) }                  
//返回切片长度
func (a CountDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }       
//交换元素位置
func (a CountDesc) Less(i, j int) bool { return a[i].Count > a[j].Count } 
//比较元素大小

func main() {
	wordCount := make(map[string]int) 
	scanner := bufio.NewScanner(os.Stdin)
	//循环读取标准输入的每一行数据
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t") //将当前行按照制表符分割成两部分
		if len(parts) != 2 {              
            //如果分割后的长度不是2，说明格式不正确，跳过此行
			continue
		}
		word := parts[0]                     
		count, err := strconv.Atoi(parts[1]) 
		if err != nil {                      //如果转化出错，说明格式不正确，跳过此行
			continue
		}
		if _, ok := wordCount[word]; !ok { //如果单词不存在于map中，初始化其出现次数为0
			wordCount[word] = 0
		}
		wordCount[word] += count 
	}
	var wc []WordCount            
	for w, c := range wordCount { 
		wc = append(wc, WordCount{w, c}) //将单词和单词出现的次数封装成WordCount结构体，添加到切片中
	}
	sort.Sort(CountDesc(wc)) //将切片按照出现次数降序排序
	for _, w := range wc {  
		fmt.Printf("%s\t%d\n", w.Word, w.Count)
	}
}
```

### 多个统计结果的过滤、排序(collect)

#### mapper

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) 
	for scanner.Scan() {
		//只需要全部打印即可
		fmt.Printf("%s\n", scanner.Text())
	}
}
```

#### reducer

- 排序
- 过滤：需要删除的单词文件每一行只出现单词，不出现数字，将符合这些特征的单词写入单独的数组，最后删除这些键

```go
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
	wordlist := make([]string, 0) //要删除的单词列表
	wordCount := make(map[string]int)
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
		word := parts[0]                     
		count, err := strconv.Atoi(parts[1])
		if err != nil {                      
            //如果转化出错，说明格式不正确，跳过此行
			continue
		}
		if _, ok := wordCount[word]; !ok { 
            //如果单词不存在于map中，初始化其出现次数为0
			wordCount[word] = 0
		}
		wordCount[word] += count //将单词出现的次数加入到map中
	}
	//删除
	for _, word := range wordlist {
		delete(wordCount, word)
	}
	var wc []WordCount            
	for w, c := range wordCount { 
		wc = append(wc, WordCount{w, c})
        //将单词和单词出现的次数封装成WordCount结构体，添加到切片中
	}
	sort.Sort(CountDesc(wc)) //将切片按照出现次数降序排序
	for _, w := range wc {   //遍历切片，输出每个单词和其出现次数
		fmt.Printf("%s\t%d\n", w.Word, w.Count)
	}
}
```

## 待解决的bug

- collect作业统计可以在本地测试时可以完成过滤，但在hadoop进行mapreduce时无法过滤
- 作业列表不能根据id查询，只能搜索全部

## Api文件一览

### `user.api`

```groovy
syntax = "v1"

//注册
type (
	RegisterReq {
		Username string `form:"username"`//用户名
		Password string `form:"password"`//密码
	}

	RegisterResp {
		Status int32  `json:"status"`
		Msg    string `json:"msg"`
	}
)

//登录
type (
	LoginReq {
		Username string `form:"username"`//用户名
		Password string `form:"password"`//密码
	}

	LoginInfo {
		ID           int64  `json:"id,omitempty"`//用户id
		Role         string `json:"role,omitempty"`//角色
		LoginTime    string `json:"login_time,omitempty"`//登录时间
		AccessToken  string `json:"access_token,omitempty"`//登录token
		RefreshToken string `json:"refresh_token,omitempty"`//刷新token
	}

	LoginResp {
		Status int32     `json:"status"`
		Msg    string    `json:"msg"`
		Data   LoginInfo `json:"data,omitempty"`
	}
)

@server(
	middleware :CORSMIDDLEWARE
)
service user {
	@doc "用户注册"
	@handler Register
	post /user/register (RegisterReq) returns (RegisterResp)
	@doc "用户登录"
	@handler Login
	post /user/login (LoginReq) returns (LoginResp)
}

//刷新token
type (
	RefreshReq {
		RefreshToken string `form:"refresh_token"`//刷新token
	}

	RefreshResp {
		Status       int32  `json:"status"`
		Msg          string `json:"msg"`
		AccessToken  string `json:"access_token,omitempty"`//登录token
		RefreshToken string `json:"refresh_token,omitempty"`//刷新token
	}
)

@server(
	middleware :CORSMIDDLEWARE,JWTMIDDLEWARE
)

service user{
	@doc "刷新token"
	@handler Refresh
	post /user/refresh (RefreshReq) returns (RefreshResp)
}
```

### `file.api`

```groovy
syntax = "v1"

//上传文件
type (
    UploadReq {
        Level int32 `form:"level"`//试卷等级
        Year int32 `form:"year"`  //年份
        Month int32 `form:"month"`//月份
        Set int32 `form:"set"`    //第几套
    }

    UploadResp {
        Status int32 `json:"status"`
        Msg string `json:"msg"`
    }
)

//下载文件
type (
    DownloadReq {
        Type int32 `form:"type"`  //可以当作没有，做保留
        Id int64 `form:"id"`      //文件id
        Level int32 `form:"level"`//试卷等级
        Year int32 `form:"year"`  //年份
        Month int32 `form:"month"`//月份
        Set int32 `form:"set"`    //第几套
    }

    DownloadResp {
        Status int32 `json:"status"`
        Msg string `json:"msg"`
    }
)

//获取文件列表
type (
    GetFileListReq {
        Level int32 `form:"level,optional"`//试卷等级
        Year int32 `form:"year,optional"`  //年份
        Month int32 `form:"month,optional"`//月份
        Set int32 `form:"set,optional"`    //第几套
    }

    GetFileListResp {
        Status int32 `json:"status"`
        Msg string `json:"msg"`
        List interface{} `json:"list,omitempty"`//文件列表
    }
)

//删除文件
type (
    RemoveFileReq {
        Id int64 `form:"id"`//文件id
    }

    RemoveFileResp {
        Status int32 `json:"status"`
        Msg string `json:"msg"`
    }
)

@server(
    middleware : CORSMIDDLEWARE,JWTMIDDLEWARE
)
service file {
    @doc "获取文件信息"
    @handler GetFileList
    post /file/list (GetFileListReq) returns (GetFileListResp)
    @doc "下载文件"
    @handler Download
    post /file/download (DownloadReq) returns (DownloadResp)
}

@server(
    middleware : CORSMIDDLEWARE,JWTMIDDLEWARE,AuthMIDDLEWARE
)

service file {
    @doc "上传文件"
    @handler Upload
    post /file/upload (UploadReq) returns (UploadResp)
    @doc "删除文件"
    @handler DeleteFile
    delete /file/delete (RemoveFileReq) returns (RemoveFileResp)
}
```

### `job.api`

```groovy
syntax = "v1"

//提交计数任务
type (
	CountJobReq {
		Level int32 `form:"level"`//4/6
		Year  int32 `form:"year"`//年份
		Month int32 `form:"month"`//月份
		Set   int32 `form:"set"`//第几套
	}

	CountJobResp {
		Status  int32  `json:"status"`
		Msg     string `json:"msg"`
		JobName string `json:"job_name"`//作业名称
	}
)

//聚合作业数据
type (
	JoinDataReq {
		Type   int32 `form:"type,optional"` //1-count数据聚合  2-collect数据聚合
		Year   int32 `form:"year,optional"`//1 年份
		Month  int32 `form:"month,optional"`//1 月份
		Set    int32 `form:"set,optional"`//1 第几套
		Level  int32 `form:"level,optional"`//1 2 等级
		Begin  int32 `form:"begin,optional"`//2 开始年份
		End    int32 `form:"end,optional"`//2 结束年份
		Except int32 `form:"except,optional"`//2 过滤的词汇类别，如：1对应3500
	}
	JoinDataResp {
		Status int32  `json:"status"`
		Msg    string `json:"msg"`
	}
)

//查看作业列表
type (
	ViewJobsReq {
		Id string `path:"id"`//0表示查看全部 否则为查询某个固定的
	}

	ViewJobsResp {
		Status  int32       `json:"status"`
		Msg     string      `json:"msg"`
		JobList interface{} `json:"jobList"`//作业列表
	}
)

//查看历年单词出现次数（可过滤）
type (
	ViewWordReq {
		Level  int32 `form:"level"`//4/6
		Begin  int32 `form:"begin"`//开始年份
		End    int32 `form:"end"`//结束年份
		Count  int32 `form:"count"`//显示前多少个
		Except int32 `form:"except"`//过滤的词汇类别，如：1对应3500
	}
	ViewWordResp {
		Status int32       `json:"status"`
		Msg    string      `json:"msg"`
		Words  interface{} `json:"words"`
	}
)

//提交统计作业
type (
	CollectJobReq {
		Begin  int32 `form:"begin"`//开始年份
		End    int32 `form:"end"`//结束年份
		Level  int32 `form:"level"`//4/6
		Except int32 `form:"except"`//过滤的词汇类别，如：1对应3500
	}
	CollectJobResp {
		Status  int32  `json:"status"`
		Msg     string `json:"msg"`
		JobName string `json:"job_name"`//作业名称
	}
)

@server(
	middleware : CORSMIDDLEWARE,JWTMIDDLEWARE
)
service job{
	@doc "统计历年单词"
	@handler CollectJob
	post /collect/job (CollectJobReq) returns (CollectJobResp)
}

@server(
	middleware : CORSMIDDLEWARE,JWTMIDDLEWARE,AuthMIDDLEWARE
)
service job {
	@doc "提交计数任务"
	@handler CountJob
	post /submit/job (CountJobReq) returns (CountJobResp)
	@doc "提交计数任务"
	@handler JoinData
	post /join/Data (JoinDataReq) returns (JoinDataResp)
	@doc "查看作业情况"
	@handler ViewJob
	get /view/job/:id (ViewJobsReq) returns (ViewJobsResp)
	@doc "查看高频单词"
	@handler ViewWord
	post /view/words (ViewWordReq) returns (ViewWordResp)
}
```

