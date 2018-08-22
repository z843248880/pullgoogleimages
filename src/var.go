package main

import (
	"time"
)

//变量
var (
	//拉取镜像的地址，即源地址
	ImageURL                  			= "anjia0532"

	//推送镜像地址，即目标地址；记住必须在服务器上先docker login登录到docker仓库，否则推送会报错。
	PullURL 							= "wcr.wecash.net"

	//服务器上存放镜像的目录
	RootDir                   			= "/data/git_root/gcr.io_mirror"

	//线程数，根据服务器配置调整，最好别超过10个，否则网络IO和磁盘IO会遇到瓶颈，导致拉取和推送镜像都很慢很慢。
	PollNum								= 10
	
	//以下参数不用动 
	AllRegistryWithTagChan    			= make(chan []string, 100000)
	PullOrPushFailedChan      			= make(chan string, 10000)
	T                         			= time.NewTicker(time.Second * 3)
	StartWorkChan             			= make(chan bool, 100)
	StartWorkChanTag          			= 0 
	NeedPullImagelist         			= make(map[string]map[string]string)
	NeedPullImagelistGoogleContainers   = make(map[string]string)
	NeedPullImagelistKubernetesHelm   		= make(map[string]string)
	NeedPullImagelistGoogleSamples   	= make(map[string]string)
)
