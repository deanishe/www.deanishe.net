---
title: Multiple app instances and scripting weirdness
date: 2017-06-23T10:12:40+02:00
draft: false
platforms:
    - macOS
tags:
    - jxa
    - apps
    - macos
---

It's possible to start multiple instances of an application with `open -n`:

<!--more-->

```sh
#!/bin/bash
# First instance of application
open -a 'Safari' 'https://www.google.com'
# Open a second application process
open -n -a 'Safari' 'https://www.yahoo.com'
```

Any subsequent calls to `Safari` via AppleScript will go to the newest instance of the application, but weirdly, JXA calls go to the first:

```sh
#!/bin/bash
open -a 'Safari' 'https://www.google.com'
sleep 10  # give app and page time to load
osascript -l JavaScript -e "Application('Safari').windows[0].currentTab.name()"
# Google
osascript -e 'tell application "Safari" to return the name of the current tab of the first window as text'
# Google

# Open a new instance
open -n -a 'Safari' 'https://www.yahoo.com'
sleep 10  # give app and page time to load
osascript -l JavaScript -e "Application('Safari').windows[0].currentTab.name()"
# Google
osascript -e 'tell application "Safari" to return the name of the current tab of the first window as text'
# Yahoo
```

To talk to a specific instance, sort them by PID. The lowest PID is the oldest instance, the highest the newest:

```javascript
// Compare PIDs
function sortByPid(proc1, proc2) {
    var pid1 = proc1.unixId()
    var pid2 = proc2.unixId()
    if (pid1 < pid2) return -1
    if (pid2 < pid1) return 1
    return 0
}

// Return array of processes, sorted by PID
function namedProcesses(name) {
    var results = []
    var procs = Application('System Events').processes.whose({name: name})
    for (i=0; i<procs.length;i++) {
        results.push(procs[i])
    }
    results.sort(sortByPid)
    return results
}

function run() {
    var procs = namedProcesses('Safari')
    var oldest = procs[0]
    var newest = procs[procs.length-1]
    console.log('oldest proc (' + newest.unixId() + ')', oldest)
    console.log('newest proc (' + oldest.unixId() + ')', newest)
}
```

Unfortunately, I haven't yet figured out how to get a Safari `Application` from the Safari `Process` object returned by System Events.
