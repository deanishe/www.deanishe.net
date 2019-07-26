---
title: "Demo workflow: Search engine List Filter"
date: "2019-07-26T19:00:00.000Z"
draft: false
tags:
    - alfred-workflow
    - demo
platforms:
    - Alfred
---

Alfred workflow that allows you to choose from a list
of search engines, then enter the query to search for.

<!--more-->

It demonstrates the use of the [Args & Vars utility][argsvars]
to capture more than one input from the user (by saving the first
to a workflow variable), and the [JSON utility][jsonutil] to
dynamically configure an Open URL action to use the selected URL.

The latter means that Alfred still takes care of URL-encoding the
search query for you.

Based on [this forum thread][thread].

[argsvars]: https://www.alfredapp.com/help/workflows/utilities/argument/
[jsonutil]: https://www.alfredapp.com/help/workflows/utilities/json/
[thread]: https://www.alfredforum.com/topic/13380-how-to-use-web-searches-in-a-list-filter/
