// 编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

// 接收客户端 request，并将 request 中带的 header 写入 response header
// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 当访问 localhost/healthz 时，应返回 200

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func headers(response http.ResponseWriter, request *http.Request) {

	//question1
	headers := request.Header
	for header := range headers {
		values := headers[header]
		//fmt.Println(values)
		for index, _ := range values {
			values[index] = strings.TrimSpace(values[index])
		}
		response.Header().Set(header, strings.Join(values, ","))
	}

	//question 2
	os.Setenv("VERSION", "go1.18.3")
	version := os.Getenv("VERSION")
	response.Header().Set("VERSION-GO", version)

	//question 3
	ipPort := request.RemoteAddr
	println("Client->ip:port=" + ipPort)
	ip := strings.Split(ipPort, ":")
	println("Client->ip=" + ip[0])
	println("Client->response code=" + strconv.Itoa(http.StatusOK))

	io.WriteString(response, "================Details of the http response header:============\n")
	for k, v := range response.Header() {
		io.WriteString(response, fmt.Sprintf("%s=%s\n", k, v))
	}

}

//question 4
func healthz(w http.ResponseWriter, r *http.Request) {
	HealthzCode := "200"
	w.Write([]byte(HealthzCode))
}
