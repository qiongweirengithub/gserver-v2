package application


import (
	"bytes"
	"fmt"
	"os/exec"
	"math/rand"
	"time"
	"io/ioutil"
	"bufio"
	"io"
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
	cmd := exec.Command(command, arg...)
	return runCmdAndPrint(cmd, workdir)
}


func ExecCICmd( workdir string, command string, arg ...string) (_ string, err error) {
	cmd := exec.Command(command, arg...)
	return runCmdAndPrint(cmd, workdir)
}



func ExecCmd(name string, arg ...string) (_ string, err error) {

	cmd := exec.Command(name, arg...)
	return runCmdAndPrint(cmd, "")
}


func runCmdAndPrint(cmd *exec.Cmd, wd string) (string, error) {

	fmt.Println(Blue, "WD: ", wd, Reset)
	fmt.Println(Blue, "CMD: ", cmd, Reset)

	if wd != "" {
		cmd.Dir = wd 
	}
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	
	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		
		if err2 != nil && io.EOF != err2 {
			return "error", err2
		}

		if io.EOF == err2 {
			return "success", nil
		}

		fmt.Print(Green , line, Reset)
	}

}



func runCmdAndPrintDepracated(cmd *exec.Cmd, workdir string) (string, error) {

	var out bytes.Buffer
	var stderr bytes.Buffer
	if workdir != "" {
		cmd.Dir = workdir 
	}
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	
	errs:=cmd.Run()
	if errs != nil {
		fmt.Println(Red, "---> work dir :", workdir, Reset)
		fmt.Println(Red, "---> cmd      :", cmd)
		fmt.Println(Red, "---> cmd run error", errs)
		return "", errs
	}
	a:= out.Bytes();
	fmt.Println(Red, "---> work dir :", workdir, Reset)
	fmt.Println(Green, "---> cmd :", cmd, Reset)
	fmt.Println(Green, "---> res :", string(a), Reset)

	return string("a"), nil

}
