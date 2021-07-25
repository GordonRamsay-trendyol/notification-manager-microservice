package repository

import (
	"github.com/couchbase/gocb/v2"
)

const (
	hostname   = "localhost"
	username   = "Administrator"
	password   = "password"
	bucketName = "users"
)

var cls *gocb.Cluster

func init() {
	cluster, err := gocb.Connect(
		hostname,
		gocb.ClusterOptions{
			Username: username,
			Password: password,
		},
	)

	if err != nil {
		panic(err)
	}

	cls = cluster

}
