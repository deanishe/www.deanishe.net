---
title: "Demo workflow: Random Images"
date: "2020-06-13T10:21:33+02:00"
draft: false
platforms:
    - Alfred
tags:
    - alfred-workflow
    - demo
---

Simple workflow that opens a given number of random images at a given interval.

<!--more-->

Usage
-----

`randimg <interval> <count>` — Show `<count>` images at intervals of `<interval>` seconds.


Configuration
-------------

The workflow can be configured via its configuration sheet.

|   Variable   |   Default    |               Description               |
|--------------|--------------|-----------------------------------------|
| `APP_NAME`   | `Preview`    | Application to open images in           |
| `IMAGE_DIR`  | `~/Pictures` | Directory to read images from           |
| `PLAY_SOUND` | `true`       | Play a sound halfway through and at end |

If Preview is the chosen application, the images are opened in fullscreen mode.

From [this forum thread][thread].

[thread]: https://www.alfredforum.com/topic/15077-get-random-image-from-folder-for-x-seconds/