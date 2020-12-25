#!/bin/bash
workdir=$(cd $(dirname $0); pwd)
dir=`echo $workdir|grep -o '[^/]*$'`
go build
ps -ef|grep $dir |grep -v 'grep'|awk '{print $2}'|xargs kill -9
nohup ./$dir > $dir.log 2>&1 &