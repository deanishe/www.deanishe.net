// This file exists to stop mage/go list complaining about a lack of
// .go files in 1.13

package main

import (
	_ "github.com/magefile/mage"
	_ "github.com/magefile/mage/mg"
	_ "github.com/magefile/mage/sh"
)
