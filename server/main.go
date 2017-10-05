package main
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main(){
	//port that will be used for server side listening
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()
	//detects if there is an incoming connection from client side
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}
//handles the incoming request from client side 
func handle(conn net.Conn) {
	defer conn.Close()
	// read request
	request(conn)
	// write response
	//respond1(conn)
}

//prints the method request
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			uri:=strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			handleRoute(uri, conn);

		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func handleRoute(uri string, conn net.Conn){
	switch(uri){
	case "/Home":
		respond1(conn);
		break;
	default:
		respond2(conn);
		break;
	}
}

//respond to the client request which is still get
func respond1(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
func respond2(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Route Unavailable</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}