###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-16 02:26:42
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-16 02:52:34
 # @FilePath: /pb/user/compile.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -x
go mod init github.com/kokutas/pb/user
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   user.proto
go mod tidy