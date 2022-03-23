package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func tracefile(str_content string)  {
	fd,_:=os.OpenFile("output.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_time:=time.Now().Format("2006-01-02 15:04:05");
	fd_content:=strings.Join([]string{"======",fd_time,"=====",str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}

func Index(context *gin.Context) {
	ip:=context.PostForm("ip")
	fmt.Println(ip)
	tracefile(ip)
}