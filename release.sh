#!/bin/bash
MPATH=/data/push_server/mpush_server
GOPATH=$GOPATH:$MPATH
echo $MPATH
go build -o $MPATH/bin/mpush src/mpush.go
cp -fr src/conf $MPATH/bin/
echo "done build"
killall -9 mpush
#fuser -k ./bin/push_srv
echo "pid killed"
nohup ./bin/mpush -push ":7243" &
echo "done restart"