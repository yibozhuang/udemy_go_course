package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type genericClient struct {
	bytesWritten int
	bytesRead    int
}

func (c *genericClient) summary() {
	fmt.Printf("Total bytes read: %d\n", c.bytesRead)
	fmt.Printf("Total bytes written: %d\n", c.bytesWritten)
}

type Reader interface {
	Read(bytes int, client *genericClient)
}

type Writer interface {
	Write(bytes int, client *genericClient)
}

type ReaderWriter interface {
	Reader
	Writer
}

type dbReader struct{}

var _ Reader = dbReader{}

func (dbr dbReader) Read(bytes int, client *genericClient) {
	client.bytesRead += bytes
	log.Debugf("read %d bytes\n", bytes)
}

type dbWriter struct{}

var _ Writer = dbWriter{}

func (dbw dbWriter) Write(bytes int, client *genericClient) {
	client.bytesWritten += bytes
	log.Debugf("wrote %d bytes\n", bytes)
}

type dbReaderWriter struct {
	Reader
	Writer
}

type dbOperatorType string

const (
	DBREAD  dbOperatorType = "READ"
	DBWRITE dbOperatorType = "WRITE"
)

type operation struct {
	operatorType dbOperatorType
	size         int
	content      string
	client       *genericClient
}
