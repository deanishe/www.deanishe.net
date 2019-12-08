---
title: "Pinboard Bookmarks to Chrome"
date: "2019-12-06T00:37:23+01:00"
draft: false
tags: ['pinboard', 'chrome', 'bookmarks', 'python']
platforms: ['macos']
---

Script to Overwrite a Chrome profile's bookmarks with your [Pinboard][pinboard] bookmarks.

<!--more-->

The idea is to overwrite the bookmarks of an *unused* Chrome profile with your Pinboard bookmarks to get rudimentary integration with apps (such as [Alfred](/tags/alfred)) that understand Chrome's bookmarks, but not Pinboard's.

Don't use it with an active profile: the generated file isn't a valid Chrome bookmarks file, and overwriting apps' private data is a bad idea in any case.

Usage
-----

The script requires two settings, the name of the Chrome profile to use and your Pinboard API key.

You can either set the `PROFILE_NAME` and `PB_API_TOKEN` variables in the script itself, or pass them via the `PROFILE` and `PINBOARD_API_TOKEN` environment variables respectively.

[Here's where you can find your Pinboard API key][api-key].

Create a new Chrome profile for use with this script and set the `PROFILE` environment variable to its name.

To view your available Chrome profiles, run the script with the `-l` option, which will list the names of your profiles and the directories they're stored in. **Take care not to overwrite an important profile**.

Run the script every hour or so using `cron` or a [launch agent][launchd]. You can [create a launch agent online here][launched] or use [the awesome LaunchControl app][launch-control].


Script
------

{{< code lang="python" script="pb2chrome.py" >}}


[pinboard]: https://pinboard.in
[launchd]: https://www.launchd.info/
[launched]: http://launched.zerowidth.com/
[launch-control]: https://www.soma-zone.com/LaunchControl/
[api-key]: https://pinboard.in/settings/password
