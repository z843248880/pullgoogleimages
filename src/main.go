package main

import (
	"fmt"
	"strings"
	// "fmt"
	"flag"
	"time"
)

//函数入口
func main() {
	P(time.Now())

	Init()
	
	var projects string

	flag.StringVar(&projects, "p", "kubernetes-helm", "please enter docker-registry-project-list seperated by a comma.")
	flag.Parse()
	P("projects:", projects)
	projectSplit := strings.Split(strings.TrimSpace(projects), ",")
	startWork(projectSplit)
	
	// 执行失败的命令会在这里打印
	fmt.Println("--------The following is the command that failed to execute:")
	close(PullOrPushFailedChan)
	fmt.Println("--------Disply failed command end.")
	for v := range PullOrPushFailedChan {
		fmt.Println(v)
	}

	P(time.Now())
}
