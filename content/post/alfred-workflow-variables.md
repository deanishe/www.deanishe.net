---
title: "Workflow/environment variables in Alfred"
date: 2018-10-14
tags:
    - alfred
    - variables
    - tutorial
    - guide
aliases:
    - "/post/2018/10/workflow/environment-variables-in-alfred/"
---

This is a brief look at how to get, set and save variables in code
(i.e. in Script Filters, Run Script Actions, etc.).

<!--more-->

<!-- MarkdownTOC autolink=true -->

- [Introduction](#introduction)
- [Setting variables](#setting-variables)
  - [From Run Script actions](#from-run-script-actions)
  - [From Script Filters](#from-script-filters)
    - [Root-level variables](#root-level-variables)
    - [Item-level variables](#item-level-variables)
    - [Modifier-level variables](#modifier-level-variables)
- [Using variables](#using-variables)
  - [bash](#bash)
  - [Python](#python)
  - [AppleScript](#applescript)
  - [JavaScript (JXA)](#javascript-jxa)
  - [PHP](#php)
  - [Ruby](#ruby)
- [Saving variables](#saving-variables)
  - [AppleScript](#applescript-1)
  - [JavaScript (JXA)](#javascript-jxa-1)

<!-- /MarkdownTOC -->


## Introduction ##

In Alfred 2, you had one single variable to work with: the `{query}`
macro. Alfred 3 adds the ability to specify as many variables as you
want. [Alfred's own help][alfred-help] provides a great description of
working with variables in Alfred's own UI. I'm going to look more
closely about getting and setting workflow/environment variables in
your own code within a workflow.

First of all, it bears mentioning that all variables are strings. Sure,
you can set a variable to a number in JSON, but when it reaches your
next script or one of Alfred's Filter Utilities, it will be a string.
If you set a variable to an array (e.g.
`[1, 2, 3, "mach dat Mäh mal ei"]`), Alfred will turn it into a single, tab-delimited string (`"1\t2\t3\tmach dat Mäh mal ei"`).


## Setting variables ##

There are several ways to set variables. The most obvious ones are in
the [Workflow Environment Variables table][config-sheet] in the
workflow configuration sheet and using the [Arg and Vars
Utility][args-and-vars]. The configuration sheet is largely without
magic, but in an Args and Vars Utility, you can use variable expansion
macros: `{query}` expands (as always) to the input (which may be a
user-entered query or the output from a previous element), and you can
use `{var:VARIABLE_NAME}` macros for your own custom variables.  This
is described in detail in the above-mentioned Alfred help pages.

More interestingly, you can also set variables via the output of your
scripts (i.e. dynamically) by emitting appropriate JSON. How you set
variables depends on whether you are using a Script Filter or a Run
Script action.

**NOTE: You must use the appropriate mechanism, or it won't work!**


### From Run Script actions ###

Let's say your script outputs a URL, e.g. https://www.google.com.
Normally you just do `print('https://www.google.com')` (or `echo` or
`puts`) and that gets passed as the input to the next action. To also
pass variables, you instead emit JSON in a very specific format:

```json
{"alfredworkflow": {
    "arg": "https://www.google.com",
    "variables": {"browser": "Google Chrome"}}}
```

The root `alfredworkflow` object is required. If it's missing, Alfred
won't parse the JSON, but will pass it as-is as input to the next
action (which can also be very useful). Your output (i.e. the next
Action's input/`{query}`) goes in `arg`, and any variables you wish to
set go in the `variables` object.


### From Script Filters ###

You can also set workflow variables via Script Filter feedback at three
different levels: the **root** level, the **item** level and the
**modifier** level. (**Note**: This only applies to JSON feedback; XML
feedback is now deprecated and does not support the features described
here.)

In each case, variables are set via a `variables` object at the
appropriate level (feedback root, `item` or `mod`).


#### Root-level variables ####

Root-level variables are always passed to downstream elements
regardless of which item is actioned. They are also passed back to the
same Script Filter if you've set `rerun`, so you can use root-level
variables to implement a [progress bar][progress-bar].

`browser` is set to `Safari` for all items:

```json
{"variables": {"browser": "Safari"},
 "items": [{"title": "Google",
   "arg": "https://www.google.com"}]}
```


#### Item-level variables ####

Item-level variables are only passed downstream when the item they're set on is actioned, and they override root-level variables. Root-level variables are also passed downstream when you action an item.

browser is set to `Safari` by default, but `Google Chrome` for `Reddit`:

```json
{"variables": {"browser": "Safari"},
 "items": [
   {"title": "Google",
     "arg": "https://www.google.com"},
   {"title": "Reddit",
     "arg": "https://reddit.com",
     "variables": {"browser": "Google Chrome"}}]}
```


#### Modifier-level variables ####

Modifier-level variables are only passed downstream when the corresponding `item` is actioned with the appropriate modifier key pressed. They **replace** item-level variables (i.e. if a modifier sets any variables, Alfred ignores any variables set on its parent `item`) and override root-level variables.

As above, `browser` is set to `Safari` by default and `Google Chrome` for Reddit. But you can also pass `browser=Google Chrome` for Google by holding ⌘ when actioning it:

```json
{"variables": {"browser": "Safari"},
 "items": [
   {"title": "Google",
     "arg": "https://www.google.com",
     "mods": {"cmd": {"variables": {"browser": "Google Chrome"}}}},
   {"title": "Reddit",
     "arg": "https://reddit.com",
     "variables": {"browser": "Google Chrome"}}]}
```


## Using variables ##

So you've set a few variables, and now you want to use them. Within
Alfred elements like [Arg and Vars][args-and-vars] or [Filter][filter]
Utilities, you use the above-mentioned `{var:VARIABLE_NAME}` macros.
Very simple.

Where it gets a little more complicated is in your own code. First and foremost, __`{var:VARIABLE_NAME}` macro expansion does not work in Run Script actions, Script Filters or any other script boxes in Alfred.__

When Alfred runs your code, it does not use `{var:...}` macros, but
rather takes any workflow variables and sets them as environment
variables for your script. Using the above example again, Alfred would
pass "https://www.google.com" to my script as input (either via ARGV or
`{query}` depending on the settings) and it would set the environment
variable `browser` to `Safari` or `Google Chrome`. How you retrieve
environment variables depends on the language you're using.


### bash ###

The variables are already in the global namespace. Just use `$browser`


### Python ###

Use the `os.environ` dictionary or `os.getenv('VARIABLE_NAME')`:

```python
import os
browser = os.environ['browser']

# Or
browser = os.getenv('browser')
```

### AppleScript ###

Use `system attribute`:

```applescript
set theBrowser to (system attribute "browser")
```


### JavaScript (JXA) ###

Use `$.getenv()`:

```javascript
ObjC.import('stdlib');
var browser = $.getenv('browser');
```


### PHP ###

Use `getenv()`:

```php
$browser = getenv('browser');

// Or
$browser = $_ENV['browser'];
```

(Please see [this comment by juliosecco][php-comment] on why you should
use `getenv()` over `$_ENV`.)


### Ruby ###

Use `ENV`:

```ruby
browser = ENV["browser"]
```


## Saving variables ##

Any variables you set in a running workflow are **not** saved. They
exist as long as the workflow is running and then disappear. Any
Workflow Environment Variables will "reset" to their values in the
[workflow configuration sheet][config-sheet] on the next run.

Generally, this is what you want, but sometimes you want to save a
variable's value. For example, you might have an `API_KEY` Workflow
Environment Variable in the configuration sheet. The user can enter
their API key for the service in the configuration sheet, but you'd
also like to add the ability to set it from within your workflow, e.g.
with a `setapikey` Keyword and corresponding Run Script action.

As of version 3.6, Alfred provides the `set configuration` and
`remove configuration` AppleScript functions to manipulate the
variables set in the Workflow Configuration Sheet.


### AppleScript ###

**NOTE:** The `with exportable` clause is optional. If not specified,
the variable defaults to "Don't Export".

To set variable `BROWSER` to value `Safari` in workflow
`net.deanishe.demo`:

```applescript
tell application "Alfred 3" to set configuration "BROWSER" to value "Safari" in workflow "net.deanishe.demo" with exportable
```

As Alfred exports the bundle ID of the running workflow to the
environment variable `alfred_workflow_bundleid`, you can use this
instead of hard-coding the bundle ID:

```applescript
set bundleID to (system attribute "alfred_workflow_bundleid")

tell application "Alfred 3"
    set configuration "BROWSER" to value "Safari" in workflow bundleID with exportable
end tell
```

The corresponding call to remove a variable is:

```applescript
tell application "Alfred 3" to remove configuration "BROWSER" in workflow "net.deanishe.demo"
```


### JavaScript (JXA) ###

The equivalents to the above in JXA JavaScript (again, the `exportable`
variable is optional):

```javascript
Application('Alfred 3').setConfiguration('BROWSER', {
    toValue: 'Safari',
    inWorkflow: 'net.deanishe.demo',
    exportable: true
});
```

Or using the `alfred_workflow_bundleid` variable:

```javascript
ObjC.import('stdlib');
Application('Alfred 3').setConfiguration('BROWSER', {
    toValue: 'Safari',
    inWorkflow: $.getenv('alfred_workflow_bundleid'),
    exportable: true
});
```


[alfred-help]: https://www.alfredapp.com/help/workflows/advanced/variables/
[progress-bar]: https://www.alfredforum.com/topic/9718-progress-bar/
[php-comment]: https://www.alfredforum.com/topic/9070-how-to-workflowenvironment-variables/?p=46151
[args-and-vars]: https://www.alfredapp.com/help/workflows/utilities/argument/
[config-sheet]: https://www.alfredapp.com/help/workflows/advanced/variables/#environment
[filter]: https://www.alfredapp.com/help/workflows/utilities/filter/
