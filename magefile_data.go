// +build mage

package main

import (
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	// CLI scripts in themes/alabastard/bin
	github   func(args ...string) error
	forecast func(args ...string) error
	pinboard func(args ...string) error
)

func init() {
	github = sh.RunCmd(filepath.Join(BinDir, "github"))
	forecast = sh.RunCmd(filepath.Join(BinDir, "forecast"))
	pinboard = sh.RunCmd(filepath.Join(BinDir, "pinboard-public"))
}

// Data fetch all remote data from GitHub, Pinboard etc.
func Data() {
	mg.Deps(Repos, Events, Weather, Posts)
}

// Repos fetch list of user's repos from GitHub
func Repos() error {
	if err := os.MkdirAll(filepath.Dir(ReposFile), 0755); err != nil {
		return err
	}
	return github("repos", ReposFile)
}

// Events fetch list of user's actions from GitHub
func Events() error {
	if err := os.MkdirAll(filepath.Dir(EventsFile), 0755); err != nil {
		return err
	}
	return github("events", EventsFile)
}

// Weather fetch weather forecast data from DarkSky
func Weather() error {
	if err := os.MkdirAll(filepath.Dir(ForecastFile), 0755); err != nil {
		return err
	}
	return forecast(ForecastFile)
}

// Posts fetch recent posts from Pinboard
func Posts() error {
	if err := os.MkdirAll(filepath.Dir(PostsFile), 0755); err != nil {
		return err
	}
	return pinboard(PostsFile)
}
