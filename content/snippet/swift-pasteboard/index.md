---
title: "Retrieve pasteboard contents by type on macOS"
date: "2020-02-16T22:54:57+01:00"
draft: false
platforms:
    - macOS
    - Swift
tags:
    - pasteboard
    - cli
    - swift
---

Swift CLI helper to retrieve pasteboard content by type.

<!--more-->

This script retrieves HTML and text content from the pasteboard, encodes it to JSON and writes the result to STDOUT.

{{< code script="pboard.swift" lang="swift" >}}

Run with:

```bash
swift pboard.swift
```

Build with:

```bash
xcrun -sdk macosx swiftc pboard.swift -o pboard
```