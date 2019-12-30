#!/usr/bin/env bash

project_path=$(pwd)
app_path=${project_path}"/../.."
app_name=demo
service_name=demo

path=/data/apps/goWebDemo

env=${1:-"dev"}
echo "deploy app to env:${env} ..."

if [[ "env_"${1} = "env_prd" ]]; then
    # prd server host and path
    host=prd_server
elif [[ "env_"${1} = "env_dev" ]]; then
    # test server host and path
    host=test_server
else
    echo "env ERROR, deploy script exit now..."
    exit
fi

echo "copy app to ${env} server now ..."
echo "app path:"${app_path}/${app_name}
scp ${app_path}/${app_name} root@${host}:${path}/${app_name}".tmp"

if [[ $? -ne 0 ]]; then
    echo "SCP ERROR, Exit Now..."
    exit 1
fi

echo "compare md5 now ..."
local_md5=`md5sum ${app_path}/${app_name} | awk '{print \$1;}'`   # for linux
# local_md5=`md5 ${app_path}/${app_name} | awk '{print \$4;}'`    # for mac
remote_md5=`ssh -l root ${host} "md5sum ${path}/${app_name}.tmp"`
remote_md5=`echo ${remote_md5} | awk '{print \$1;}'`
echo ${remote_md5} ${local_md5}
if [[ ${local_md5} != ${remote_md5} ]]; then
    echo ${local_md5} ${remote_md5}
    echo "Scp File Failed, Md5 Check ERROR"
    exit
fi

# app file copy right, start copy config file
if [[ "env_"${1} = "env_prd" ]]; then
    # prd env, copy config_prd.json hosts_prd.json to server
    scp ${project_path}/conf/config_prd.json root@${host}:${path}/conf/config.json
    scp ${project_path}/conf/hosts_prd.json root@${host}:${path}/conf/hosts.json
elif [[ "env_"${1} = "env_dev" ]]; then
    # prd env, copy config_prd.json hosts_prd.json to server
    scp ${project_path}/conf/config_test.json root@${host}:${path}/conf/config.json
    scp ${project_path}/conf/hosts_test.json root@${host}:${path}/conf/hosts.json
else
    echo "env ERROR, deploy script exit now..."
    exit
fi

echo "stop remote process now ..."
ssh -l root ${host} "supervisorctl stop ${service_name}"

echo "backup remote app now ..."
ssh -l root ${host} "mv ${path}/${app_name} ${path}/${app_name}.bak"

echo "replace remote app now ..."
ssh -l root ${host} "mv ${path}/${app_name}.tmp ${path}/${app_name}"

echo "start remote service now ..."
ssh -l root ${host} "supervisorctl start ${service_name}"

echo "check remote service status"
ssh -l root ${host} "supervisorctl status ${service_name}"

echo "Remove Local File Now ..."
rm ${app_path}/${app_name}

echo "Work Done Successful!!!"