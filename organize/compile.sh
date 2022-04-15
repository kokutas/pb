#!/bin/sh
###
 # @Author: kokutas
 # @Email: xxx
 # @Phone: xxx
 # @Date: 2022-04-16 01:18:00
 # @LastEditors: kokutas
 # @LastEditTime: 2022-04-16 02:12:52
 # @FilePath: /pb/organize/compile.sh
 # @Description: TODO
 # Copyright (c) 2022 by kokutas, All Rights Reserved. 
### 
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   organize.proto