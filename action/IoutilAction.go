package action 

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"bufio"
)

func IoutilReadDir(writer http.ResponseWriter, request *http.Request) {
	rd, err := ioutil.ReadDir("/")
	if err != nil {
		fmt.Fprintf(writer, "read dir / failed")
		return	
	}

	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Fprintf(writer, "[" + fi.Name() + "]")
		}else{
			fmt.Fprintf(writer, fi.Name())
		}
	}
}

func IoutilReadFile(writer http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadFile("/tmp/test")
	if err != nil {
		fmt.Fprintf(writer, "ReadFile failed")
		return	
	}
	fmt.Fprintf(writer, string(content))
}

func IoutilWriteFile(writer http.ResponseWriter, request *http.Request) {
	content := "this is text"
	err := ioutil.WriteFile("/tmp/test", []byte(content), 777)
	if err != nil {
		fmt.Fprintf(writer, "WriteFile failed")
		return	
	}
	fmt.Fprintf(writer, "WriteFile sucess")
}

func IoutilTempDirFile(writer http.ResponseWriter, request *http.Request) {
	dir, err := ioutil.TempDir("", "Test")
	if err != nil {
		fmt.Fprintf(writer, "TempDir test failed")
		return	
	}
	defer os.Remove(dir)
	fmt.Fprintf(writer, dir)
	
	file, err := ioutil.TempFile(dir, "test")
	if err != nil {
		fmt.Fprintf(writer, "TempFile test failed")
		return	
	}
	defer os.Remove(file.Name())
	fmt.Fprintf(writer, file.Name())
}

//分块读取文件
func BufioReadBlock(writer http.ResponseWriter, request *http.Request) {
	file := "/tmp/test"
	fileHandle, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(writer, "file open failed")
		return	
	}
	defer fileHandle.Close()
	
	readHandle := bufio.NewReader(fileHandle)
	block := make([]byte, 3)
	for {
		_, err := readHandle.Read(block)   //到达文件最后的时候会返回一个EOF错误
		if err != nil {
			break
		}
		fmt.Fprintf(writer, string(block))
	}
}

func ReadFileBlock(writer http.ResponseWriter, request *http.Request) {
	file := "/tmp/test"
	fileHandle, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(writer, "file open failed")
		return	
	}
	defer fileHandle.Close()
		
	content := []byte{}								//存储文件内容
	buf := make([]byte, 100)						//临时的缓冲，按块读取，一次最多读取100字节
	
	for {
		n, err := fileHandle.Read(buf)				//读取文件到缓冲
		if n > 0 {
			content = append(content, buf[:n]...)	//将读取到的文件聚合起来
		}
		if err != nil {
			break      								//遇到流结束或者其他错误
		}
	}
	
	fmt.Fprintf(writer, string(content))
}

func BufioReadLine(writer http.ResponseWriter, request *http.Request) {
	file := "/tmp/test"
	fileHandle, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(writer, "file open failed")
		return	
	}
	defer fileHandle.Close()
	
	scanHandle := bufio.NewScanner(fileHandle)
	for scanHandle.Scan() {
		fmt.Fprintf(writer, scanHandle.Text())
	}
}



