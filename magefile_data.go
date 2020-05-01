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
	github = sh.RunCmd(filepath.Join(BinDir, "github"))
	forecast = sh.RunCmd(filepath.Join(BinDir, "forecast"))
	pinboard = sh.RunCmd(filepath.Join(BinDir, "pinboard-public"))

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
