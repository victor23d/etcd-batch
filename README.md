# ETCD BATCH

Etcd batch can import a json/yaml/toml/envfile to etcd or dump etcd to a json/yaml/toml/envfile

Support file format JSON, YAML, TOML, envfile

etcd_batch [ import | import-file | dump | flat ] "prefix" json


## EXAMPLE

import json from stdin

etcd_batch import "prefix" < foo.json

import a json file

etcd_batch import-file "/PREFIX/" foo.json

dump a json file

etcd_batch dump /prefix/ > bar.json

flat json file

etcd_batch flat bar.json


## Other

This tool is created due to this issue: https://github.com/etcd-io/etcd/issues/8205

Similiar tools: [etcdtool](https://github.com/mickep76/etcdtool)   but not support for etcd v3.

TODO


kubectl style command line

etcd_batch apply/delete -n [namespace] -f foo.json

etcd_batch get -n [namespace] > bar.json
