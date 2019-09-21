# ETCD BATCH

Etcd batch can import a json to etcd or dump etcd to a json

Also support file format JSON, 

TODO YAML, TOML, envfile, WORK IN PROGRESS

## EXAMPLE

Import kv from a json file

etcd-batch apply -f foo.json --prefix "/"

Dump kv to a json file

etcd-batch dump -o json --prefix "" > bar.json

Flat json file, will output flatted keys.

etcd-batch flat foo.json

## Option

Share most options from etcdctl e.g. ETCDCTL_ENDPOINTS

Also share environment variables.


## Install

Just download from Release page.

## Other

This tool is created due to this issue: https://github.com/etcd-io/etcd/issues/8205

Similiar tools: [etcdtool](https://github.com/mickep76/etcdtool)   but not support for etcd v3.


### TODO

CI/CD auto build and release by tag

go tool chain report

