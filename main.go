package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func Connect(host, port string){
	d := net.Dialer{Timeout: 300*time.Millisecond}

	address := host + ":" + port	
	conn, err := d.Dial("tcp", address)

	if err != nil {
		return
	}

	defer conn.Close()
	fmt.Println("Port", port, "for Host", host, "is Open.")
}

func VanillaScan(host string){
	var port uint64 = 1
	for ; port < 65536; port++ {
		go Connect(host, strconv.FormatUint(port, 10))
	}
}

func SweepScan(hostListString string){
	hostList := strings.Split(hostListString, ",")
	var n int = len(hostList)
	for i:=0; i<n; i++ {
		fmt.Println(i, ":", hostList[i])
	}
}

func main(){
	startTime := time.Now()
	fmt.Println(`
	____   ___  ____ _____   ____   ____    _    _   _ _   _ _____ ____  
	|  _ \ / _ \|  _ \_   _| / ___| / ___|  / \  | \ | | \ | | ____|  _ \ 
	| |_) | | | | |_) || |   \___ \| |     / _ \ |  \| |  \| |  _| | |_) |
	|  __/| |_| |  _ < | |    ___) | |___ / ___ \| |\  | |\  | |___|  _ < 
	|_|    \___/|_| \_\|_|   |____/ \____/_/   \_\_| \_|_| \_|_____|_| \_\
																		  
	`)
	args := os.Args[1:]
	hostListString := args[0]
	SweepScan(hostListString)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Time Elapsed in Running this Program :", elapsedTime)
}