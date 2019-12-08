---
title: "Text Table in Python"
date: 2016-12-25T16:33:06+01:00
draft: false
tags:
    - python
    - text
platforms:
    - Python
---

Python class to pretty-print tabluar data in a terminal.

```python
t = Table(titles=['Name', 'Position', 'Goals'])
t.add_row(['Dave Smith', 'striker', 12])
t.add_row(['Angus McGregor', 'full back', 1])
print(t)
```
produces:

```
     Name      | Position  | Goals
----------------------------------
Dave Smith     | striker   |    12
Angus McGregor | full back |     1
```

<!--more-->

Titles are centre-aligned, text left-aligned, and numbers
right-aligned.

{{< code script="table.py" lang="python" >}}
