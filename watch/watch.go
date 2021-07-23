/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : watch.go
#   Last Modified : 2021-07-23 17:16
#   Describe      :
#
# ====================================================*/

package watch

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/fsnotify/fsnotify"
	"github.com/wandercn/hotbuild/config"
	"github.com/wandercn/hotbuild/run"
	"github.com/wandercn/hotbuild/tree"
)

var (
	GreeBg    = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	redBg     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	Green     = string([]byte{27, 91, 51, 50, 109})
	white     = string([]byte{27, 91, 51, 55, 109})
	yellow    = string([]byte{27, 91, 51, 51, 109})
	red       = string([]byte{27, 91, 51, 49, 109})
	blue      = string([]byte{27, 91, 51, 52, 109})
	magenta   = string([]byte{27, 91, 51, 53, 109})
	cyan      = string([]byte{27, 91, 51, 54, 109})
	Reset     = string([]byte{27, 91, 48, 109})
)

func Start() {
	// 运行中的进程id
	var currentPid int

	hotBuild := func() {
		var err error
		fmt.Println("")
		fmt.Println(".............................. ", GreeBg, "( Start rebuilding )", Reset, " .................................")
		if err = run.BuildCode(); err != nil {
			log.Printf("BuildCode failed: %v", err)
			return
		}
		// 重新编译运行之前退出之前的进程
		if currentPid > 0 {
			fmt.Println(".............................. ", redBg, "< GracefulStop running on pid=", currentPid, " >", Reset, " ................")
			proc, err := os.FindProcess(currentPid)
			if err != nil {
				log.Printf("find old proc failed: %v", err)
			} else {
				err := proc.Signal(os.Interrupt)
				if err != nil {
					log.Printf("Signal error: %v", err)
					log.Printf("gracefulStop doing kill")
					err = proc.Signal(os.Kill)
					if err != nil {
						os.Exit(1)
						panic(err)
					}
				}
			}
		}
		fmt.Println(".............................. ", magentaBg, "[ Build successfully ]", Reset, " ...............................")
		fmt.Println(".............................. ", cyanBg, "{ Start running }", Reset, " ....................................")
		if currentPid, err = run.Run(); err != nil {
			log.Printf("Run Failed:%v", err)
			return
		}
		fmt.Println(Green, "Hotbuild is running. Press Ctrl+C to stop", Reset)
		fmt.Printf("\n")
	}

	hotBuild()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	waited := make(chan os.Signal, 1)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create && tree.IsExistDir(event.Name) {
					err = watcher.Add(event.Name)
					if err != nil {
						log.Printf("added Add watch failed:%v\n", err)
					}

					log.Println("added file:", event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					hotBuild()
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove && tree.IsExistDir(event.Name) {
					err = watcher.Remove(event.Name)
					if err != nil {
						log.Printf("deleted Remove watch failed: %v\n", err)
					}
					log.Println("deleted file:", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename && tree.IsExistDir(event.Name) {
					err = watcher.Remove(event.Name)
					if err != nil {
						log.Printf("rename Remove watch failed: %v\n", err)
					}
					log.Println("rename file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	conf := config.New()
	projectDir, err := conf.GetProjectDir()
	if err != nil {
		log.Printf("GetProjectDir failed: %v\n", err)
		return
	}
	watchDirs := make([]string, 100)
	excludeDirs, err := conf.GetExcludeDirs()
	if err != nil {
		log.Printf("GetExcludeDirs failed: %v\n", err)
		return
	}
	err = tree.TreeDirs(projectDir, &watchDirs, excludeDirs)
	if err != nil {
		log.Fatalf("tree dirs failed: %v\n", err)
		return
	}
	for _, d := range watchDirs {
		err = watcher.Add(d)
		if err != nil {
			log.Fatalf(redBg, "add watch dir: %s failed: %v\n", d, err, Reset)
			return
		}
	}

	signal.Notify(waited, os.Kill)
	// ctrl + C 强制退出
	<-waited
}
