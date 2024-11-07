# 编译前；一定需要打开Docker Desktop，否则会报错
# 编译前；一定需要打开Docker Desktop，否则会报错
# 编译前；一定需要打开Docker Desktop，否则会报错

# 编译前端文件
cd /Users/nuanxinqing/Code/Vue/arco-lddp-vue || exit
yarn build

# 返回项目目录
cd /Users/nuanxinqing/Code/Golang/LDDP || exit

# 删除现有前端文件
rm -rf static/assets/*

# 复制前端文件到项目目录
cp -r /Users/nuanxinqing/Code/Vue/arco-lddp-vue/dist/* static/assets/

# 打包前端文件
cd static || exit
go-bindata -o=bindata/bindata.go -pkg=bindata ./assets/...
cd ..

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o LDDP-windows-amd64 -ldflags '-s -w -extldflags "-static"'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o LDDP-linux-amd64.exe -ldflags '-s -w -extldflags "-static"'

# 压缩打包文件
upx LDDP-*