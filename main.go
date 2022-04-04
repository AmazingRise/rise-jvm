package main

import (
	"bytes"
	"github.com/AmazingRise/rise-jvm-core"
	"os"
	"syscall/js"
)

func main() {
	js.Global().Set("run", js.FuncOf(runWrapper))
	<-make(chan bool)
}

func Run(raw []byte) []byte {
	buf := bytes.NewBufferString("")
	core.Run(bytes.NewReader(raw), buf, os.Stdin, false)
	return buf.Bytes()
}

func runWrapper(this js.Value, args []js.Value) interface{} {
	arr := args[0]
	inBuf := make([]byte, arr.Get("byteLength").Int())
	//fmt.Println(inBuf)
	js.CopyBytesToGo(inBuf, arr)
	output := Run(inBuf)
	//fmt.Println(string(output))
	return js.ValueOf(string(output))
}
