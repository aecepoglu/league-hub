package reactFileServer

import (
	"net/http"
	"os"
	"strings"
	"path"
)

type FileServer struct {
	dir string
}

func New(dir string) FileServer {
	return FileServer {
		dir: dir,
	}
}

func (fs FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rpath := r.URL.Path
	if strings.Count(path.Base(rpath), ".") > 0 {
		fpath := path.Join(fs.dir, rpath)
		if f, err := os.Stat(fpath); err == nil && !f.IsDir() {
			http.ServeFile(w, r, fpath)
			return
		}
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, path.Join(fs.dir, "index.html"))
}

