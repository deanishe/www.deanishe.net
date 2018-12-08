// +build mage

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

var (
	workerFilename = "sw.js"
	workerTemplate = "serviceworker.template.js"
	workerGlobs    = []string{
		// Assets
		// "img/*.png",
		// "img/*.jpg",
		// "img/*.svg",
		// "img/people/*.jpg",
		"js/*.min.*.js",
		"style/*.*.css",
		// Pages
		// "/",
		// "botschafter/",
		// "datenschutz/",
		// "gaeste/",
		// "helfen/",
		// "kontakt/",
		// "kulturpartner/",
		// "presse/",
		// "rechner/",
		// "schirmherr/",
		// "sozialpartner/",
		// "verein/",
	}
	workerExcludeGlobs = []string{
		// "favicons/**",
		// "bilder/**",
		// "download/**",
		// "img/icons/**",
		// "sw.js",
	}
)

// ServiceWorker generates a service worker script.
func ServiceWorker() error {

	fmt.Println("generating service worker ...")

	files, err := findWorkerFiles()
	if err != nil {
		return err
	}

	// fmt.Printf("%d URL(s) to pre-cache\n", len(files))

	data, err := makeWorker(files)
	if err != nil {
		return err
	}

	p := filepath.Join(BuildDir, workerFilename)
	if err := ioutil.WriteFile(p, data, 0644); err != nil {
		return err
	}

	fmt.Printf("wrote service worker to %s\n", p)

	return nil
}

func hashFiles(files []workerFile) ([]byte, error) {
	h := sha256.New()

	for _, wf := range files {
		f, err := os.Open(wf.Path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		if _, err := io.Copy(h, f); err != nil {
			return nil, err
		}
		f.Close()
	}

	return h.Sum(nil), nil
}

type workerFile struct {
	Path string
	URL  string
}

func (f workerFile) String() string {
	return fmt.Sprintf("workerFile{Path: %q, URL: %q}", f.Path, f.URL)
}

func fileMatches(path string, globs []string) bool {
	for _, pat := range globs {
		if ok, err := doublestar.PathMatch(pat, path); ok {
			return true
		} else if err != nil {
			panic(err)
		}
	}

	return false
}

// func isWorkerFile(path string) bool {
// 	for _, pat := range workerGlobs {
// 		if ok, err := filepath.Match(pat, path); ok {
// 			return true
// 		} else if err != nil {
// 			panic(err)
// 		}
// 	}

// 	return false
// }

// return list of files for service worker to pre-cache.
func findWorkerFiles() ([]workerFile, error) {
	files := []workerFile{}

	fmt.Printf("BuildDir=%s\n", BuildDir)

	err := filepath.Walk(BuildDir, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		p, err := filepath.Rel(BuildDir, path)
		if err != nil {
			return err
		}

		if fi.Name() == "index.html" {
			p = filepath.Dir(p) + "/"
		}

		// fmt.Printf("path=%q, p=%q\n", path, p)

		if fileMatches(p, workerExcludeGlobs) || !fileMatches(p, workerGlobs) {
			// fmt.Printf("- %s\n", p)
			return nil
		}

		if p != "/" {
			p = "/" + p
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		// h := sha256.New()
		// if _, err := io.Copy(h, f); err != nil {
		// 	return err
		// }

		wf := workerFile{Path: path, URL: p}

		fmt.Printf("+ %s\n", wf.URL)
		files = append(files, wf)
		// fmt.Printf("path=%s, p=%s, fi=%+v\n", path, p, fi)

		return nil
	})

	return files, err
}

// generate the JS for the service worker
func makeWorker(files []workerFile) ([]byte, error) {

	// calculate cumulative hash for all files
	h, err := hashFiles(files)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("cumulative hash: %x\n", h)

	urls := make([]string, len(files))
	for i, f := range files {
		urls[i] = f.URL
	}

	fmt.Printf("%d URL(s) to pre-cache\n", len(urls))

	data, err := ioutil.ReadFile(workerTemplate)
	if err != nil {
		return nil, err
	}

	s := fmt.Sprintf("<script>\n%s\n</script>", string(data))
	tpl := template.Must(template.New("_").Parse(s))
	buf := &bytes.Buffer{}
	err = tpl.Execute(buf, map[string]interface{}{
		"Hash": fmt.Sprintf("%x", h),
		"URLs": urls,
	})
	if err != nil {
		return nil, err
	}

	s = buf.String()
	s = s[9 : len(s)-10] // remove <script> tags

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	s, err = m.String("text/javascript", s)
	if err != nil {
		return nil, err
	}

	return []byte(s), nil
}
