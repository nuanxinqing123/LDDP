# 编译前；一定需要打开Docker Desktop，否则会报错
# 编译前；一定需要打开Docker Desktop，否则会报错
# 编译前；一定需要打开Docker Desktop，否则会报错

# 编译前端文件
cd /Users/nuanxinqing/Code/Vue/arco-lddp-vue
yarn build

# 返回项目目录
cd /Users/nuanxinqing/Code/Golang/LDDP

# 删除现有前端文件
rm -rf static/assets/*

# 复制前端文件到项目目录
cp -r /Users/nuanxinqing/Code/Vue/arco-lddp-vue/dist/* static/assets/

# 打包前端文件
cd static
go-bindata -o=bindata/bindata.go -pkg=bindata ./assets/...
cd ..

# 打包普通版
cd server/const
gsed -i '13s/const EmpowerVersion = "v2"/const EmpowerVersion = "v1"/' const.go
cd ../..
xgo -out LDDP-V1 --targets=linux/amd64,windows/amd64 ./

# 打包订阅版
cd server/const
gsed -i '13s/const EmpowerVersion = "v1"/const EmpowerVersion = "v2"/' const.go
cd ../..
xgo -out LDDP-V2 --targets=linux/amd64,windows/amd64 ./

# 压缩打包文件
upx LDDP-*