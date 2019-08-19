# ETCD BATCH


Etcd batch can import a json/yaml/toml/envfile to etcd or dump etcd to a json/yaml/toml/envfile

Support file format JSON, YAML, TOML, envfile

import a json file

etcd /prefix < foo.json

dump a json file

etcd /prefix > bar.json


This tool is created due to this issue: https://github.com/etcd-io/etcd/issues/8205

Similiar tools: [etcdtool](https://github.com/mickep76/etcdtool)   but not support for etcd v3.

