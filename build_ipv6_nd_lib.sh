#!/bin/bash

SCRIPT_PATH=$(dirname `which $0`)

cd $SCRIPT_PATH
#Generate the c++ binding from proto files

./gen-cpp-bindings-ipv6-nd.sh

# Compile the object files and build library
make
