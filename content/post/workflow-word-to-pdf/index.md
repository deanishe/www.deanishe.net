---
title: "Workflow: Word to PDF"
date: "2021-05-05T17:49:22+02:00"
draft: false
tags:
    - alfred-workflow
    - microsoft
    - word
    - pdf
    - export
platforms:
    - Alfred
    - macOS
    - Microsoft Word
---

An Alfred 4+ workflow to export Word documents to PDF. Uses Microsoft Word for the best possible conversion.

<!--more-->

The workflow can be run via its **Export as PDF** File Action or by searching for a Word document to convert with the keyword `word2pdf`. It can either use macOS's built-in ability to generate PDFs via the Print dialog (the default) or Word's built-in PDF export, which is based on a web service provided by Microsoft.


## Installation ##

Download [Word to PDF.alfredworkflow][workflow] and double-click to install.


## Usage ##

- `word2pdf <query>` — Select a Word document to export.
    - `↩` — Export to selected document to PDF.
- `Export as PDF` — [File Action][actions] to export selected files.
- `File Buffer` — Don't forget you can collect files in [Alfred's File Buffer][file-buffer] and send them to the File Action all at once.

## Configuration ##

The workflow has two configuration settings:

`OPEN_WITH` (optional)
: Set to the name of the application you'd like to open the exported PDFs with. Default value is `Preview`. Set to `Finder` to reveal the PDFs in Finder. Clear the value to not open the PDFs.

`USE_ONLINE_SERVICE` (optional)
: Set to `true` to use Word's native PDF export (i.e. Microsoft's online service). This can be much faster and more reliable than UI scripting the Print dialog, but it has obvious privacy implications.


This workflow was built for [this thread on the Alfred forum][request].

[actions]: https://www.alfredapp.com/help/features/actions/
[file-buffer]: https://www.alfredapp.com/help/features/file-search/#file-buffer
[request]: https://www.alfredforum.com/topic/16844-converting-a-word-document-to-pdf-and-opening-in-a-program/
[workflow]: Word%20to%20PDF.alfredworkflow
