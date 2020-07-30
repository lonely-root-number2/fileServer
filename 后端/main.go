package main

import (
	//"bytes"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
)

func errHandle(i int,err error){
	if err != nil {
		fmt.Println(runtime.Caller(i))
		os.Exit(0)
	}

}

func getFile_dep(){

	//filepath.Walk("D:\\code\\Go", func(path string, info os.FileInfo, err error) error {
	//
	//	return nil
	//})
}

type fileInfo struct {
	FileName string `json:"fileName"`
	IsDir bool	`json:"IsDir"`
	LastModTime time.Time `json:"lastModTime"`
	FileSize int64 `json:"FileSize"`
}
func getFile(path string)[]fileInfo{
	// 必须特殊处理，，，
	if path == "c:"{
		path = "c:/"
	}
	ret := make([]fileInfo,0,25)
	tmp := make([]fileInfo,0,25)
	dirList, err := ioutil.ReadDir(path)
	errHandle(1,err)
	for _,v := range dirList{
		if v.IsDir(){
		ret = append(ret,fileInfo{
			FileName:    v.Name(),
			IsDir:       v.IsDir(),
			LastModTime: v.ModTime(),
			FileSize:    v.Size()/1024.0,
		})
		}else {
			tmp = append(tmp,fileInfo{
				FileName:    v.Name(),
				IsDir:       v.IsDir(),
				LastModTime: v.ModTime(),
				FileSize:    v.Size()/1024.0,
			})
		}

	}
	ret = append(ret,tmp...)
	//bf := bytes.NewBuffer([]byte{})
	//jdcode:=json.NewEncoder(bf)
	//jdcode.SetEscapeHTML(false)
	//jdcode.Encode(ret)
	//return bf.String()
	return ret
}
func Cors(origin,Method,Headers string)func(c *gin.Context){
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS"{
			c.Header("Access-Control-Allow-Origin",origin)
			c.Header("Access-Control-Request-Method","*")
			c.Header("Access-Control-Request-Headers","*")
		}else{
			c.Header("Access-Control-Allow-Origin",origin)
		}
		c.Next()
	}
}

func pathMerge(path []string,os string)string{
	if os == "windows"{
		path = path[1:]
		path[0] = path[0]+":"
		fmt.Printf("path merge is %s\n",path)
		return strings.Join(path,"/")
	}
	return strings.Join(path,"/")
}


func main() {

	os := runtime.GOOS
	fmt.Println(os)
	//	fmt.Printf("%v",getFile())
	r := gin.Default()
	r.Use(Cors("*","",""))
	// 获取操作系统
	r.GET("/getos", func(c *gin.Context) {
		if os == "windows"{
			c.JSON(200,gin.H{
				"os":os,
				"data": GetLogicalDrives(),
			})
		}

	})
	indexHandler := r.Group("/index")
	indexHandler.GET("/*act", func(c *gin.Context) {
		fmt.Println("-------------------")
		pathParam,ok := c.Params.Get("act")
		fmt.Printf("the pathParam is %s\n",pathParam)
		if !ok {
			// URI获取失败

		}
		if pathParam == "/" {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "ok",
				"data": GetLogicalDrives(),
				"os":   os,
			})
			return
		}
		// 舍去index[0]
		pathSlice := strings.Split(pathParam,`/`)
		pathEnd := pathMerge(pathSlice,os)
		fmt.Printf("the pathEnd is %v,lenth = %d\n",pathEnd,len(pathEnd))
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": getFile(pathEnd),
			"os":   os,
		})
	})
	r.Run(":8000")
}
