---
title: "Workflow: Video Conferences"
date: "2020-05-01T12:06:23+02:00"
draft: false
tags:
    - alfred-workflow
    - calendar
    - video-conference
    - zoom
    - lync
platforms:
    - Alfred
    - Zoom
    - Lync
---

An Alfred 4+ workflow that shows upcoming calendar events that are video conferences, and opens the links to them. It supports Lync and Zoom out of the box, but can be easily configured to support additional services.

<!--more-->

The workflow was inspired by [this thread on the Alfred forum][thread].


## Installation ##

Download [Video Conferences.alfredworkflow](Video%20Conferences.alfredworkflow) and double-click to install.

**Note:** You must grant the workflow/Alfred permission to access your calendars or it won't work. Obviously.


## Usage ##

- `.vc [<query>]` — Search upcoming video-conference events
    - <kbd>⌘↩</kbd> — Open video-conference URL in your default browser
- `.vc reload` — Force reload the cached list of events


## Configuration ##

The workflow has a few knobs to turn in its configuration sheet.

`lookahead_days` (default: `3`)
: The number of days to fetch events for.

`max_cache_seconds` (default: `1200`)
: How long (in seconds, obvs) to cache the list of events for. Fetching events is pretty slow, so adjust this to find the right compromise between speed and freshness. You can refresh the cached list of events at any time by entering the search query "reload".

`account_1` (default: empty)
: Set `account_*` variables to the names of accounts you'd like to limit the search to. For example, if you only want to search your iCloud account, set `account_1` to `iCloud`. If you'd also like to search your Nextcloud account, create another variable starting with `account_` and set it to `Nextcloud` etc.

`calendar_1`, `calendar_2` (default: empty)
: As with `account_*` variables, you can use variables starting with `calendar_` to restrict the search to only calendars whose names match the variables, e.g. if you've set `calendar_1` to `Work`, only calendars named "Work" will be searched.

`regex_zoom`, `regex_lync` etc.
: Variables starting with `regex_` are used to search for video conference URLs. There are currently regular expressions for Zoom (`regex_zoom`) and Lync (`regex_lync`). Add new variables starting with `regex_` to add regular expressions that match the URLs of other services you'd like to add.


[aw]: https://www.deanishe.net/alfred-workflow/
[magic]: https://www.deanishe.net/alfred-workflow/guide/magic-arguments.html
[thread]: https://www.alfredforum.com/topic/12894-workflow-to-get-next-meeting-locationurl-and-open-it/