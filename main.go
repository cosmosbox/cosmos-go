package main

import (
	"github.com/valyala/gorpc"
	"log"
)

func main() {
	go func() {
		s := &gorpc.Server{
			// Accept clietsnts on this TCP address.
			Addr: ":12333",

			// Echo handler - just return back the message we received from the client
			Handler: func(clientAddr string, request interface{}) interface{} {
				log.Printf("Obtained request %+v from the client %s\n", request, clientAddr)
				return request
			},
		}
		if err := s.Serve(); err != nil {
			log.Fatalf("Cannot start rpc server: %s", err)
		}
	}()
	c := &gorpc.Client{
		// TCP address of the server.
		Addr: "127.0.0.1:12333",
	}
	log.Print("here")
	c.Start()
	log.Print("here2")
	resp, err := c.Call("foobar")
	if err != nil {
		log.Fatalf("Error when sending request to server: %s", err)
	}
	if resp.(string) != "foobar" {
		log.Fatalf("Unexpected response from the server: %+v", resp)
	}

	log.Printf("The result: %+v", resp)
}
