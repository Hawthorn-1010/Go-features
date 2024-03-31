package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	//定义请求
	req, err := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 定义上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*80)
	defer cancel()
	req = req.WithContext(ctx)

	// 执行请求
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// 输出日志
	out, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}
