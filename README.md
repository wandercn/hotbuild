# hotbuild

A cross platform hot compilation tool

By monitoring the modification of the project directory file, the recompilation and running are automatically triggered. Running directory and monitoring change directory can be different.

# install

    go get -u github.com/wandercn/hotbuild@latest 

# run
```
wander$ cd $your_project_dir/
wander$ hotbuild
 _   _  ___ _____   ____  _   _ ___ _     ____
| | | |/ _ \_   _| | __ )| | | |_ _| |   |  _ \
| |_| | | | || |   |  _ \| | | || || |   | | | |
|  _  | |_| || |   | |_) | |_| || || |___| |_| |
|_| |_|\___/ |_|   |____/ \___/|___|_____|____/ v1.0.0, built with Go 1.16.5

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
