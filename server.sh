#!/bin/bash
set -e

# server script
appName="theAppNameVarHolder"

##########################

cd $(dirname $0)

getpid() {
  echo $(ps -ef | grep -E "\s\.?\/${appName}" | awk '{print $2}')
}

status() {
  local pid=$(getpid)
  if [ ! -z $pid ]; then
    echo "$appName is runing pid: $pid"

    echo ""
    echo "ps status"
    ps -p "$pid" -o "user,pid,ppid,lstart,etime,rss,%mem,%cpu,command"

    echo ""
    echo "netstat:"
    net
  else
    echo "$appName is not runing"
  fi
}

start() {
  local pid=$(getpid)
  if [ -z $pid ]; then
    echo "starting $appName"
    if [ $# -gt 1 ]; then
      echo "start $appName with arguments: ${@:1}"
      ./$appName ${@:1}
    else
      ./$appName &>/dev/null &
      echo "$appName is runing pid: $!"
    fi
  else
    echo "$appName is already runing pid:$pid"
  fi
}

stop() {
  echo "stopping $appName"
  local pid=$(getpid)
  if [ ! -z $pid ]; then
    kill "$pid"
    sleep 1s
    pid=$(getpid)
    if [ ! -z $pid ]; then
      echo "$appName is still runing, try force stop!"
      kill -9 "$pid"
    fi
  fi

  echo "$appName stopped"
}

reload() {
  local pid=$(getpid)
  if [ ! -z $pid ]; then
    kill -USR2 "$pid"
    echo "$appName reloaded"
  else
    echo "$appName is not runing"
  fi
}

startOrReload() {
  local pid=$(getpid)
  if [ -z $pid ]; then
    start ${@:2}
  else
    echo "reloading $appName"
    reload
  fi
}

net() {
  local platform=$(uname)
  local pid=$(getpid)
  if [ -z $pid ]; then
    echo "$appName is not runing"
    return 1
  fi
  if [ $platform == "Linux" ]; then
    local data=$(netstat -tuan -p 2>/dev/null | awk "\$7~/^$pid|Address/")
    echo "$data" | grep -E 'LISTEN|Address'
    echo ""
    echo "connection count status:"
    echo "$data" | awk '($6 != "LISTEN" && $7 !~/Address/) {print $5 ":" $6}' |
      awk -F : '{print $1 " " $3}' | sort | uniq -c
  elif [ $platform == "Darwin" ]; then
    lsof -Pn -i4 -i6 | awk "\$2~/$pid/"
  else
    echo "unsupported platform"
  fi
}

version() {
  ./$appName -version
}

help() {
  ./$appName -h
}

case "$1" in
status)
  status
  ;;
start)
  start ${@:2}
  ;;
stop)
  stop
  ;;
restart)
  stop
  start
  ;;
reload)
  reload
  ;;
startOrReload)
  startOrReload ${@:2}
  ;;
net)
  net
  ;;
version)
  version
  ;;
help)
  help
  ;;
*)
  cat <<EOF

usage: $0 action [args]

  $0 script for simple controlling $appName

  actions supported:
    status          show $appName current status
    start           start $appName
                      see more arguments please execute: $0 help
    stop            stop $appName
    restart         stop then start
    reload          send reload signal (SIGUSR2) via kill command
    startOrReload   start $appName or just reload it if it is already runing
    net             netstat about $appName
    version         show version
    help            show $appName help

EOF
  ;;
esac

exit 0