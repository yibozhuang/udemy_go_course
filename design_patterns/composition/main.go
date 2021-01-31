package main

import (
	"fmt"
	"unicode/utf8"
)

type executor struct{}

func (l *executor) scan(r Reader, bytes int, client *genericClient) {
	read := 0
	for read < bytes {
		r.Read(1, client)
		read++
	}
}

func (l *executor) print(w Writer, data string, client *genericClient) {
	toWrite := utf8.RuneCountInString(data)
	for toWrite > 0 {
		w.Write(1, client)
		toWrite--
	}
}

func (l *executor) execute(rw ReaderWriter, operations []operation) {
	for _, op := range operations {
		if op.operatorType == DBREAD {
			l.scan(rw, op.size, op.client)
		}

		if op.operatorType == DBWRITE {
			l.print(rw, op.content, op.client)
		}
	}
}

func main() {
	client := &genericClient{}

	operations := []operation{
		{
			operatorType: DBREAD,
			size:         20,
			client:       client,
		},
		{
			operatorType: DBWRITE,
			content:      "design patterns in Go",
			client:       client,
		},
		{
			operatorType: DBREAD,
			size:         10,
			client:       client,
		},
		{
			operatorType: DBWRITE,
			content:      "how to write effective Go",
			client:       client,
		},
	}

	rw := dbReaderWriter{
		Reader: dbReader{},
		Writer: dbWriter{},
	}

	e := executor{}
	e.execute(rw, operations)

	fmt.Println("Summary:")
	client.summary()
}
