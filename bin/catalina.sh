#!/bin/bash
set -e
WORK_DIR=$(cd "$(dirname $0)/.." && pwd)
cd "${WORK_DIR}"
APP_NAME="${WORK_DIR##*/}"
# 添加全路径是为了进程检测
PROCESS_CHECK="${WORK_DIR}/bin/${APP_NAME}"
CONF_DIR=./conf
LOG_PATH=./logs
SHUTDOWN_SECONDS=30
if [ -f ${CONF_DIR}/setenv.sh ]  ; then
  source ${CONF_DIR}/setenv.sh
fi
echo "WORK_DIR=${WORK_DIR},APP_NAME=${APP_NAME}"

# 启动
start() {
  echoRed "Application Starting...."
  RUN_CMD=${WORK_DIR}/bin/$APP_NAME
  if [ ! -d logs ]; then
    mkdir logs
  fi
  if [ "${IN_CONTAINER}" = true ]; then
    $RUN_CMD >logs/catalina.out 2>&1
  else
    nohup $RUN_CMD >logs/catalina.out 2>&1 &
  fi
  echo "waiting and checking 10s"
  for i in {1..10}; do
    echo -n "$i..."
    sleep 1
  done
  echo
  health
  if [ "${IN_CONTAINER}" != true ]; then
    startSupervisor
  fi
  echoGreen "Application Started"
}
# 停止
stop() {
  if [ "${IN_CONTAINER}" != true ]; then
        uninstallSupervisor
    fi
  running=$(pidRunning)
  if [ $running -eq 0 ]; then
    echoYellow "${APP_NAME} not running"
  else
    pid=$(getPid)
    echoRed "${APP_NAME} pid:${pid} stopping..."
    kill $pid
    while [ $SHUTDOWN_SECONDS -gt 0 ]; do
      running=$(pidRunning)
      if [ $running -eq 1 ]; then
        sleep 1
        SHUTDOWN_SECONDS=$(($SHUTDOWN_SECONDS - 1))
      else
        echoGreen "${APP_NAME} stoped"
        return
      fi
    done
    running=$(pidRunning)
    if [ $running -eq 1 ]; then
      kill -9 $pid
      echoRed "${APP_NAME} force stoped"
    else
      echoGreen "${APP_NAME} stoped"
    fi
  fi
}

restart() {
  stop
  sleep 2
  start
}
health() {
  running=$(pidRunning)
  if [ $running -eq 0 ]; then
    echoYellow "${APP_NAME} not running"
    exit 1
  else
    echoGreen "${APP_NAME} is running，pid is $(getPid)"
  fi
}

getPid() {
  ps_pid=$(ps -ef | grep "${PROCESS_CHECK}" | grep -v grep | awk '{print $2}' | head -n 1)
  echo ${ps_pid:-0}
}

# ---- 根据pid查询是否进程在运行
pidRunning() {
  pid=$(getPid)
  ps -p ${pid} >/dev/null 2>&1
  # $?上个命令退出状态
  if [ $? -eq 0 ]; then
    echo 1
  else
    echo 0
  fi
}
startSupervisor() {
  cmd="${WORK_DIR}/bin/catalina.sh supervise"
  exists=$(crontab -l | grep "${cmd}" | grep -v grep | wc -l)
  if [ ${exists} -lt 1 ]; then
    crontab <<EOF
$(crontab -l)
* * * * * ${cmd} > /dev/null 2>&1 &
EOF
  fi
}
# 关闭自动拉起
uninstallSupervisor() {
  cmd="${WORK_DIR}/bin/catalina.sh supervise"
    crontab <<EOF
$(crontab -l | grep -v "${cmd}" | grep -v grep)
EOF
echoRed "uninstalled supervisor"
}

supervise() {
  log=${LOG_PATH}/supervior.log
  running=$(pidRunning)
  if [ ${running} -eq 0 ]; then
    time=$(date +"%Y-%m-%d %H:%M:%S")
    echo "=============== ${time} ===============" >>${log}
    echo "${APP_NAME} not running， now start a new process." >>${log}
    start
  fi
}

# ---- 带颜色的输出
echoWithColor() {
  c=$1
  shift
  echo $'\e[0;'"${c}"'m'"$@"$'\e[0m'
}

echoRed() {
  echoWithColor 31 "$@"
}

echoGreen() {
  echoWithColor 32 "$@"
}

echoYellow() {
  echoWithColor 33 "$@"
}

usage() {
  echo "Usage: catalina.sh [start|stop|restart|health]"
}

case $1 in
start)
  start
  ;;
stop)
  stop
  ;;
restart)
  restart
  ;;
health)
  health
  ;;
supervise)
  supervise
  ;;
*)
  usage
  exit 1
  ;;
esac
