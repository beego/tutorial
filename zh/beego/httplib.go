package main

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/httplib"
)

func main() {
	req := httplib.Get("http://qiniu.com")
	body, _ := req.String()
	lines := strings.Split(body, "\n")
	fmt.Println(strings.Join(lines[2:8], "\n"))
}
