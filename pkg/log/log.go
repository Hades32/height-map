package log

import (
	"fmt"
)

var DebugEnabled = false

func Debug(a ...interface{}) {
	if DebugEnabled {
		fmt.Println(a...)
	}
}

func Debugf(s string, a ...interface{}) {
	if DebugEnabled {
		fmt.Printf(s, a...)
	}
}

func Debugfln(s string, a ...interface{}) {
	if DebugEnabled {
		fmt.Printf(s+"\n", a...)
	}
}
