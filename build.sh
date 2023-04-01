# 编译前端文件
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