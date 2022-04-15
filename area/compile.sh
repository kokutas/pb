###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-16 02:46:13
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-16 03:01:33
 # @FilePath: /area/compile.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -x
go mod init github.com/kokutas/pb/area
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   area.proto
   
go mod tidy
go mod vendor