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


Configuration
-------------

Change the name of the app in the workflow's configuration sheet in Alfred Preferences:

![Workflow configuration sheet][screen]

`delay` is the number of seconds the workflow waits after simulating `⌘F` before it types the query (i.e. time for the app's search dialog to appear and take focus).


[screen]: screenshot-setup.png "screenshot of Alfred's workflow configuration sheet"
