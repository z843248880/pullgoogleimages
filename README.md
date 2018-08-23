# pullgoogleimages
This project is used to pull google official docker images.It is can pull the newest tag from all tags.  

# 前提
安装go环境  
请自行搜索安装go，版本不限；建议：go version go1.10.3 linux/amd64  
将go运行程序移动到/usr/bin/目录下  
执行go version验证go命令是否可用。  

# 运行
mkdir -p /data/gopro/src &&  \
cd /data/gopro/src &&  \
git clone https://github.com/z843248880/pullgoogleimages.git &&  \
cd /data/gopro/src/pullgoogleimages/src &&  \
go build . &&  \
./src -p kubernetes-helm,google-samples,google-containers

注："-p"参数是可选的，即指定多个项目，不加“-p”参数默认只同步kubernetes-helm一个项目。
可在服务器上执行“ls /data/git_root/gcr.io_mirror/”列出所有可同步项目。
