## 说明

创建一个http主要使用到了net包下面的http子包。

## 服务端

### 启动一个http服务
```golang
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 定义返回数据内容
type HeaderJson struct {
	StatusOk      int
	StatusMessage string
}

func Server() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		headerJson := &HeaderJson{StatusOk: http.StatusOK, StatusMessage: http.StatusText(http.StatusOK)}

		writer.Header().Set("Content-type", "Application/json")
		writer.WriteHeader(http.StatusOK) // 写在设置头信息后面，否则设置的header无法生效

		json.NewEncoder(writer).Encode(headerJson)// 对返回数据使用json进行序列化
	})

	err := http.ListenAndServe(":8100", nil)// 绑定IP和端口号， 格式（IP:端口号），默认的端口号是80
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

## 客户端

### 发起http-get请求

```golang
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1/")
	/** 对应服务的代码
	<?php
	echo time();
	 */
	if err != nil {
		fmt.Println("请求失败", err)
	}

	respContent := resp.Body
	val, ioErr := ioutil.ReadAll(respContent)
	if ioErr != nil {
		fmt.Println("解析失败", val)
	}
	fmt.Println(string(val))

	defer resp.Body.Close()// 结束请求之后一定要关闭
	// output:1636701618
}
```

### 发起http-get请求(带参数)

```golang
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	apiUrl := "http://127.0.0.1/"

	data := url.Values{}
	data.Set("name", "张三")
	data.Set("age", "12")

	u, uErr := url.ParseRequestURI(apiUrl)
	if uErr != nil {
		fmt.Println("url解析错误")
	}
	u.RawQuery = data.Encode()// 序列化请求的url参数

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("请求失败", err)
	}

	respContent := resp.Body
	val, ioErr := ioutil.ReadAll(respContent)
	if ioErr != nil {
		fmt.Println("解析失败", val)
	}
	
	defer resp.Body.Close()// 结束请求之后一定要关闭
}
```

### 发起http-post请求

```golang
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	apiUrl := "http://127.0.0.1/"

	contentType := "application/json"
	data := `{"name":"张三","age":"12"}`
	resp, err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("请求失败", err)
	}

	respContent := resp.Body
	val, ioErr := ioutil.ReadAll(respContent)
	if ioErr != nil {
		fmt.Println("解析失败", val)
	}
	fmt.Println(string(val), resp.StatusCode, resp.Status)

	defer resp.Body.Close()// 结束请求之后一定要关闭
}
```