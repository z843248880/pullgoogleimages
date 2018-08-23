# pullgoogleimages
This project is used to pull google official docker images.It is can pull the newest tag from all tags.  
该项目用于一键拉取google相关docker镜像并推送到私有docker仓库  

# 代码文件说明
- common.go：       存放执行linux命令的相关代码
- getimageinfo.go： 存放处理镜像的相关代码
- getnewtag.go：    从众多镜像tag中选取可能性最大的最新tag
- init.go：         初始化操作；initMap()方法可以控制想同步的项目和镜像
- main.go：         程序入口
- var.go：          相关全局变量；修改镜像源或者目标镜像地址，修改该文件的ImageURL和PullURL即可；请把PullURL修改为你的私有docker仓库地址

# 前提
1. 安装go环境  
请自行搜索安装go，版本不限；建议：go version go1.10.3 linux/amd64  
将go运行程序移动到/usr/bin/目录下  
执行go version验证go命令是否可用。 
2. 在服务器上登录私有docker仓库，并在私有docker仓库上创建项目:kubernetes-helm、google-samples、google-containers，想同步哪些镜像项目就在私有docker仓库创建哪些项目


# 运行
mkdir -p /data/gopro/src &&  \
cd /data/gopro/src &&  \
git clone https://github.com/z843248880/pullgoogleimages.git &&  \
cd /data/gopro/src/pullgoogleimages/src &&  \
go build . &&  \
./src -p kubernetes-helm,google-samples,google-containers

注："-p"参数是可选的，即指定多个项目，不加“-p”参数默认只同步kubernetes-helm一个项目。
可在服务器上执行“ls /data/git_root/gcr.io_mirror/”列出所有可同步项目。
