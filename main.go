package main

import (
    "fmt"
    "net"
    "strconv"
    "sync"
	"flag"
)

func scanPort(host string, port int, wg *sync.WaitGroup) {
    defer wg.Done()

    address := host + ":" + strconv.Itoa(port)
    conn, err := net.Dial("tcp", address)
    if err != nil {
        return
    }
    conn.Close()
    fmt.Printf("%s:%d\n", host, port)
}

func main() {
	// Adding commandline flags
	host := flag.String("host", "localhost", "The host you want to scan")
	start_port := flag.Int("start-port", 1, "The port you want your scan to start with")
	end_port := flag.Int("end-port", 65535, "The port you want your scan to end with")

	flag.Parse()

	// Start the scan
    var wg sync.WaitGroup

    for port := *start_port; port <= *end_port; port++ {
        wg.Add(1)
        go scanPort(*host, port, &wg)
    }

    wg.Wait()
}