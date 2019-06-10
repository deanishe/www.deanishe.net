---
title: "Workflow: Search in App"
date: "2019-06-09"
draft: false
tags:
    - alfred-workflow
    - demo
    - jxa
platforms:
    - Alfred
aliases:
    - "/post/2019/06/demo-workflow-search-in-app/"
---

Alfred workflow to pass a query to an application's own search by simulating <kbd>⌘F</kbd>.

For applications that aren't scriptable via AppleScript or URL scheme, you can simulate keypresses to open an app's own search function and enter a query. This workflow is a relatively flexible implementation of that.

<!--more-->

[Download here](#downloads).

By default, the workflow puts your query on macOS's [Find Pasteboard](#find-pasteboard), activates the target application, waits for it to be visible, and then simulates <kbd>⌘F</kbd>.

If the application doesn't support the Find Pasteboard, set `use_find_pasteboard` to `0` (see [below](#configuration)) and the workflow attempts to type the query into the search field instead.

Finally, it simulates <kbd>↩</kbd> (by default) to start the search.


Configuration
-------------

Configure the workflow via its configuration sheet in Alfred Preferences:

![Workflow configuration sheet][screen]

### Options

`app_name` is the name of the application you want to search in.

`delay` is how long (in seconds) the workflow waits for the app's search dialog to activate after sending <kbd>⌘F</kbd> before it attemps to type the query or send <kbd>↩</kbd>.

If `press_return` is set to `1`, the workflow simulates <kbd>↩</kbd> to start the search.

`timeout` is how long (in seconds) the workflow waits for the application to become visible.

If `use_find_pasteboard` is `1`, the workflow puts the search query on macOS's [Find pasteboard](#find-pasteboard). Otherwise, the workflow attempts to type the search query in the application's search field.

If `wait_visible` is set to `1`, the workflow waits until the application is visible before simulating any keypresses.


Find Pasteboard
---------------

macOS has multiple pasteboards (clipboards), and one of these is called the Find Pasteboard (or Search Pasteboard). In applications that use macOS's native search, the Find Pasteboard is automatically used for the search query (including live updates). You can put the currently selected text on the Find Pasteboard with <kbd>⌘E</kbd>.

In Alfred 4+, the [Hotkey trigger][hotkey] can also pass the contents of the Find Pasteboard to your workflow.


[screen]: screenshot-setup.png "screenshot of Alfred's workflow configuration sheet"
[hotkey]: https://www.alfredapp.com/help/workflows/triggers/hotkey/ "Alfred's help for Hotkey triggers"
