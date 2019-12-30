#!/usr/bin/env bash

app_name=demo
service_name=demo

path=/data/apps/goWebDemo

env=${1:-"dev"}
echo "rollback app on env:${env} ..."

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

echo "stop remote process now ..."
ssh -l root ${host} "supervisorctl stop ${service_name}"

echo "rollback remote app now ..."
ssh -l root ${host} "\\cp ${path}/${app_name}.bak ${path}/${app_name}"

echo "start remote service now ..."
ssh -l root ${host} "supervisorctl start ${service_name}"

echo "check remote service status"
ssh -l root ${host} "supervisorctl status ${service_name}"

if [[ $? -ne 0 ]]; then
    echo "in ${host} rollback ${service_name} failed"
    exit -1
fi

echo "Work Done Successful!!!"