# [hotbuild](https://hotbuild.ffactory.org)

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/wandercn/hotbuild?color=peru)](https://github.com/wandercn/hotbuild/releases/latest)
[![Released API docs](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white)](github.com/wandercn/hotbuild)
[![Build](https://img.shields.io/github/actions/workflow/status/wandercn/hotbuild/.github/workflows/go.yml?branch=master)](#)
[![Go Report Card](https://goreportcard.com/badge/github.com/wandercn/hotbuild?color=pink)](https://goreportcard.com/report/github.com/wandercn/hotbuild)
[![Lines of code](https://img.shields.io/tokei/lines/github/wandercn/hotbuild.svg?color=beige)](#)
[![Downloads of releases](https://img.shields.io/github/downloads/wandercn/hotbuild/total.svg?color=lavender)](https://github.com/wandercn/hotbuild/releases/latest)
[![Languages](https://img.shields.io/github/languages/top/wandercn/hotbuild.svg?color=yellow)](#)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/wandercn/hotbuild)](#)

A cross platform hot compilation tool

By monitoring the modification of the project directory file, the recompilation and running are automatically triggered. Running directory and monitoring change directory can be different. Monitor the file modification in all recursive subdirectories under the project path.

<!-- ![hotbuild](./logo.jpg =100x100) -->
 <img src="./logo.jpg" width = "100" height = "100" alt="Hotbuild" align=center />

# Install

    go get -u github.com/wandercn/hotbuild@latest 

# Initialize
```
wander$ cd $your_project_dir/
wander$ hotbuild initconf
2021/07/28 11:10:03 create config file: .hotbuild.toml
```

# Run
```
wander$ cd $your_project_dir/
wander$ hotbuild run
 _   _  ___ _____   ____  _   _ ___ _     ____
| | | |/ _ \_   _| | __ )| | | |_ _| |   |  _ \
| |_| | | | || |   |  _ \| | | || || |   | | | |
|  _  | |_| || |   | |_) | |_| || || |___| |_| |
|_| |_|\___/ |_|   |____/ \___/|___|_____|____/ v1.0.5, built with Go 1.16.5

.............................. ( Start rebuilding ) .................................
.............................. [ Build successfully ] ...............................
.............................. { Start running } ....................................
2021/07/06 09:56:26.913 [I] [asm_amd64.s:1371]  http server Running on http://0.0.0.0:8080
2021/07/06 09:56:26.913 [I] [asm_amd64.s:1371]  Admin server Running on :8088
```
# Custom configuration

Auto create default configuration file  in $your_project_dir/.hotbuild.toml,you can modify the configuration yourself.

```
wander$ cd $your_project_dir/
wander$ cat .hotbuild.toml
buildcmd = "go build -o ./tmp_bin"
excludedir = [".git", "tmp", "docs", "vendor"]
projectdir = "/Users/apple/workspace/src/example.com/"
runcmd = "./tmp_bin"

```
|           | en                           |         zh-cn          |
|:-----     |:-----                      |:-----            |
|buildcmd   | Compile command            |编译命令          |
|excludedir | Exclude tracked folders    |排除跟踪的文件夹  |
|projectdir | Project directory to track |需要跟踪的项目目录|
|runcmd     | Run command                |运行命令          |

# FAQ
1. the "Too many open files" error of MacOSX
   
  ```
# maxfiles is 256 too small.

wander$ launchctl limit
	cpu         unlimited      unlimited
	filesize    unlimited      unlimited
	data        unlimited      unlimited
	stack       8388608        67104768
	core        0              unlimited
	rss         unlimited      unlimited
	memlock     unlimited      unlimited
	maxproc     2784           4176
	maxfiles    256           10240

# change maxfiles to 4096.

wander$ sudo launchctl limit maxfiles 4096 unlimited
wander$ launchctl limit
	cpu         unlimited      unlimited
	filesize    unlimited      unlimited
	data        unlimited      unlimited
	stack       8388608        67104768
	core        0              unlimited
	rss         unlimited      unlimited
	memlock     unlimited      unlimited
	maxproc     2784           4176
	maxfiles    4096           10240

  ```
remind:
> Some terminal modifications cannot take effect immediately and need to be restarted, such as iterm2.
