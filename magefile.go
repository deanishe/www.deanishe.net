// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

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

var (
	// PublishRepo   = "https://github.com/deanishe/deanishe.github.io"
	PublishRepo   = "git@github.com:deanishe/deanishe.github.io.git"
	PublishBranch = "master"
	BinDir        = "./themes/alabastard/bin"
	BuildDir      = "./public"
	DataDir       = "./data"

	// Subdirectories of DataDir
	ForecastFile = "darksky/forecast.json"
	ReposFile    = "github/repos.json"
	EventsFile   = "github/events.json"
	PostsFile    = "pinboard/posts.json"
)

func init() {
	ForecastFile = filepath.Join(DataDir, ForecastFile)
	ReposFile = filepath.Join(DataDir, ReposFile)
	EventsFile = filepath.Join(DataDir, EventsFile)
	PostsFile = filepath.Join(DataDir, PostsFile)
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
	if err := sh.Run("hugo"); err != nil {
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
	return sh.RunWith(env, "hugo")
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

// Manage your deps, or running package managers.
func Deps() error {
	fmt.Println("installing deps ...")
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}
	return sh.Run("go", "mod", "verify")
}

// Clean delete files in ./public
func Clean() error {
	fmt.Printf("cleaning %s ...\n", BuildDir)
	err := os.RemoveAll(BuildDir)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	fmt.Println("cleaning mage cache ...")
	return sh.Run("mage", "-clean")
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
