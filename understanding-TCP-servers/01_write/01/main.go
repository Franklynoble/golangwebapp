package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

  	li, err := net.Listen("tcp",":8080")

      if err != nil {
     	       log.Panic(err)
	}
	defer li.Close()

 for {
 	 conn, err := li.Accept()
 	  if err != nil {
 	  	   log.Println(err)
 	  	    }
 	  	     //write to a connection and return the connection
 	  	    io.WriteString(conn,"\n Hello from TCP server\n")
 	       fmt.Fprint(conn,"How is Your Day?")
 	 fmt.Fprintf(conn, "%v","Well, i gop!")
 }
}
