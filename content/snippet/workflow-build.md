---
date: 2016-12-22T19:14:00+01:00
draft: false
tags: ["python", "alfred"]
title: "Workflow build script for Alfred"
platforms: ["OS X", "Alfred"]
---

A script for building [Alfred][alfred] workflows. Focussed on Python-based workflows.

Creates an `.alfredworkflow` file from the contents of the specified
directory. The generated file's name is based on the workflow's name
and version extracted from `info.plist`.

<!--more-->

```none
workflow-build [options] <workflow-dir>

Build Alfred Workflows.

Compile contents of <workflow-dir> to a ZIP file (with extension
`.alfredworkflow`).

The name of the output file is generated from the workflow name,
which is extracted from the workflow's `info.plist`. If a `version`
file is contained within the workflow directory, it's contents
will be appended to the compiled workflow's filename.

Usage:
    workflow-build [-v|-q|-d] [-f] [-o <outputdir>] <workflow-dir>...
    workflow-build (-h|--version)

Options:
    -o, --output=<outputdir>    Directory to save workflow(s) to.
                                Default is current working directory.
    -f, --force                 Overwrite existing files.
    -h, --help                  Show this message and exit.
    -V, --version               Show version number and exit.
    -q, --quiet                 Only show errors and above.
    -v, --verbose               Show info messages and above.
    -d, --debug                 Show debug messages.
```


Source code
-----------

{{< gist deanishe b16f018119ef3fe951af >}}

[alfred]: https://www.alfredapp.com