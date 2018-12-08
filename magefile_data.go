// +build mage

package main

import (
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Data fetch all remote data
func Data() {
	mg.Deps(Repos, Events, Weather, Posts)
}

// Repos fetch list of projects/repos
func Repos() error {
	if err := os.MkdirAll(filepath.Dir(ReposFile), 0755); err != nil {
		return err
	}
	return sh.Run(filepath.Join(BinDir, "github"), "repos", ReposFile)
}

// Events fetch list of events/actions
func Events() error {
	if err := os.MkdirAll(filepath.Dir(EventsFile), 0755); err != nil {
		return err
	}
	return sh.Run(filepath.Join(BinDir, "github"), "events", EventsFile)
}

// Weather fetch weather forecast data
func Weather() error {
	if err := os.MkdirAll(filepath.Dir(ForecastFile), 0755); err != nil {
		return err
	}
	return sh.Run(filepath.Join(BinDir, "forecast"), ForecastFile)
}

// Posts fetch recent posts from Pinboard
func Posts() error {
	if err := os.MkdirAll(filepath.Dir(PostsFile), 0755); err != nil {
		return err
	}
	return sh.Run(filepath.Join(BinDir, "pinboard-public"), PostsFile)
}
