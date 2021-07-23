/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : watch.go
#   Last Modified : 2021-07-23 11:09
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

func Start() {
	// 运行中的进程id
	var currentPid int

	hotBuild := func() {
		var err error
		fmt.Println(".............................. ( Start rebuilding ) .................................")
		if err = run.BuildCode(); err != nil {
			log.Printf("BuildCode failed: %v", err)
			return
		}
		// 重新编译运行之前退出之前的进程
		if currentPid > 0 {
			fmt.Printf(".............................. < GracefulStop running on pid=%v > ................\n", currentPid)
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
		fmt.Println(".............................. [ Build successfully ] ...............................")
		fmt.Println(".............................. { Start running } ....................................")
		if currentPid, err = run.Run(); err != nil {
			log.Printf("Run Failed:%v", err)
			return
		}
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
		log.Printf("GetProjectDir failed: %v", err)
		return
	}
	watchDirs := make([]string, 100)
	excludeDirs, err := conf.GetExcludeDirs()
	if err != nil {
		log.Printf("GetExcludeDirs failed: %v", err)
		return
	}
	err = tree.TreeDirs(projectDir, &watchDirs, excludeDirs)
	if err != nil {
		log.Fatalf("tree dirs failed: %v", err)
		return
	}
	for _, d := range watchDirs {
		err = watcher.Add(d)
		if err != nil {
			log.Fatalf("add watch dir: %s failed: %v", d, err)
			return
		}
	}
	log.Println("Hotbuild is running. Press Ctrl+C to stop")
	signal.Notify(waited, os.Interrupt, os.Kill)
	// ctrl + C 强制退出
	<-waited
}
