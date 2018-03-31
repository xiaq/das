package eval

import (
	"strconv"
	"syscall"

	"github.com/elves/elvish/eval/vals"
	"github.com/elves/elvish/eval/vars"
)

var builtinNs = Ns{
	"_":         vars.NewBlackhole(),
	"pid":       vars.NewRo(strconv.Itoa(syscall.Getpid())),
	"ok":        vars.NewRo(OK),
	"true":      vars.NewRo(true),
	"false":     vars.NewRo(false),
	"-json-nil": vars.NewRo(vals.JsonNil{}),
	"paths":     &EnvList{envName: "PATH"},
	"pwd":       PwdVariable{},
}

func addBuiltinFns(fns map[string]interface{}) {
	builtinNs.AddBuiltinFns("", fns)
}
