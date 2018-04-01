#!/bin/bash
#
# Copyright (c) 2016 by cisco Systems, Inc. 
# All rights reserved.
#
set -x

printf "Generating Python bindings..."
mkdir -p src/genpy

cd ./proto_archive

for dir in `find . -type d -links 2`; do
  protoc -I ./ --python_out=../src/genpy/ --grpc_out=src/genpy/ --plugin=protoc-gen-grpc=`which grpc_python_plugin` $dir/*.proto
  mkdir -p ../src/genpy/$dir
  touch ../src/genpy/$dir/__init__.py
done

for file in .; do
  protoc -I ./ --python_out=../src/genpy/ --grpc_out=src/genpy/ --plugin=protoc-gen-grpc=`which grpc_python_plugin` *.proto
done

echo "Done"
