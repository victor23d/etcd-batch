GOOS=darwin GOARCH=amd64 go build  -o etcd-batch-darwin etcd-batch.go
GOOS=linux GOARCH=amd64 go build  -o etcd-batch-linux etcd-batch.go
GOOS=windows GOARCH=amd64 go build etcd-batch.go
