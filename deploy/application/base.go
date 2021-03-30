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

func ExecCDCmd(workdir string, command string, arg ...string) (_ string, err error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	defer printLine()

	cmd := exec.Command(command, arg...)
	cmd.Dir = workdir

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	errs:=cmd.Run()
	if errs != nil {
		fmt.Println("=======error==========")
		fmt.Println("cd cmd run error", errs)
		fmt.Println("work dir :", workdir)
		fmt.Println("cmd      :", cmd)
		fmt.Println("=======error==========")
		return "", errs
	}
	a:= out.Bytes();
	fmt.Println("=======success==========")
	fmt.Println("work dir :", workdir)
	fmt.Println("cmd      :", cmd)
	fmt.Println("res      :", string(a))
	fmt.Println("=======success=========")
	

	return string(a), nil
}


func ExecCICmd( workdir string, command string, arg ...string) (_ string, err error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(command, arg...)
	cmd.Dir = workdir 

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	
	defer printLine()
	
	errs:=cmd.Run()


	if errs != nil {
		fmt.Println("=======error==========")
		fmt.Println("ci cmd run error", errs)
		fmt.Println("work dir :", workdir)
		fmt.Println("cmd      :", cmd)
		fmt.Println("========error=========")
		return "", errs
	}
	a:= out.Bytes();
	fmt.Println("=======success==========")
	fmt.Println("work dir :", workdir)
	fmt.Println("cmd      :", cmd)
	fmt.Println("res      :", string(a))
	fmt.Println("========success=========")


	return string(a), nil
}



func ExecCmd(name string, arg ...string) (_ string, err error) {

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	defer printLine()

	errs:=cmd.Run()
	if errs != nil {
		fmt.Println("=======error==========")
		fmt.Println("cmd      :", cmd)
		fmt.Println("cmd run error", errs)
		fmt.Println("========error=========")
		return "", errs
	}
	a:= out.Bytes();
	fmt.Println("=======success=========")
	fmt.Println("cmd      :", cmd)
	fmt.Println("res      :", string(a))
	fmt.Println("=======success==========")
	return string(a), nil
}


func printLine() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}
