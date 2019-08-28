WORK IN PROGRESS

# ETCD BATCH

Etcd batch can import a json/yaml/toml/envfile to etcd or dump etcd to a json/yaml/toml/envfile

Support file format JSON, YAML, TOML, envfile

etcd-batch [ import | import-file | dump | flat ] "prefix" json


## EXAMPLE

import json from stdin

etcd-batch import "prefix" < foo.json

import a json file

etcd-batch import-file "/PREFIX/" foo.json

dump a json file

etcd-batch dump /prefix/ > bar.json

flat json file

etcd-batch flat bar.json


## Other

This tool is created due to this issue: https://github.com/etcd-io/etcd/issues/8205

Similiar tools: [etcdtool](https://github.com/mickep76/etcdtool)   but not support for etcd v3.

TODO


kubectl style command line

etcd-batch apply/delete -n [namespace] -f foo.json

etcd-batch get -n [namespace] > bar.json
