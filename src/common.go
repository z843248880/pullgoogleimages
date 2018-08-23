package main

import (
	// "log"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
	// "strconv"
)



func checkErr(err error, shell_cmd string) string {
	if err != nil {
		// log.Fatalf("exec shell command '%s' occur error: %v", shell_cmd, err)
		fmt.Printf("exec shell command '%s' occurError: %v \n", shell_cmd, err)
		return "occurError"
	}
	return ""
}

func P(s ...interface{}) {
	fmt.Println(time.Now().Format("2006-1-2 15:04:05"), "： ", s)
}

func exec_shell(shell_cmd string) string {
	cmd := exec.Command("/bin/bash", "-c", shell_cmd)
	out, err := cmd.Output()
	if errResult := checkErr(err, shell_cmd); errResult == "occurError" {
		return "occurError"
	} else {
		return strings.Trim(string(out), "\n")
	}

}

//执行linux命令并返回结果，限制命令超时时间，命令超时会被kill
func runBashCommandAndKillIfTooSlow(command string, killInSeconds time.Duration) string {
	cmd := exec.Command("/bin/bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Start()
	if err != nil {
		P("start process error: ", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	tt := time.NewTicker(killInSeconds * time.Second)
	select {
	// case <-time.After(killInSeconds * time.Second):
	//     if err := cmd.Process.Kill(); err != nil {
	//         log.Fatal("failed to kill: ", err , "command is: ", command)
	//     } else {
	//         P("kill the process that exceeds timeout value: ", command)
	//     }
	//     <-done
	case <-tt.C:
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill: ", err, "command is: ", command)
		} else {
			P("kill the process that exceeds timeout value: ", command)
			return "occurError"
		}
		<-done
	case err := <-done:
		if err != nil {
			// log.Printf("process done with error = %v", err)
			return "occurError"
		} else {
			return strings.Trim(out.String(), "\n")
		}
	}
	return "occurError"
}
