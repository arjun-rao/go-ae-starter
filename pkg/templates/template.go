/*
Package templates provides constants to use as html templates
*/
package templates

import (
	"path/filepath"
	"runtime"
	"strings"
)

// Path returns a canonical path to the templates directory
func Path(path ...string) string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	basepath = filepath.Join(basepath, "..", "..")
	templateDir := basepath + "/pkg/templates/"
	if len(path) == 0 {
		return templateDir
	}
	return templateDir + strings.Join(path, "/")
}
