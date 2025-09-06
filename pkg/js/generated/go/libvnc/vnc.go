package vnc

import (
	lib_vnc "github.com/g1i1u/nuclei-edited/v3/pkg/js/libs/vnc"

	"github.com/dop251/goja"
	"github.com/g1i1u/nuclei-edited/v3/pkg/js/gojs"
)

var (
	module = gojs.NewGojaModule("nuclei/vnc")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"IsVNC": lib_vnc.IsVNC,

			// Var and consts

			// Objects / Classes
			"IsVNCResponse": gojs.GetClassConstructor[lib_vnc.IsVNCResponse](&lib_vnc.IsVNCResponse{}),
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
