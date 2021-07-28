#!/usr/bin/env bash
lastTag=$(git describe --tags `git rev-list --tags --max-count=1`)
goVersion=$(go version | awk '{print $3'})
versionFile=version/version.go
release=bin
echo $lastTag
echo $goVersion
# 更新版本号
sed -ie "s/const Version = \"*.*.*\"/const Version = \"$lastTag\"/" $versionFile 

# 更新go版本
sed -ie "s/const GoVersion = \"*.*.*\"/const GoVersion = \"$goVersion\"/" $versionFile 
# Linux amd64
GO_ENABLED=0 GOOS=linux GOARCH=amd64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=$lastTag'" -o ./$release/$target/ 
cd $release/ 
zip -mr  $target.zip $target 
cd ..
# Linux arm
GO_ENABLED=0 GOOS=linux GOARCH=arm
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-X 'main.Version=$lastTag'" -o ./$release/$target/ 
cd $release/
zip -mr  $target.zip $target 
cd ..
# macosx
GO_ENABLED=0 GOOS=darwin GOARCH=amd64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-X 'main.Version=$lastTag'" -o ./$release/$target/ 
cd $release/
zip -mr  $target.zip $target 
cd ..
# windows
GO_ENABLED=0 GOOS=windows GOARCH=amd64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.Version=$lastTag'" -o ./$release/$target/ 
cd $release/
zip -mr  $target.zip $target 
cd ..

# md5
cd $release/
echo "" > md5checksum.txt
for i in `ls *.zip`;
do
    md5 $i >> md5checksum.txt
done;
cd ..
cat $release/md5checksum.txt
