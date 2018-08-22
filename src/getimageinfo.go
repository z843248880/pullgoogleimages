package main

import (
	"strings"
)

func startWork(projectSplit []string) {
	var mainExitChan chan bool

	mainExitChan = make(chan bool, len(projectSplit))
	for _, image_project_dir := range projectSplit {
		P("image_project_dir:", image_project_dir)
		go StartWorkFromProject(image_project_dir, mainExitChan)
	}

	P("len(projectSplit):", len(projectSplit))
	for i := 0; i < len(projectSplit); i++ {
		a := <-mainExitChan
		P("mainExitChan", a)
	}
	close(AllRegistryWithTagChan)

	for i := 0; i < PollNum; i++ {
		go PullImageAndPushImage()
	}

	for i := 0; i < PollNum; i++ {
		a := <-StartWorkChan
		P("StartWorkChan:", a)
		P("StartWorkChanTag:", StartWorkChanTag)
	}

	// DONE:
	// 	for {
	// 		select {
	// 		case v := <-AllRegistryWithTagChan:
	// 			P("tag:", v)
	// 			AllRegistryWithTagChanNum++
	// 		// case <- time.After(time.Second * 3):
	// 		case <-T.C:
	// 			P("process exit")
	// 			close(AllRegistryWithTagChan)
	// 			T.Stop()
	// 			break DONE
	// 		}
	// 	}

	// P("AllRegistryWithTagChanNum:", AllRegistryWithTagChanNum)

}

//从channel中取出数据，然后执行docker pull 、docker tag 、docker push 。
func PullImageAndPushImage() {
	for v := range AllRegistryWithTagChan {
		projectName := v[0]
		imageName := v[1]
		tagValue := v[2]
		tagValue = strings.TrimSuffix(tagValue, ".tag")
		fullImagePath := ImageURL + "/" + projectName + "." + imageName + ":" + tagValue
		pullCmd := "docker pull " + fullImagePath

		P(pullCmd + "    start.")
		execResult := runBashCommandAndKillIfTooSlow(pullCmd, 600)
		if execResult != "occurError" {
			P(pullCmd + "    done.")
			PushFullImagePath := PullURL + "/" + projectName + "/" + imageName + ":" + tagValue
			_ = runBashCommandAndKillIfTooSlow("docker tag "+fullImagePath+" "+PushFullImagePath, 10)
			pushCmd := "docker push " + PushFullImagePath
			P(pushCmd + "    start.")
			P("pushstart", pushCmd)
			execResult = runBashCommandAndKillIfTooSlow(pushCmd, 600)
			P("pushend.", pushCmd)
			if execResult != "occurError" {
				P(pushCmd + "    done.")
			} else {
				PullOrPushFailedChan <- pushCmd
			}
		} else {
			PullOrPushFailedChan <- pullCmd
		}
	}
	StartWorkChan <- true
	StartWorkChanTag++
}

func (in *ImageInfo) GetImageName(rootdir, imageprojectdir string) ([]string, int) {
	fullPathProject := rootdir + "/" + imageprojectdir
	project := runBashCommandAndKillIfTooSlow("ls -l "+fullPathProject+" | grep '^d' | awk '{print $NF}'", 60)
	projectList := strings.Split(project, "\n")
	projectLen := len(projectList)
	P("len:", projectLen)
	return projectList, projectLen
}

func StartWorkFromProject(image_project_dir string, mainExitChan chan bool) {
	var in ImageInfo
	projectList, _ := in.GetImageName(RootDir, image_project_dir)

	var projectListChan chan []string
	projectListChan = make(chan []string, 10)

	go func() {
		for _, v := range projectList {
			t := []string{RootDir, "/", image_project_dir, "/", v}
			projectListChan <- t
		}
		close(projectListChan)
	}()

	var GetImageTagExitChan chan bool
	GetImageTagExitChan = make(chan bool, 10000)
	in.GetImageTag(projectListChan, GetImageTagExitChan)

	for i := 0; i < len(projectList); i++ {
		a := <-GetImageTagExitChan
		P("GetImageTagExitChan:", a)
	}
	close(GetImageTagExitChan)

	P("image_project_dir:", image_project_dir, "    ", "len(projectList):", len(projectList))

	mainExitChan <- true
}

//根据镜像获取该镜像拥有的tag，并把需要拉取的tag信息放入channel
func (in *ImageInfo) GetImageTag(projectListChan chan []string, GetImageTagExitChan chan bool) {
	for t := range projectListChan {
		image_project_dir := t[2]
		projectname := t[4]
		if _, ok := NeedPullImagelist[image_project_dir][projectname]; ok {
			v := RootDir + "/" + image_project_dir + "/" + projectname
			// image_tag_all := runBashCommandAndKillIfTooSlow("ls "+v+"/*.tag", 60)
			image_tag_all := runBashCommandAndKillIfTooSlow("ls "+v+"/*.tag | egrep -v 'beta|alpha'", 60)
			tagList := strings.Split(image_tag_all, "\n")
			newesttagvalue := GetNewestTag(tagList)
			if newesttagvalue != "nonetag" {
				P("newesttagvalue:",projectname , newesttagvalue)
				all_registry_with_tag_slice := []string{image_project_dir, projectname, newesttagvalue}
				AllRegistryWithTagChan <- all_registry_with_tag_slice
				// for _, fullpathtag := range tagList {
				// 	vList := strings.Split(fullpathtag, "/")
				// 	vListLen := len(vList)
				// 	tagValue := vList[vListLen-1]
				// 	all_registry_with_tag_slice := []string{image_project_dir, projectname, tagValue}
				// 	// P("tag new:", all_registry_with_tag_slice)
				// 	AllRegistryWithTagChan <- all_registry_with_tag_slice
				// }
			}
		}
		GetImageTagExitChan <- true
	}
}
