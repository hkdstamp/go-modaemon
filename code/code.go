package code

import (
	"io/ioutil"
	"os"
	"path"
	"sort"
	"time"
)

type Code struct {
	URL  string
	Ref  string
	Path string
}

type Dirs []os.FileInfo

func (d Dirs) Len() int {
	return len(d)
}

func (d Dirs) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d Dirs) Less(i, j int) bool {
	return d[j].ModTime().Unix() < d[i].ModTime().Unix()
}

func (c *Code) CheckUpdate() (bool, error) {
	base := "/srv/code"

	dirs, err := ioutil.ReadDir(base)

	if err != nil {
		return false, err
	}

	if len(dirs) == 0 {
		return true, nil
	}

	sort.Sort(Dirs(dirs))

	if len(dirs) > 10 {
		for _, dir := range dirs[10:] {
			err := os.RemoveAll(path.Join(base, dir.Name()))
			if err != nil {
				return false, err
			}
		}
	}

	g := &Git{
		url:  c.URL,
		path: path.Join(base, dirs[0].Name()),
		ref:  c.Ref,
	}

	return g.checkUpdate()
}

func (c *Code) Get() (string, error) {
	t := time.Now().Format("20060102150405")
	g := &Git{
		url:  c.URL,
		path: path.Join("/srv", "code", t),
		ref:  c.Ref,
	}
	err := g.get()
	return g.path, err
}
