WORK IN PROGRESS

# ETCD BATCH

Etcd batch can import a json to etcd or dump etcd to a json

Also support file format JSON, 

TODO YAML, TOML, envfile

## EXAMPLE

import from a json file

etcd-batch apply -f foo.json --prefix ""

dump to a json file

etcd-batch dump -o json --prefix "" > bar.json

flat json file

etcd-batch flat foo.json

## Install

Just download from Release page.

## Other

This tool is created due to this issue: https://github.com/etcd-io/etcd/issues/8205

Similiar tools: [etcdtool](https://github.com/mickep76/etcdtool)   but not support for etcd v3.

