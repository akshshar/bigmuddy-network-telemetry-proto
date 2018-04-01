#!/bin/bash

SCRIPT_PATH=$(dirname `which $0`)

cd $SCRIPT_PATH
# Clean up the object files and library
make clean
