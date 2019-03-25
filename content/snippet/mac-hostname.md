---
title: "Set the hostname on macOS"
date: "2019-03-25"
draft: false
platforms:
    - macOS
tags:
    - macos
    - network
    - hostname
---

Use `scutil` to set a Mac's FQDN.

```bash
sudo scutil --set HostName fqdn.example.com
```

<!--more-->

## Other commands

```bash
scutil --get HostName        # fqdn, same as hostname -f
scutil --get LocalHostName   # same as hostname -s
scutil --get ComputerName    # your computer's "friendly" name
```

## More info

- [manpage](https://ss64.com/osx/scutil.html)
- [tldr](https://github.com/tldr-pages/tldr/blob/master/pages/osx/scutil)
