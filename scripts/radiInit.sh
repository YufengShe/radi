#!/bin/bash
export ProjectDir=/home/ubuntu/project//radi

cd $ProjectDir
cd ./src

nohup go run main.go >radi.log 2>radi.error &
