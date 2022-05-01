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

func runWrapper(this js.Value, args []js.Value) interface{} {
	//fmt.Println(args)
	arr := args[0]
	// rawMainArg := args[1].String()
	// mainArgs := strings.Split(rawMainArg, " ")
	inBuf := make([]byte, arr.Get("byteLength").Int())
	js.CopyBytesToGo(inBuf, arr)
	buf := bytes.NewBufferString("")
	core.Run(bytes.NewReader(inBuf), buf, os.Stdin, true) //, mainArgs...)
	// If we return buf.String() here, we will get a TinyGo error in the console.
	return js.ValueOf(string(buf.Bytes()))
}
