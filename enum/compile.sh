###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-16 03:45:06
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-16 03:45:07
 # @FilePath: /pb/enum/compile.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
#!/bin/sh
set -x
go mod init github.com/kokutas/pb/enum
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   enum.proto
go mod tidy