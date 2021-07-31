#!/bin/bash

# 启动一个 config server
mongod --bind_ip_all --configsvr --dbpath=/data/db

# 启动一个分片
mongod --bind_ip_all --shardsvr --dbpath=/data/db

tls_mode="${TLS_MODE}"
script_name=${0##*/}

if [[ "$AUTH" == "true" ]] ; then
  admin_user="$ADMIN_USER"
  admin_pass="$ADMIN_PASSWORD"
  admin_creds=(-u "$admin_user" -p "$admin_pass")
  auth_args=("--auth" "--keyFile=/data/configdb/key.txt")
fi

log() {
  local msg="$1"
  local timestamp
  timestamp=$(date -iso-8601=ns)
  echo "[$timestamp] [$script_name] $msg" 2>&1 | tee -a
}

retry_until() {
  local host="${1}"
  local command="${2}"
  local expected="${3}"
  local creds=("${admin_creds[@]}")

  if [[ "${host}" =~ ^localhost ]] ; then
      creds()
  fi
}