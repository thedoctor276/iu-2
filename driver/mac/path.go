package mac

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/maxence-charriere/iu"
)

var packaged bool

func init() {
	var path string
	var err error

	if path, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		log.Panic(err)
	}

	for _, dir := range strings.Split(path, "/") {
		if strings.Contains(dir, ".app") {
			packaged = true
			break
		}
	}

	iu.Path = resPath
}

func resPath(elem ...string) string {
	var resources = "resources"

	if packaged {
		resources = resourcePath()
	}

	elem = append([]string{resources}, elem...)
	return filepath.Join(elem...)
}
