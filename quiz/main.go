// 场景说明：给定一个csv文件problems.csv，文件中包含问题和答案，
// 使用程序从文件中读取问题，并将问题显示给终端用户，开始考试，
// 用户输入自己的答案后，按回车继续回答下一个问题，直到所有的问题回答完毕。
// 当用户考试完成后，显示用户的得分和用时。
// 问题回答设置超时时间，当用户回答问题的时间超时后，提示用户超时并结束考试。
// 当程序启动时，可以自定义超时时间。

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// 定义Quiz结构体，保存问题和得分
type Quiz struct {
	questions []*Question
	score     int
}

// 定义Question结构体，保存问题和问题的答案
type Question struct {
	problem string
	result  string
}

func quiz(quiz *Quiz, resChan chan int) {
	start := time.Now()
	var answer string
	counter := 1
	quiz.score = 0

	for _, q := range quiz.questions {
		fmt.Printf("问题%d: %v, 请输入答案: ", counter, q.problem)
		fmt.Scanln(&answer)
		if strings.TrimSpace(answer) == q.result {
			quiz.score++
		}
		counter++
	}

	dur := int(time.Since(start).Seconds())
	resChan <- dur
}

func readCSV(filename string) []*Question {
	// 读取并解析csv数据
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("csv file open error: %v\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	problems, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("csv file read error: %v\n", err)
	}

	var questions []*Question
	for _, problem := range problems {
		if len(problem) == 2 {
			q := &Question{problem: problem[0], result: problem[1]}
			questions = append(questions, q)
		}
	}

	return questions
}

var (
	filename string
	timeout  int
)

// 程序启动时，解析命令行参数
func init() {
	flag.StringVar(&filename, "f", "", "包含问题和答案的csv文件")
	flag.IntVar(&timeout, "t", 30, "考试超时时间")
	flag.Parse()

	if filename == "" {
		log.Fatal("请提供包含问题和答案的csv文件.")
	}
}

func main() {
	questions := readCSV(filename)
	qz := &Quiz{questions: questions}

	resCh := make(chan int)
	go quiz(qz, resCh)

	select {
	case dur := <-resCh:
		fmt.Printf("考试结束! 用时: %d秒, 得分: %d\n", dur, qz.score)
	case <-time.After(time.Duration(timeout) * time.Second):
		fmt.Println("考试已超时，结束考试! 得分:", qz.score)
	}
}
