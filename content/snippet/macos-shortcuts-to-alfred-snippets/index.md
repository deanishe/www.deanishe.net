---
title: "Sync macOS Shortcuts to Alfred Snippets"
date: "2020-06-22T12:04:12+02:00"
draft: false
tags:
    - alfred
    - macos
    - snippets
    - shortcuts
    - python
platforms:
    - Alfred
    - macOS
---

Python script to sync macOS/iOS system shortcuts to an [Alfred][alfredapp] snippet collection.

<!--more-->

The script syncs your macOS shortcuts as defined in System Preferences > Keyboard > Text to an Alfred snippets collection.

As the script syncs macOS snippets to Alfred, **it will remove all other snippets**, so don't put anything else in that collection.


## Why? ##

To have your iOS snippets work properly on your Mac. macOS snippets don't work everywhere and aren't as simple to use as Alfred ones.


## Usage ##

You can run the script from wherever, but unless you run it from your Alfred snippets directory, you'll at least need to set the `SNIPPETS_DIR` environent variable (see below), so Alfred sees the snippets.


## Configuration ##

The script has two options, set by environment variables:

`SNIPPETS_DIR` (default: `.`)
: Directory your Alfred snippets are in. In most cases, this should be the path to the `snippets` subdirectory of your `Alfred.alfredpreferences` bundle. The default location (i.e. you aren't syncing your Alfred preferences) would be `~/Library/Application Support/Alfred/Alfred.alfredpreferences/snippets`.

`COLLECTION_NAME` (default: `macOS`)
: Name of Alfred snippet collection to sync macOS shortcuts to

The script will create Alfred snippets in directory `$SNIPPETS_DIR/$COLLECTION_NAME`.

{{< code script="shortcuts2alfred.py" lang="python" >}}

Inspired by [this thread on the Alfred forums][thread].

[alfredapp]: https://www.alfredapp.com
[thread]: https://www.alfredforum.com/topic/15130-syncing-icloud-text-shortcuts-to-snippets/
