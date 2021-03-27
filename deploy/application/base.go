package application


import (
	"bytes"
	"fmt"
	"os/exec"
	"math/rand"
	"time"
	"io/ioutil"
)


var (
	r *rand.Rand
	CD_Dir = "/home/qiongwei/myapp/k8s/app/" 
) 

func init() {
    r = rand.New(rand.NewSource(time.Now().Unix()))
}

func GetPid(path string)  (string){
    f, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("read fail", err)
    }
    return string(f)
}


func RandString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        b := r.Intn(26) + 65
        bytes[i] = byte(b)
    }
    return string(bytes)
}

func ExecCDCmd(name string, arg ...string) (_ string, err error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Dir = CD_Dir

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	errs:=cmd.Run()
	if errs != nil {
		fmt.Println(errs)
	}
	a:= out.Bytes();
	fmt.Println(string(a))

	return string(a), nil
}


func ExecCICmd( workdir string, name string, arg ...string) (_ string, err error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Dir = workdir 

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	errs:=cmd.Run()
	if errs != nil {
		fmt.Println(errs)
	}
	a:= out.Bytes();
	fmt.Println(string(a))

	return string(a), nil
}



func ExecCmd(name string, arg ...string) (_ string, err error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	errs:=cmd.Run()
	if errs != nil {
		fmt.Println(errs)
	}
	a:= out.Bytes();
	fmt.Println(string(a))

	return string(a), nil
}


