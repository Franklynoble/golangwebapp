package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	/*
	open the file get the listener

	 */

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		 log.Panic(err)
	}
	defer li.Close()
	for {
		//eternally loop over the listener
		 //anything that comes into the listener, we accept it and
		conn , err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		//handle it with the go handle fu
		go handle(conn)
	}


}
//we handle the connection here, we pass in the connection
func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn) //pass connection to the scanner (what we get from the connection)
	for scanner.Scan() { // loop over
		ln := scanner.Text() // get the text from the token
		fmt.Println(ln) //prin it to one console
		fmt.Fprintf(conn, "I heard you say: %s\n",ln) // print it to one connection
	}
	defer conn.Close()


	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")

}