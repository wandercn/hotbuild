/*
   Copyright (c) 2021-present ffactory.org
   hotbuild is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
   EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
   MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/
package watch

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"

	css "github.com/flylog/colorstyle"
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
		lineFormat := "\n                       >> %s << \n"
		css.New().StyleBold().BgGreen().Printf(lineFormat, "( Start rebuilding )")
		if err = run.BuildCode(); err != nil {
			log.Printf("BuildCode failed: %v", err)
			return
		}
		// 重新编译运行之前退出之前的进程
		if currentPid > 0 {
			css.New().StyleBold().BgRed().Printf(lineFormat, fmt.Sprintf("< GracefulStop running on pid=%v >", currentPid))
			proc, err := os.FindProcess(currentPid)
			if err != nil {
				log.Printf("find old proc failed: %v", err)
			} else {
				if runtime.GOOS == "windows" {
					err = proc.Kill() //windows不支持Signal信号只能Kill
					if err != nil {
						os.Exit(1)
						panic(err)
					}

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
		}
		css.New().StyleBold().BgMagenta().Printf(lineFormat, "[ Build successfully ]")
		css.New().StyleBold().BgCyan().Printf(lineFormat, "{ Start running }")
		if currentPid, err = run.Run(); err != nil {
			log.Printf("Run Failed:%v", err)
			return
		}
		css.New().StyleItalic().ColorGreen().Println("Hotbuild is running. Press Ctrl+C to stop")
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
			log.Fatalf(css.New().Sprintf("add watch dir: %s failed: %v\n", d, err))
			return
		}
	}

	signal.Notify(waited, os.Kill)
	// ctrl + C 强制退出
	<-waited
}
