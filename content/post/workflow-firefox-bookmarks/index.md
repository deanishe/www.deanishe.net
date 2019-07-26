---
title: "Workflow: Firefox Bookmarks"
slug: workflow-firefox-bookmarks
date: "2019-06-08"
tags:
    - alfred-workflow
    - firefox
    - bookmarks
platforms:
    - Alfred
    - Firefox
---

Alfred 4+ workflow to search and open Firefox bookmarks.

<!--more-->

Setup
-----

The workflow reads Firefox's bookmarks.html file, so first you must make sure that Firefox is configured to automatically export this file:

1. Browse to `about:config` in Firefox
2. Set `browser.bookmarks.autoExportHTML` to `true`
3. Restart Firefox to export the bookmarks.html file

**Note:** Firefox only updates the bookmarks.html file when it quits, so new bookmarks unfortunately aren't available in Alfred until you restart Firefox.


### Profiles

By default, the workflow reads the bookmarks.html file for the default Firefox profile. If you wish to read a different profile's bookmarks, set `profile_name` in the workflow's configuration sheet to the appropriate name, e.g. `dev-edition-default`.
