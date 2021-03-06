package main

// This is a basic example of running an nsqd instance embedded. It creates
// and runs an nsqd with all of the default options, and then produces
// and consumes a single message. You are probably better off running a
// standalone instance, but embedding it can simplify deployment and is
// useful in testing.

// See https://github.com/nsqio/nsq/blob/master/nsqd/options.go and
// https://github.com/nsqio/nsq/blob/master/apps/nsqd/nsqd.go for
// more details on how to configure an embedded nsqd instance.
import (
	"bytes"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/nsqio/nsq/nsqd"
)

func main() {
	done := make(chan bool)

	// Run the embedded nsqd in a go routine
	go func() {
		// running an nsqd with all of the default options
		// (as if you ran it from the command line with no flags)
		// is literally these three lines of code. the nsqd
		// binary mainly wraps up the handling of command
		// line args and does something similar

		opts := nsqd.NewOptions()
		nsqd := nsqd.New(opts)
		nsqd.Main()

		// wait until we are told to continue and exit
		<-done
		nsqd.Exit()
	}()

	cfg := nsq.NewConfig()

	// // the message we'll send to ourselves
	msg := []byte("the message")

	// // Set up a Producer, pointing at the default host:port
	// p, err := nsq.NewProducer("localhost:4150", cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Publish a single message to the 'embedded' topic
	// err = p.Publish("embedded", msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Now set up a consumer
	c, err := nsq.NewConsumer("embedded", "local", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// and a single handler that just checks that the message we
	// received matches the message we sent
	c.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		if bytes.Compare(m.Body, msg) != 0 {
			log.Fatal("message didn't match:", string(m.Body))
		} else {
			log.Println("message matched:", string(m.Body))
		}
		return nil
	}))

	// // Connect the consumer to the embedded nsqd instance
	c.ConnectToNSQD("192.168.186:4150")

	// Sleep a little to give everything time to start up and let
	// our producer and consumer run
	// time.Sleep(250 * time.Millisecond)
	time.Sleep(100 * time.Second)

	// tell the nsqd instance to exit
	done <- true

}
