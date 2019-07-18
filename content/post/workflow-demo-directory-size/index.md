---
title: "Demo workflow: Directory size"
date: "2019-07-18"
draft: false
tags:
    - alfred-workflow
    - demo
platforms:
    - Alfred
---

Alfred workflows that shows the size of a directory.

<!--more-->

A specific directory can be specified in the workflow's
configuration sheet. Use keyword `dirsize` to show the
size of this directory.

There is also a "Show Size" File Action, which shows
the size of any directory.

As the size of a directory is calculated on-the-fly, the
workflow can be slow on large directory trees.

Under the hood, it uses `du -h -d 0`.

