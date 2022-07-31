package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
	session *gocql.Session
)

func init() {
	cluster = gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.EachQuorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
