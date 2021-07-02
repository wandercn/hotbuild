/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : run.go
#   Last Modified : 2021-07-02 15:31
#   Describe      :
#
# ====================================================*/

package run

import (
	"fmt"
	"hotbuild/config"
	"os"
	"os/exec"
	"strings"
)

/* 编译代码 */
func BuildCode() error {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Getwd failed: %v", err)
	}
	err = os.Chdir(pwd)
	if err != nil {
		return fmt.Errorf("Chdir failed: %v", err)
	}
	var cmd *exec.Cmd
	c := config.New()

	buildCmd, err := c.GetBuildCmd()
	if err != nil {
		return fmt.Errorf("GetBuildCmd failed: %v", err)
	}
	args := strings.Split(buildCmd, " ")
	command := args[0]
	cmd = exec.Command(command, args[1:]...)
	cmd.Env = os.Environ()

	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("CombinedOutput failed: %v; output: %s\n", err, b)
	}
	return nil
}

/* 运行编译文件并返回进行pid */
func Run() (pid int, err error) {
	pwd, err := os.Getwd()
	if err != nil {
		return -1, fmt.Errorf("Getwd failed: %v", err)
	}
	err = os.Chdir(pwd)
	if err != nil {
		return -1, fmt.Errorf("Chdir failed: %v", err)
	}
	var cmd *exec.Cmd
	c := config.New()
	runCmd, err := c.GetRunCmd()
	if err != nil {
		return -1, fmt.Errorf("GetBuildCmd ailed: %v", err)
	}
	goPath := os.ExpandEnv("$GOPATH")
	runCmd = strings.ReplaceAll(runCmd, "$GOPATH", goPath)

	args := strings.Split(runCmd, " ")
	command := strings.Join(args[:1], " ")
	cmd = exec.Command(command, args[1:]...)
	cmd.Env = os.Environ()
	// 输出运行日志到终端
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return -1, fmt.Errorf("CombinedOutput failed: %v", err)
	}
	return cmd.Process.Pid, nil
}
