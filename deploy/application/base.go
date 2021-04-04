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
	GreenBg      = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	WhiteBg      = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	YellowBg     = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	RedBg        = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	BlueBg       = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	MagentaBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	CyanBg       = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	Green        = string([]byte{27, 91, 51, 50, 109})
	White        = string([]byte{27, 91, 51, 55, 109})
	Yellow       = string([]byte{27, 91, 51, 51, 109})
	Red          = string([]byte{27, 91, 51, 49, 109})
	Blue         = string([]byte{27, 91, 51, 52, 109})
	Magenta      = string([]byte{27, 91, 51, 53, 109})
	Cyan         = string([]byte{27, 91, 51, 54, 109})
	Reset        = string([]byte{27, 91, 48, 109})
	disableColor = false

	black        = string([]byte{27, 91, 57, 48, 109})

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
		fmt.Println(Red, "--->  cd cmd run error", errs, Reset)
		fmt.Println(Red, "--->  work dir :", workdir, Reset)
		fmt.Println(Red, "--->  cmd      :", cmd, Reset)
		return "", errs
	}
	a:= out.Bytes();
	fmt.Println(Green, "--->  work dir :", workdir, Reset)
	fmt.Println(Green, "--->  cmd      :", cmd, Reset)
	fmt.Println(Green, "--->  res      :", string(a), Reset)


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
		fmt.Println(Red, "--->  ci cmd run error", errs, Reset)
		fmt.Println(Red, "--->  work dir :", workdir, Reset)
		fmt.Println(Red, "--->  cmd      :", cmd, Reset)
		return "", errs
	}
	a:= out.Bytes();
	fmt.Println(Green, "--->  work dir :", workdir, Reset)
	fmt.Println(Green, "--->  cmd      :", cmd, Reset)
	fmt.Println(Green, "--->  res      :", string(a), Reset)


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
		fmt.Println(Red, "--->  cmd      :", cmd)
		fmt.Println(Red, "--->  cmd run error", errs)
		return "", errs
	}
	a:= out.Bytes();
	fmt.Println(Green, "--->  cmd      :", cmd, Reset)
	fmt.Println(Green, "--->  res      :", string(a), Reset)
	return string(a), nil
}


func printLine() {
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("")
}
