package mac

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu-log"
)

var packaged bool

func init() {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		iulog.Panic(err)
	}

	for _, dir := range strings.Split(path, "/") {
		if strings.Contains(dir, ".app") {
			packaged = true
			break
		}
	}

	iu.SetResourcesPath(resourcesPath)
}

func resourcesPath(elem ...string) string {
	resources := "resources"

	if packaged {
		iulog.Warn("I'm packaged")
		resources = appResourcesPath()
	}

	elem = append([]string{resources}, elem...)
	return filepath.Join(elem...)
}
