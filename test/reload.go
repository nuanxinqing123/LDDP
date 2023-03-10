package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("等待20秒后，我会自动重启")
	//fmt.Println("这是第一版")
	fmt.Println("这是第NNNNNNNN版")
	time.Sleep(time.Second * 20)

	// 获取当前程序的命令行参数
	args := os.Args

	// 执行外部命令，并替换当前进程
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 调用命令，如果发生错误，记录错误信息
	if err := cmd.Run(); err != nil {
		log.Println(err)
		fmt.Println("程序自我重启失败！")
	}
}
