---
title: "Copying Alfred's environment to your shell"
date: 2018-12-08
platforms:
    - macOS
    - Alfred
tags:
    - alfred
    - environment
    - plist
    - shell
---

A simple script to create an Alfred-like environment in your shell
by extracting and exporting variables from info.plist.

<!--more-->

Alfred uses [environment variables][alfred-vars] to pass some important
information to your workflow code, like the paths to the data and cache
directories.

Source the following script in your shell and/or test runner to extract
Alfred workflow variables from `info.plist`:

{{< code script="alfredenv.sh" lang="bash" >}}

[alfred-vars]: https://www.alfredapp.com/help/workflows/script-environment-variables/ "Alfred's environment variables"
