#!/bin/bash
#
# Copyright (c) 2016 by cisco Systems, Inc. 
# All rights reserved.
#
set -x
mkdir -p src/gencpp
mkdir -p src/genobj
printf "Generating cplusplus bindings..."

cd ./proto_archive

for dir in `find . -type d -links 2`; do
  protoc -I ./  --grpc_out=../src/gencpp --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` $dir/*.proto
  protoc -I ./ --cpp_out=../src/gencpp $dir/*.proto
  mkdir -p ../src/genobj/$dir
done

for file in .; do
  protoc -I ./  --grpc_out=../src/gencpp --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` $file/*.proto
  protoc -I ./ --cpp_out=../src/gencpp $file/*.proto
done

echo "Done"
