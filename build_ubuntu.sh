#!/usr/bin/env bash
: ' 
   Copyright (c) 2021-present ffactory.org
   hotbuild is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
   EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
   MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
'
lastTag=$(git describe --tags $(git rev-list --tags --max-count=1))
goVersion=$(go version | awk '{print $3'})
versionFile=./version/version.go
release=bin
echo $lastTag
echo $goVersion
rm -f $release/*
# 更新版本号
sed -i "s/const Version = \"*.*.*\"/const Version = \"$lastTag\"/" ${versionFile}

# 更新go版本
sed -i "s/const GoVersion = \"*.*.*\"/const GoVersion = \"$goVersion\"/" ${versionFile}
# Linux amd64
GO_ENABLED=0 GOOS=linux GOARCH=amd64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'version.Version=$lastTag'" -o ./$release/$target/
cd $release/
zip -mr $target.zip $target
cd ..
# Linux arm
GO_ENABLED=0 GOOS=linux GOARCH=arm
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-X 'version.Version=$lastTag'" -o ./$release/$target/
cd $release/
zip -mr $target.zip $target
cd ..
# Linux arm64
GO_ENABLED=0 GOOS=linux GOARCH=arm64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-X 'version.Version=$lastTag'" -o ./$release/$target/
cd $release/
zip -mr $target.zip $target
cd ..
# Linux loong64
GO_ENABLED=0 GOOS=linux GOARCH=loong64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=linux GOARCH=loong64 go build -ldflags="-X 'version.Version=$lastTag'" -o ./$release/$target/
cd $release/
zip -mr $target.zip $target
cd ..
# Linux riscv64
GO_ENABLED=0 GOOS=linux GOARCH=riscv64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -ldflags="-X 'version.Version=$lastTag'" -o ./$release/$target/
cd $release/
zip -mr $target.zip $target
cd ..
# macosx amd64
GO_ENABLED=0 GOOS=darwin GOARCH=amd64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-X 'version.Version=$lastTag'" -o ./$release/$target/
cd $release/
zip -mr $target.zip $target
cd ..
# macosx arm64
GO_ENABLED=0 GOOS=darwin GOARCH=arm64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-X 'version.Version=$lastTag'" -o ./$release/$target/
cd $release/
zip -mr $target.zip $target
cd ..
# windows
GO_ENABLED=0 GOOS=windows GOARCH=amd64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-X 'version.Version=$lastTag'" -o ./$release/$target/
cd $release/
zip -mr $target.zip $target
cd ..

# md5
cd $release/
echo "" >md5checksum.txt
for i in $(ls *.zip); do
	md5sum $i >>md5checksum.txt
done
cd ..
cat $release/md5checksum.txt
