#!/usr/bin/env bash
lastTag=$(git describe --tags `git rev-list --tags --max-count=1`)
echo $lastTag > version.txt
# Linux amd64
GO_ENABLED=0 GOOS=linux GOARCH=amd64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=$lastTag'" -o ./build/$target/ 
cd build/
zip -mr  $target.zip $target 
cd ..
# Linux arm
GO_ENABLED=0 GOOS=linux GOARCH=arm
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-X 'main.Version=$lastTag'" -o ./build/$target/ 
cd build/
zip -mr  $target.zip $target 
cd ..
# macosx
GO_ENABLED=0 GOOS=darwin GOARCH=amd64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-X 'main.Version=$lastTag'" -o ./build/$target/ 
cd build/
zip -mr  $target.zip $target 
cd ..
# windows
GO_ENABLED=0 GOOS=windows GOARCH=amd64
target="hotbuild_${lastTag}_${GOOS}_${GOARCH}"
GO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.Version=$lastTag'" -o ./build/$target/ 
cd build/
zip -mr  $target.zip $target 
cd ..

# md5
cd build/
echo "" > md5checksum.txt
for i in `ls *.zip`;
do
    md5 $i >> md5checksum.txt
done;
cd ..
cat build/md5checksum.txt
