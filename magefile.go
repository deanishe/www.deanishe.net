// +build mage

package main

import (
	"net/url"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Aliases for Mage commands

// Aliases for Mage commands

// Aliases for Mage commands

// Aliases for Mage commands
var Aliases = map[string]interface{}{
	"b": Build,
	"c": Clean,
	"p": Publish,
	"d": Data,
	"w": Weather,
	"e": Events,
	"r": Repos,
	"a": All,
}

const (
	// PublishRepo   = "https://github.com/deanishe/deanishe.github.io"
	PublishRepo   = "git@github.com:deanishe/deanishe.github.io.git"
	PublishBranch = "master"
	BinDir        = "./themes/alabastard/bin"
	BuildDir      = "./public"
	DataDir       = "./data"
)

var Hostname string

func hostname() string {
	if Hostname != "" {
		return Hostname
	}
	s := struct{
		BaseURL string `toml:"baseURL"`
	}{}

	if _, err := toml.DecodeFile("config.toml", &s); err != nil {
		panic(err)
	}

	URL, err := url.Parse(s.BaseURL)
	if err != nil {
		panic(err)
	}
	Hostname = URL.Hostname()

	return Hostname
}

// Assets compile .coffee files to JS
func Assets() error {
	fmt.Println("compiling assets ...")
	files, err := findFiles("./themes/alabastard", ".coffee")
	if err != nil {
		return err
	}

	args := append([]string{"-c"}, files...)
	return sh.Run("coffee", args...)
}

// All update, generate and publish website
func All() error {
	mg.Deps(Data)
	if err := Build(); err != nil {
		return err
	}

	return Publish()
}

// Build generate website in ./public
func Build() error {
	mg.Deps(Deps, Clean)
	if isInstalled("coffee") {
		mg.Deps(Assets)
	}
	fmt.Println("building ...")
	if err := hugo(nil); err != nil {
		return err
	}

	return ServiceWorker()
}

// Dev generate dev (ENV=dev) website in ./public
func Dev() error {
	mg.Deps(Deps, Clean)
	if isInstalled("coffee") {
		mg.Deps(Assets)
	}
	fmt.Println("building dev ...")
	env := map[string]string{"ENV": "dev"}
	return hugo(env)
}

func hugo(env map[string]string) error {
	var err error
	if env != nil {
		err = sh.RunWith(env, "hugo")
	} else {
		err = sh.Run("hugo")
	}
	if err != nil {
		return err
	}
	// needed to unfuck go.mod
	return Deps()
}

// Publish push website to hosting server
func Publish() error {
	fmt.Println("publishing site ...")
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.Chdir(BuildDir); err != nil {
		return err
	}

	msg := "rebuilt site " + time.Now().UTC().Format(time.RFC3339)
	fmt.Println("creating git repo ...")
	// make sure this isn't a git repo
	if err := sh.Rm(".git"); err != nil {
		return err
	}
	if err := sh.Run("git", "init", "."); err != nil {
		return err
	}
	if err := sh.Run("git", "add", "."); err != nil {
		return err
	}
	if err := sh.Run("git", "commit", "-m", msg); err != nil {
		return err
	}
	if err := sh.Run("git", "remote", "add", "origin", PublishRepo); err != nil {
		return err
	}

	fmt.Println("pushing site to GitHub pages ...")
	if err := sh.Run("git", "push", "-f", "origin", PublishBranch); err != nil {
		return err
	}

	return os.Chdir(wd)
}

// Deps update dependencies
func Deps() error {
	fmt.Println("installing deps ...")
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}
	return sh.Run("go", "mod", "verify")
}

// Clean delete files in ./public
func Clean() { mg.Deps(cleanMage, cleanBuild) }

func cleanMage() error { return sh.Run("mage", "-clean") }

func cleanBuild() error {
	fmt.Printf("cleaning %s ...\n", BuildDir)
	return sh.Rm(BuildDir)
}

// find any files under directory with given extension
func findFiles(dir, ext string) ([]string, error) {
	files := []string{}
	ext = strings.ToLower(ext)
	err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		x := filepath.Ext(path)
		if strings.ToLower(x) == ext {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// return true if a command is installed
func isInstalled(name string) bool {
	if err := sh.Run("hash", name); err == nil {
		return true
	}
	return false
}
