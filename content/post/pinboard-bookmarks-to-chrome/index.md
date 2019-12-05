---
title: "Pinboard Bookmarks to Chrome"
date: "2019-12-06T00:37:23+01:00"
draft: false
tags: ['pinboard', 'chrome', 'bookmarks', 'python']
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

To view your available Chrome profiles, run the script with the `-l` option, which will list the names of your profiles and the directories they're stored in. **Take care not to screw up and important profile**.

Run the script every hour or so using `cron` or a [launch agent][launchd]. You can [create a launch agent online here][launched] or use [the awesome LaunchControl app][launch-control].


Script
------

[You can download the script here](pb2chrome.py.txt). Remove the `.txt` extension and make the script executable if necessary.

```python
#!/usr/bin/env python
# encoding: utf-8
#
# Copyright (c) 2019 Dean Jackson <deanishe@deanishe.net>
# MIT Licence. See http://opensource.org/licenses/MIT
# Created on 2019-12-05


"""pb2chrome.py [-h|-l]

Overwrite a Google Chrome profile's bookmarks with your Pinboard bookmarks.

You must set the following environment variables for the script to work:

PROFILE_NAME
    The name of the profile whose bookmarks should be overwritten.
    You should create a new profile for use exclusively by this script.

PINBOARD_API_TOKEN
    Your Pinboard API key. Find it here:
    https://pinboard.in/settings/password

Options:
    -l   list your Chrome profiles
    -h   show this message and exit
"""

from __future__ import print_function, absolute_import

from collections import namedtuple
from getopt import getopt
import json
import os
import sys
try:
    from urllib2 import urlopen, URLError
except ImportError:
    from urllib.request import urlopen, URLError

# Name of Chrome profile to write Pinboard bookmarks to
PROFILE = 'Pinboard Sync'

# Pinboard API token from https://pinboard.in/settings/password
# You can set the API token here in the script or via the
# PINBOARD_API_TOKEN environment variable.
PB_API_TOKEN = ''

PROFILE = os.getenv('PROFILE_NAME') or PROFILE
PB_API_TOKEN = os.getenv('PINBOARD_API_TOKEN') or PB_API_TOKEN
# URL to fetch all bookmarks from Pinboard
PB_URL = 'https://api.pinboard.in/v1/posts/all?auth_token={token}&format=json'


Profile = namedtuple('Profile', 'name dir')


def log(s, *args, **kwargs):
    """Log to STDERR."""
    if args:
        s = s % args
    elif kwargs:
        s = s % kwargs
    print(s, file=sys.stderr)


def get_chrome_profiles():
    """Get names & paths of Chrome profiles."""
    profiles = []
    dirpath = os.path.expanduser('~/Library/Application Support/Google/Chrome/')
    with open(os.path.join(dirpath, 'Local State')) as fp:
        state = json.load(fp)
    for dirname, data in state['profile']['info_cache'].items():
        name = data['name']
        p = Profile(data['name'], os.path.join(dirpath, dirname))
        profiles.append(p)

    profiles.sort(key=lambda p: p.dir)
    return profiles


def get_pinboard_bookmarks():
    """Retrieve bookmarks from Pinboard API."""
    r = urlopen(PB_URL.format(token=PB_API_TOKEN))
    log('[%d] %s', r.getcode(), r.geturl().replace(PB_API_TOKEN, '<secret>'))
    if r.getcode() > 200:
        raise URLError('bad response: %d' % r.getcode())
    return json.load(r)


def convert_bookmark(bookmark):
    """Convert Pinboard bookmarks to Chrome bookmarks."""
    return {
        'name': bookmark['description'],
        'url': bookmark['href'],
        'type': 'url',
    }


def write_chrome_bookmarks(bookmarks, path):
    """Save Chrome bookmarks to file."""
    data = {
        'roots': {
            'bookmark_bar': {
                'children': bookmarks,
                'name': 'Pinboard',
                'type': 'folder',
            },
        },
        'version': 1,
    }
    with open(path, 'w') as fp:
        json.dump(data, fp, indent=2, separators=(',', ': '))


def main():
    """Run script."""
    opts, _ = getopt(sys.argv[1:], 'hl')
    for opt, _ in opts:
        if opt == '-h':
            print(__doc__)
            return
        if opt == '-l':
            for p in get_chrome_profiles():
                print('%s -- %s' % (p.name, p.dir))
            return

    if not PB_API_TOKEN:
        raise ValueError(
            'Pinboard API token not set. Set the PB_API_TOKEN variable '
            'in this script or the environment variable PINBOARD_API_TOKEN.'
        )
    if not PROFILE:
        raise ValueError(
            'Chrome profile not set. Set the PROFILE variable in this script '
            'or the environment variable PROFILE_NAME.'
        )

    for profile in get_chrome_profiles():
        if profile.name == PROFILE:
            break
    else:
        raise ValueError('Chrome profile "%s" not found' % PROFILE)

    path = os.path.join(profile.dir, 'Bookmarks')
    pinboard = get_pinboard_bookmarks()
    log('%d bookmark(s) loaded from Pinboard', len(pinboard))
    chrome = [convert_bookmark(d) for d in pinboard]
    write_chrome_bookmarks(chrome, path)
    log('Saved bookmarks to Chrome profile "%s" (%s)' % (profile.name, path))


if __name__ == '__main__':
    main()
```

[pinboard]: https://pinboard.in
[launchd]: https://www.launchd.info/
[launched]: http://launched.zerowidth.com/
[launch-control]: https://www.soma-zone.com/LaunchControl/
[api-key]: https://pinboard.in/settings/password
