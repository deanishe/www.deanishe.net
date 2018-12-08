---
title: "Copying Alfred's environment to your shell"
date: 2018-12-08
platforms: ["macOS", "Alfred"]
tags: ["alfred", "environment", "plist", "shell"]
draft: false
---

A simple script to create an Alfred-like environment in your shell
by extracting and exporting variables from info.plist.

<!--more-->

Alfred uses [environment variables][alfred-vars] to pass some important
information to your workflow code, like the paths to the data and cache
directories.

Source the following script in your shell and/or test runner to extract
Alfred workflow variables from `info.plist`:

```bash
# When sourced, creates an Alfred-like environment

# getvar <name> | Read a value from info.plist
getvar() {
    local v="$1"
    /usr/libexec/PlistBuddy -c "Print :$v" info.plist
}

# Minimal workflow settings
export alfred_workflow_bundleid=$( getvar "bundleid" )
export alfred_workflow_name=$( getvar "name" )
# Required if you're using AwGo's built-in update mechanism
export alfred_workflow_version=$( getvar "version" )

# Standard locations of Alfred 3 cache & data directories
export alfred_workflow_cache="${HOME}/Library/Caches/com.runningwithcrayons.Alfred-3/Workflow Data/${alfred_workflow_bundleid}"
export alfred_workflow_data="${HOME}/Library/Application Support/Alfred 3/Workflow Data/${alfred_workflow_bundleid}"

# If you want to extract workflow variables set in the workflow
# configuration sheet:
# export SCHEDULE_DAYS=$( getvar "variables:SCHEDULE_DAYS" )
# export EVENT_CACHE_MINS=$( getvar "variables:EVENT_CACHE_MINS" )
```


[alfred-vars]: https://www.alfredapp.com/help/workflows/script-environment-variables/ "Alfred's environment variables"

