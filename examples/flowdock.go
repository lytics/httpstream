/*

https://stream.flowdock.com/flows?filter=example/main,example/flow

*/

package main

import (
	"flag"
	"log"
	"net/url"
	"os"

	. "github.com/araddon/httpstream"
)

var (
	token    *string = flag.String("token", "password", "Password")
	flow     *string = flag.String("flow", "organization/flow", "Flowdock url path:  organization/flow")
	logLevel *string = flag.String("logging", "debug", "Which log level: [debug,info,warn,error,fatal]")
)

func main() {

	flag.Parse()

	// make a go channel for msgs
	stream := make(chan []byte, 200)
	flowUrl, _ := url.Parse("https://" + *token + "@stream.flowdock.com/flows" + *flow)

	// set the logger and log level
	SetLogger(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile), *logLevel)

	// the stream listener effectively operates in one "thread"
	client := NewBasicAuthClient("", "", func(line []byte) {
		stream <- line
	})
	_ = client.Connect(flowUrl, "")

	for {
		evt := <-stream
		Debug(string(evt))
	}
}