package main

import (
	"os"
	"fmt"
	"flag"
	"net/http"
	"strconv"
	"image/png"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

var levels = map[string]qr.ErrorCorrectionLevel{"L":qr.L,"M":qr.M,"Q":qr.Q,"H":qr.H}
func main() {
	server  := flag.Bool("server", false, "服务器模式")
	port    := flag.Int("port", 9527, "服务器监听端口")

	content := flag.String("content", "helloworld", "二维码内容")
	output  := flag.String("output", "/tmp/qrcode.png", "二维码输出路径")
	level   := flag.String("level", "M", "容错级别L/M/Q/H")
	size    := flag.Int("size", 200, "图片像素尺寸")

	flag.Parse()

	if *server==true {
		serv(*port)
	} else {
		gen(*content, *output, *size, *level)
	}
}

func gen(content string, output string, size int, level string) {
	var qr_level = qr.M
	if _,ok := levels[level];ok {
		qr_level = levels[level]
	}

	file, _ := os.Create(output)
	defer file.Close()

	qrCode, _ := qr.Encode(content, qr_level, qr.Auto)
	qrCode, _  = barcode.Scale(qrCode, size, size)
	png.Encode(file, qrCode)
}

func serv(port int) {
	http.HandleFunc("/", render)

    err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}

   <-make(chan int)
}

func render(w http.ResponseWriter, r *http.Request) {
	content := r.URL.Query().Get("content")
	size    := r.URL.Query().Get("size")
	level   := r.URL.Query().Get("level")

	if content=="" {
		return
	}

	var height int
	var width int
	if size=="" {
		height = 200
		width  = 200
	} else {
		height,_ = strconv.Atoi(size)
		width,_  = strconv.Atoi(size)
	}

	var qr_level = qr.M
	if _,ok := levels[level];ok {
		qr_level = levels[level]
	}

	qrCode, _ := qr.Encode(content, qr_level, qr.Auto)
	qrCode, _  = barcode.Scale(qrCode, width, height)
	png.Encode(w, qrCode)
}
