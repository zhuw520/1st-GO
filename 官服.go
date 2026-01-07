package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 在这里直接修改配置
	port := 8086       // 改端口
	code := 200        // 改状态码
	html := `<!DOCTYPE html>
<html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.29.4</center>
</body>
</html>
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->`              // 改HTML
	
	// 处理所有请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, html)
	})
	
	addr := fmt.Sprintf(":%d", port)
	log.Printf("启动服务: 端口=%d, 状态码=%d", port, code)
	log.Printf("访问: http://localhost:%d", port)
	
	http.ListenAndServe(addr, nil)
}