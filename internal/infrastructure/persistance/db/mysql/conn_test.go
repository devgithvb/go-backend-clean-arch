package mysql

import (
	"log"
	"testing"
)

func TestConn(t *testing.T) {
	cfg := Config{
		Host:            "localhost", // usage docker: hostName(container name)
		Port:            "5001",      // usage docker: postName(5342)
		Username:        "admin",
		Password:        "123456",
		Database:        "simorgh_db",
		SSLMode:         "disable",
		MaxIdleConns:    2,
		MaxOpenConns:    15,
		ConnMaxLiftTime: 5,
	}

	mysql := New(cfg)
	if err := mysql.ConnectTo(); err != nil {
		log.Fatal(err)
	}

	log.Println(mysql.Conn())
}
