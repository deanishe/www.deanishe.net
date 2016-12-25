---
title: Auto-swap day/night Alfred themes
date: 2016-12-25T16:32:05+01:00
draft: false
platforms: ["Alfred", "OS X"]
tags: ["alfred", "theme", "launchd", "flux"]
---

Swap Alfred theme at sunrise and sunset. Goes well with [f.lux][flux] and
macOS's night mode.

It uses [launchd][launchd] to call itself at sunrise and sunset (or
as soon thereafter as your machine boots/wakes), and tells
[Alfred][alfred] to change its theme.

<!--more-->

```none
toggle_alfred_theme.py [-v|-q] [-d|--dark <theme>] [-l|--light <theme>]

    **You must first edit this script (or the accompanying bash wrapper)
    to set your location!**

    These scripts (Python script and bash wrapper) live at:
    https://gist.github.com/deanishe/ce442c3a768adedc9c39

    Change Alfred's theme depending on whether it's dark outside. Theme
    is changed immediately when the script is run and the script will
    also call itself again at sunrise and sunset (via launchd) to change
    Alfred's theme. It's works well together with F.lux, which can
    switch to Yosemite's dark mode at sunset.

    Just run the script *once* with your preferred themes:

        python toggle_alfred_theme.py --dark 'Dark Theme' --light 'Light Theme'

    or if you're using the wrapper:

        toggle-alfred-theme.bash --dark 'Dark Theme' --light 'Light Theme'

    and it will ensure Alfred's theme is changed every day at sunrise
    and sunset.

    To change your preferred themes, just run the script again.

    Note: Because the script calls itself via launchd, if you move the
    script, it will stop working until you run it again.

Usage:

    toggle_alfred_theme.py (-h|--help)
    toggle_alfred_theme.py --timezones
    toggle_alfred_theme.py (-t|--times)
    toggle_alfred_theme.py [-n] [-v|-q] [--dark <theme>] [--light <theme>]

Options:
    -h, --help           Show this help message
    -n, --nothing        Show what would be set, but make no changes
    -t, --times          Show sunrise and sunset times for next 7 days
    --timezones          Show a list of (>500) supported timezones
    -l, --light <theme>  Alfred theme to use after sunrise
    -d, --dark <theme>   Alfred theme to use after sunset
    -v, --verbose        Show debugging info
    -q, --quiet          Only show warnings and errors

Installation & Setup:

    This script requires the `astral` and `pytz` libraries. Install with:

        pip install astral

    It's better to install them in the same directory as this script (or
    use a virtualenv), in order not to muck up your Python installation
    or break other software:

        pip install --target=/directory/where/this/script/is astral

    Adjust the settings at the top of this script in the CONFIG section
    (or in the bash wrapper) to match your location.

    `TZ_NAME` must be one of the timezones recognised by `pytz`. To see
    a list of all supported timezones, run this script with the
    --timezones option. (Note there are over 500 timezones.)

    You can usually find your town's latitude, longitude and elevation
    on its Wikipedia page.

How it works:

    When run, the script will immediately set Alfred's theme according
    to whether it's light or dark out, then tell OS X to run the script
    again at the next sunset/sunrise. Even if your computer is off/asleep
    when the script is supposed to run, it will be run immediately on
    boot/wake.

    Note: Yosemite has some issues with running LaunchAgents on wake. If
    the script isn't running when it's supposed to on Yosemite, but the
    script reports the correct times, it's a problem with Yosemite, not
    this script.

    The script has to fork into the background (i.e. exit successfully
    immediately) because `launchctl` doesn't like the script updating
    the Launch Agent while it's running it.
```


There are two files, the Python script and a bash wrapper. The wrapper's
only purpose is to store your configuration, so you can upgrade the
Python script without having to edit it again.

You **must** set your location in one of the scripts before usage.

You can usually find the latitude, longitude and elevation of your town
on its Wikipedia page.


{{< gist deanishe ce442c3a768adedc9c39 >}}


[alfred]: https://www.alfredapp.com/
[flux]: https://justgetflux.com
[launchd]: http://www.launchd.info
