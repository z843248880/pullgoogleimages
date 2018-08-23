package main

import (
	"fmt"
)


func Init() {
	initMap()
	initLinuxOperation()
}


func initLinuxOperation() {
	gitcmd := "[ -d /data/git_root/gcr.io_mirror ] && rm -rf /data/git_root/gcr.io_mirror ; git clone https://github.com/anjia0532/gcr.io_mirror.git /data/git_root/gcr.io_mirror"
	err := runBashCommandAndKillIfTooSlow(gitcmd, 120)
	if err == "occurError" {
		fmt.Printf("exec linux command error: %s", gitcmd)
	}
	
}

func initMap() {
	//不想同步哪些镜像，直接在这里注释即可；如果要添加新的项目或者镜像，也是在这添加
	NeedPullImagelistGoogleContainers["busybox"] = ""
	NeedPullImagelistGoogleContainers["cadvisor"] = ""
	NeedPullImagelistGoogleContainers["cassandra"] = ""
	NeedPullImagelistGoogleContainers["cloud-controller-manager"] = ""
	NeedPullImagelistGoogleContainers["coredns"] = ""
	NeedPullImagelistGoogleContainers["etcd"] = ""
	NeedPullImagelistGoogleContainers["etcd-amd64"] = ""
	NeedPullImagelistGoogleContainers["ingress-gce-glbc-amd64"] = ""
	NeedPullImagelistGoogleContainers["kube2sky"] = ""
	NeedPullImagelistGoogleContainers["kube-apiserver"] = ""
	NeedPullImagelistGoogleContainers["kube-apiserver-amd64"] = ""
	NeedPullImagelistGoogleContainers["kube-controller-manager"] = ""
	NeedPullImagelistGoogleContainers["kube-controller-manager-amd64"] = ""
	NeedPullImagelistGoogleContainers["kubectl"] = ""
	NeedPullImagelistGoogleContainers["kubedns-amd64"] = ""
	NeedPullImagelistGoogleContainers["kube-proxy"] = ""
	NeedPullImagelistGoogleContainers["kube-proxy-amd64"] = ""
	NeedPullImagelistGoogleContainers["kubernetes-dashboard"] = ""
	NeedPullImagelistGoogleContainers["kubernetes-dashboard-amd64"] = ""
	NeedPullImagelistGoogleContainers["kube-scheduler"] = ""
	NeedPullImagelistGoogleContainers["kube-scheduler-amd64"] = ""
	NeedPullImagelistGoogleContainers["pause"] = ""
	NeedPullImagelistGoogleContainers["pause-amd64"] = ""

	NeedPullImagelistKubernetesHelm["expandybird"] = ""
	NeedPullImagelistKubernetesHelm["manager"] = ""
	NeedPullImagelistKubernetesHelm["resourcifier"] = ""
	NeedPullImagelistKubernetesHelm["tiller"] = ""

	// NeedPullImagelistGoogleSamples["appengine-tensorboard"] = ""
	// NeedPullImagelistGoogleSamples["cassandra"] = ""
	// NeedPullImagelistGoogleSamples["container-analysis-webhook"] = ""
	// NeedPullImagelistGoogleSamples["echo-go"] = ""
	// NeedPullImagelistGoogleSamples["echo-java"] = ""
	// NeedPullImagelistGoogleSamples["echo-node"] = ""
	// NeedPullImagelistGoogleSamples["echo-php"] = ""
	// NeedPullImagelistGoogleSamples["echo-python"] = ""
	// NeedPullImagelistGoogleSamples["echo-ruby"] = ""
	// NeedPullImagelistGoogleSamples["env-backend"] = ""
	// NeedPullImagelistGoogleSamples["env-show"] = ""
	// NeedPullImagelistGoogleSamples["freshpod"] = ""
	// NeedPullImagelistGoogleSamples["gb-frontend"] = ""
	// NeedPullImagelistGoogleSamples["gb-frontend-amd64"] = ""
	// NeedPullImagelistGoogleSamples["gb-frontend-arm"] = ""
	// NeedPullImagelistGoogleSamples["gb-frontend-arm64"] = ""
	// NeedPullImagelistGoogleSamples["gb-frontend-ppc64le"] = ""
	// NeedPullImagelistGoogleSamples["gb-redisslave"] = ""
	// NeedPullImagelistGoogleSamples["gb-redisslave-amd64"] = ""
	// NeedPullImagelistGoogleSamples["gb-redisslave-arm"] = ""
	// NeedPullImagelistGoogleSamples["gb-redisslave-arm64"] = ""
	// NeedPullImagelistGoogleSamples["gb-redisslave-ppc64le"] = ""
	// NeedPullImagelistGoogleSamples["gb-redisworker"] = ""
	// NeedPullImagelistGoogleSamples["gke-serviceaccounts-initializer"] = ""
	// NeedPullImagelistGoogleSamples["go-echo"] = ""
	// NeedPullImagelistGoogleSamples["grpc-cxx"] = ""
	// NeedPullImagelistGoogleSamples["hello-app"] = ""
	// NeedPullImagelistGoogleSamples["hello-app-metrics"] = ""
	// NeedPullImagelistGoogleSamples["hello-app-tls"] = ""
	// NeedPullImagelistGoogleSamples["hello-frontend"] = ""
	// NeedPullImagelistGoogleSamples["hello-go-gke"] = ""
	// NeedPullImagelistGoogleSamples["k8sdocs"] = ""
	// NeedPullImagelistGoogleSamples["k8s-filebeat"] = ""
	// NeedPullImagelistGoogleSamples["k8skafka"] = ""
	// NeedPullImagelistGoogleSamples["k8szk"] = ""
	// NeedPullImagelistGoogleSamples["kubernetes-bootcamp"] = ""
	// NeedPullImagelistGoogleSamples["kubernetes-zookeeper"] = ""
	// NeedPullImagelistGoogleSamples["lobsters"] = ""
	// NeedPullImagelistGoogleSamples["lobsters-db"] = ""
	// NeedPullImagelistGoogleSamples["ml-pipeline-cmle-op"] = ""
	// NeedPullImagelistGoogleSamples["ml-pipeline-dataflow-tfma-taxi"] = ""
	// NeedPullImagelistGoogleSamples["ml-pipeline-dataflow-tftbq-taxi"] = ""
	// NeedPullImagelistGoogleSamples["ml-pipeline-kubeflow-tfserve-taxi"] = ""
	// NeedPullImagelistGoogleSamples["ml-pipeline-kubeflow-tf-taxi"] = ""
	// NeedPullImagelistGoogleSamples["ml-pipeline-kubeflow-trainer-taxi"] = ""
	NeedPullImagelistGoogleSamples["mysql"] = ""
	// NeedPullImagelistGoogleSamples["nfs-server"] = ""
	// NeedPullImagelistGoogleSamples["node-hello"] = ""
	// NeedPullImagelistGoogleSamples["prometheus-dummy-exporter"] = ""
	// NeedPullImagelistGoogleSamples["pubsub-bq-pipe"] = ""
	// NeedPullImagelistGoogleSamples["pubsub-sample"] = ""
	// NeedPullImagelistGoogleSamples["README.md"] = ""
	// NeedPullImagelistGoogleSamples["redis"] = ""
	// NeedPullImagelistGoogleSamples["redis-bq-pipe"] = ""
	// NeedPullImagelistGoogleSamples["redis-proxy"] = ""
	// NeedPullImagelistGoogleSamples["sd-dummy-exporter"] = ""
	// NeedPullImagelistGoogleSamples["steps-twotier"] = ""
	// NeedPullImagelistGoogleSamples["style-transfer"] = ""
	// NeedPullImagelistGoogleSamples["tf-jupyter-server"] = ""
	// NeedPullImagelistGoogleSamples["tf-k8s-worker"] = ""
	// NeedPullImagelistGoogleSamples["tf-worker-example"] = ""
	// NeedPullImagelistGoogleSamples["tf-workshop"] = ""
	// NeedPullImagelistGoogleSamples["twilio-vision"] = ""
	// NeedPullImagelistGoogleSamples["xtrabackup"] = ""
	// NeedPullImagelistGoogleSamples["zone-printer"] = ""

	NeedPullImagelist["google-containers"] = NeedPullImagelistGoogleContainers
	NeedPullImagelist["kubernetes-helm"]   = NeedPullImagelistKubernetesHelm
	NeedPullImagelist["google-samples"]    = NeedPullImagelistGoogleSamples
}
