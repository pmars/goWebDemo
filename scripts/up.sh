#!/usr/bin/env bash

app_name=demo
service_name=demo
path=/data/apps/goWebDemo
local_path=~/go/src/goWebDemo

env=${1:-"test"}

# ~/.ssh/config 文件中录入以下内容
# Host test_server
#     HostName 172.16.18.29
#     Port 22022
#     User root

if [[ "env_"${1} = "env_prd" ]]; then
    echo "build app for pay env ..."
    # prd server host and path
    host=prd_server
else
    echo "build app for test env ..."
    # test server host and path
    host=test_server
fi

cd ${local_path}/app

echo "start build app now ..."
GOOS=linux ARCH=amd64 go build -o ${app_name}

if [[ $? -ne 0 ]]; then
    echo "Build ERROR, Exit Now..."
    exit
fi

echo "start copy app to server now ..."
scp ${app_name} root@${host}:${path}/${app_name}.tmp

echo "compare md5 now ..."
local_md5=`md5 ${app_name} | awk '{print $4;}'`
remote_md5=`ssh ${host} "md5sum ${path}/${app_name}.tmp"`
remote_md5=`echo ${remote_md5} | awk '{print \$1;}'`
echo ${remote_md5}
if [[ ${local_md5} != ${remote_md5} ]]; then
    echo ${local_md5} ${remote_md5}
    echo "Scp File Failed, Md5 Check ERROR"
    exit
fi

if [[ "env_"${1} = "env_prd" ]]; then
    scp ${local_path}/conf/config_prd.json root@${host}:${path}/conf/config.json
    scp ${local_path}/conf/hosts_prd.json root@${host}:${path}/conf/hosts.json
else
    scp ${local_path}/conf/config_test.json root@${host}:${path}/conf/config.json
    scp ${local_path}/conf/hosts_test.json root@${host}:${path}/conf/hosts.json
fi

echo "backup remote app now ..."
ssh ${host} "cp ${path}/${app_name} ${path}/${app_name}.bak"

echo "rename remote app now ..."
ssh ${host} "mv ${path}/${app_name}.tmp ${path}/${app_name}"

echo "restart remote process now ..."
echo "restart service now ..."
ssh ${host} "supervisorctl restart ${service_name}"

echo "status remote service now ..."
ssh ${host} "supervisorctl status"

echo "Remove Local File Now ..."
rm ${app_name}

echo "Work Done Successful!!!"

cd -

