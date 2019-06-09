---
title: "Demo workflow: Search in App"
date: "2019-06-09"
draft: false
tags:
    - alfred-workflow
    - demo
    - jxa
platforms:
    - Alfred
---

Alfred workflow showing how to pass a query to an application's own search by simulating `⌘F`.

<!--more-->

For applications that aren't scriptable via AppleScript or URL scheme, you can simulate keypresses to open an app's own search function and enter a query.

This workflow activates the configured application (Safari by default), simulates `⌘F`, waits a bit, types the user's query and simulates `↩`.

