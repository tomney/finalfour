package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// GetSQLDB sets up access to the bolg SQL db
func GetSQLDB() sql.DB {
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("/etc/certs/server-ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	clientCert := make([]tls.Certificate, 0, 1)
	certs, err := tls.LoadX509KeyPair("/etc/certs/client-cert.pem", "/etc/certs/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	clientCert = append(clientCert, certs)
	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs:            rootCertPool,
		Certificates:       clientCert,
		InsecureSkipVerify: true,
	})

	var username string
	if username = os.Getenv("DB_USERNAME"); username == "" {
		log.Fatal("Could not get the username for the Cloud SQL database")
	}

	var password string
	if password = os.Getenv("DB_PASSWORD"); password == "" {
		log.Fatal("Could not get the password for the Cloud SQL database")
	}

	// SQL Client
	databasePublicIP := "34.66.232.103"
	datasource := fmt.Sprintf("%s:%s@tcp(%s:3306)/finalfour?tls=custom", username, password, databasePublicIP)
	db, err := sql.Open("mysql", datasource)
	return *db
}
