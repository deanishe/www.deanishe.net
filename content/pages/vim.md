title: VIM Cheat Sheet
author: Dean Jackson <de@njackson.com>
date: 2013-09-13
save_as: vim.html


* [General](#general)
    * [Modes](#modes)
    * [Moving around](#movingaround)
    * [Window & buffer management](#window-buffer-management)
    * [Code folding](#codefolding)
    * [Custom mappings](#custommappings)
        * [General](#general_1)
        * [Window navigation](#window-navigation)
        * [Tab management](#tab-management)
        * [Plugins / External](#plugins-external)
    * [Search & replace](#search-replace)
* [Plugins](#plugins)
    * [Buffer Bye](#bufferbye). Close buffers in a sensible fashion.
    * [Command-T](#commandt). Faster CtrlP.
    * [GitGutter](#gitgutter). Show git status in the gutter.
    * [Indentpython.vim](#indentpythonvim). PEP-8-ish indentation.
    * [Pymode](#pymode). All-in-one Python toolset.
    * [SimpylFold](#simpylfold). Python folding.
    * [Tabular](#tabular). Align lines by specified characters.
    * [Tagbar](#tagbar). Navigate tags in source.
    * [CtrlP](#ctrlp). Works like CMD+T in TextMate.
    * [MarkdownFootnotes](#markdownfootnotes). Add footnotes.
    * [NERDTree](#nerdtree). Explore your filesystem.
    * [NERDCommenter](#nerdcommenter). Comment/uncomment.
    * [Pytest.vim](#pytestvim). Run `pytest` tests from vim.
    * [Sensible.vim](#sensiblevim). Solid editor defaults.
    * [Supertab](#supertab). Make `<TAB>` work like it should.
    * [Syntastic](#syntastic). Syntax checking for many file formats.
    * [RagTag](#ragtag). Easily add HTML.
    * [Repeat](#repeat). Repeat the last command. Intelligently.
    * [Sparkup](#sparkup). HTML generation from CSS selectors.
    * [BufferGator](#buffergator). Navigate buffers.
    <!--* [snipMate](#snipmate). Expand snippets like TextMate.-->
    * [UltiSnips](#ultisnips). Expand snippets.
    * [Surround.vim](#surroundvim). Wrap text in pairs of quotes/brackets/etc.
    * [Abolish.vim](#abolishvim). Refined, intelligent search and replace.
    * [Fugitive.vim](#fugitivevim). `git` integration.
    * [Tasklist](#tasklist). Show TODOs, FIXMEs
    * [Gundo](#gundo). Navigate `vim` undo trees.
    * [Scriptease](#scriptease). A script for making scripts.
    * [YouCompleteMe](#youcompleteme). Superb autocompletion.
    * [TaskPaper](#taskpaper). Process TaskPaper files in `vim`.
    * [Unimpaired](#unimpaired). Mappings of complementary pairs, e.g. encoding.
    * [Vim-go](#vimgo). Golang support.
    * [Vim-snippets](#vimsnippets). Collection of snippets.
    * [Installing plugins](#installingplugins)


## General ##


### Modes ###

| Name            | Key                        | Description                                    |
| --              | --                         | --                                             |
| normal          | `<ESC>`                    | Navigate document                              |
| insert          | `i` or `I` (insert at end) | For entering text                              |
| plain visual    | `v`                        | For navigation and manipulation of text        |
| block visual    | `<c-v>`                    | Always maintains a rectangular selection       |
| linewise visual | `<s-v>`                    | Highlight entire lines                         |
| select          |                            | Like `visual`, but Windows-y                   |
| command-line    |                            | For entering editor commands, e.g. `:help ...` |
| ex-mode         |                            | Like command-line, but for batch processing    |


#### Links ####

- [Vim modes at Wikibooks](http://en.wikibooks.org/wiki/Learning_the_vi_Editor/Vim/Modes)


### Moving around

* Left/up/down/right: `h`, `j`, `k` and `l`
* Move up (scroll down):
    * `CTRL+e` 1 line
    * `CTRL+d` 0.5 page
    * `CTRL+f` 1 page
* Move down (scroll up):
    * `CTRL+y` 1 line
    * `CTRL+u` 0.5 page
    * `CTRL+b` 1 page
* Jump back and forth between matching tags/brackets/keyword pairs (matchit.vim): `%` and `%g`


### Window & buffer management ###

|      Mapping      |  Command   |                     Action                     |
|-------------------|------------|------------------------------------------------|
| `CTRL+w -`        | `:split`   | Split window horizontally                      |
| `CTRL+w /`        |            | Split window vertically                        |
| `CTRL+w c`        | `:close`   | Close current window                           |
| `CTRL+k`/`CTRL+j` |            | Move to window above/below                     |
| `CTRL+h`/`CTRL+l` |            | Move to window to left/right                   |
|                   | `:bdelete` | Delete current buffer (closes window)          |
|                   | `:Bclose`  | Close current buffer (leaves window unchanged) |


### Code folding

Vim can't fold some filetypes by default. Add a `python.vim` file to
`~/.vim/after/ftplugin` or any file to `~/.vim/after/ftplugin/python/` with the
following content:

```viml
set foldmethod=indent
set foldnestmax=2
set foldlevel=2
" show a folding status bar on the left-hand side
set foldcolumn=1
```

This will set reasonable defaults (2 levels is method-level) and open the folds.

Control folds with the following commands:

| Mapping         | Action                          |
| --------------- | ------                          |
| `zo`            | Open current fold               |
| `zO`            | Open current fold recursively   |
| `zc`            | Close current fold              |
| `zf`            | Create a fold (in visual mode)  |
| `zf{motion}`    | Create a fold (in normal mode)  |
| `zd`            | Delete current fold             |
| `zD`            | Delete current fold recursively |
| `za`            | Toggle current fold             |
| `zA`            | Toggle current fold recursively |
| `zm`            | Fold more                       |
| `zr`            | Fold less                       |
| `zM`            | Fold everything                 |
| `zR`            | Unfold everything               |
| `:help folding` | VIM's own help                  |


### Custom mappings

`<Leader>` is `<space>`


#### General ####

| Mapping         | Action                                                  |
| --------------- | --------------------------------------------------      |
| `<Leader>e`     | Run current file with interpreter specified in hashbang |
| `<Leader>ev`    | Edit `~/.vimrc`                                         |
| `<Leader>rv`    | Reload (source) `~/.vimrc`                              |
| `w!!`           | Save with `sudo`                                        |
| `<Leader>x`     | Set executable bit                                      |
| `<Leader>l`     | Toggle relative line numbers                            |
| `<Leader>ll`    | Toggle line numbers (e.g. to copy from terminal)        |
| `<Leader>rts`   | Trim trailing whitespace                                |


#### Window navigation ####

| My Mapping | Default Mapping |          Action         |
| ---------- | --------------- | ----------------------- |
| `CTRL+h`   | `CTRL+w h`      | Move to window to left  |
| `CTRL+j`   | `CTRL+w j`      | Move to window below    |
| `CTRL+k`   | `CTRL+w k`      | Move to window above    |
| `CTRL+l`   | `CTRL+w l`      | Move to window to right |


#### Tab management ####

| Mapping      | Command     | Action            |
| ---          | ---         | ---               |
| `<leader>tn` | `:tabnew`   | Create a new tab  |
| `<leader>to` | `:tabonly`  |                   |
| `<leader>tc` | `:tabclose` | Close current tab |
| `<leader>tm` | `:tabmove`  | Move current tag  |
| `<leader>t`  |             | Show tab list     |


#### Plugins / External ####

| Mapping                | Action                                    |
| ---------------------- | ---------------------------------------   |
| `D-F1`                 | Open this cheat sheet in Marked 2.app     |
| `F2`                   | Toggle Buffergator                        |
| `F3`                   | Toggle NERDTree                           |
| `F4` or `<Leader>rt`   | Toggle Tagbar                             |
| `F5`                   | Run Script (no output)                    |
| `F6`                   | Run Script (show ouput)                   |
| `F7`                   | Toggle Paste Mode                         |
| `F8`                   | Toggle Gundo                              |
| `<Leader>p`            | Open [CtrlP](#ctrlp)                      |
| `<Leader>t`            | Open [Command-T](#commandt) files         |
| `<Leader>b`            | Open [Command-T](#commandt) buffers       |
| `<Leader>r`            | Open [Command-T](#commandt) tags          |
| `<Leader>c`            | Toggle comment                            |
| `<Leader>cl`           | Comment out, aligning left comments       |
| `<Leader>td`           | Show interactive list of TODO, FIXME etc. |


### Search & Replace ###


#### Basics ####

`:%s/foo/bar/g`
: Find each occurrence of `foo` (in all lines), and replace it with `bar`.

`:s/foo/bar/g`
: Find each occurrence of `foo` (in the current line only), and replace it with `bar`.

`:%s/foo/bar/gc`
: Change each `foo` to `bar`, but ask for confirmation first.

`:%s/\<foo\>/bar/gc`
: Change only whole words exactly matching `foo` to `bar`; ask for confirmation.

`:%s/foo/bar/gci`
: Change each `foo` (case insensitive) to `bar`; ask for confirmation.

This may be wanted after using `:set noignorecase` to make searches case
sensitive (the default).

`:%s/foo/bar/gcI`
: Change each `foo` (case sensitive) to `bar`; ask for confirmation. This may
be wanted after using `:set ignorecase` to make searches case insensitive.


#### Search scope ####

| Example                 | Meaning                                                                                                                                                                                                                                     |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `:s/foo/bar/g`          | Change each `foo` to `bar` in the current line.                                                                                                                                                                                             |
| `:%s/foo/bar/g`         | Change each `foo` to `bar` in all lines.                                                                                                                                                                                                    |
| `:5,12s/foo/bar/g`      | Change each `foo` to `bar` for all lines from line 5 to line 12 inclusive.                                                                                                                                                                  |
| `:'a,'bs/foo/bar/g`     | Change each `foo` to `bar` for all lines from mark a to mark b inclusive (see Note below).                                                                                                                                                  |
| `:'<,'>s/foo/bar/g`     | When compiled with +visual, change each `foo` to `bar` for all lines within a visual selection. Vim automatically appends the visual selection range (`'<,'>`) for any ex command when you select an area and enter:. Also, see Note below. |
| `:.,$s/foo/bar/g`       | Change each `foo` to `bar` for all lines from the current line (.) to the last line ($) inclusive.                                                                                                                                          |
| `:.,+2s/foo/bar/g`      | Change each `foo` to `bar` for the current line (.) and the two next lines (+2).                                                                                                                                                            |
| `:g/^baz/s/foo/bar/g`   | Change each `foo` to `bar` in each line starting with 'baz'.                                                                                                                                                                                |

**Note:** As of Vim 7.3, substitutions applied to a range defined by marks or a
visual selection (which uses a special type of marks `'<` and `'>`) are not
bounded by the column position of the marks by default. Instead, Vim applies
the substitution to the entire line on which each mark appears unless the
`\%V` atom is used in the pattern like: `:'<,'>s/\%Vfoo/bar/g`.


#### More info ####

* [Vim wiki](http://vim.wikia.com/wiki/Search_and_replace)


## Plugins ##


### [Buffer Bye](https://github.com/moll/vim-bbye) ###

Delete buffers (close files) without closing the window or bolloxing the layout.

Adds `:Bdelete` command. Use in preference to `:bdelete`.

Mapped to `<leader>C`


### [Command-T](https://github.com/wincent/Command-T) ###

Faster version of CtrlP.

| Mapping     | Action       |
| --          | --           |
| `<Leader>p` | Show files   |
| `<Leader>b` | Show buffers |
| `<Leader>r` | Show tags    |



### [GitGutter](https://github.com/airblade/vim-gitgutter) ###

Show git status in the gutter.

|       Command       |             Action             |
|---------------------|--------------------------------|
| `:GitGutterEnable`  |                                |
| `:GitGutterDisable` |                                |
| `:GitGutterToggle`  |                                |
| `]c`                | Jump to next hunk (change)     |
| `[c`                | Jump to previous hunk (change) |
| `<Leader>hs`        | Stage hunk                     |
| `<Leader>hr`        | Revert stage hunk              |
| `<Leader>hp`        | Preview hunk's changes         |


### [Indentpython.vim](https://github.com/vim-scripts/indentpython.vim) ###

Almost PEP-8-compliant indentation for Python.


### [Pymode](https://github.com/klen/python-mode) ###

Vim python-mode. PyLint, Rope, Pydoc, breakpoints from box.

|    Command    |                 Action                |
|---------------|---------------------------------------|
| `K`           | Show docs for current word/selections |
| `<Leader>r`   | Run current script                    |
| `<Leader>b`   | Set/unset breakpoint                  |
| `CTRL+c, g`   | Goto definition                       |
| `CTRL+c, d`   | Show documentation                    |
| `CTRL+c, f`   | Find occurences                       |
| `[[`          | Jump to previous class/function       |
| `]]`          | Jump to next class/function           |
| `[M`          | Jump to previous class or method      |
| `]M`          | Jump to next class or method          |
| `CTRL+c, r1r` | Rename module                         |
| `CTRL+c, ra`  | Insert import for current word        |
| `CTRL+c, ro`  | Organise imports                      |
| `CTRL+c, r1p` | Convert current module to package     |
| `CTRL+c, rm`  | Extract method from selected lines    |
| `CTRL+c, rl`  | Extract variable from selected lines  |
| `CTRL+c, ru`  | Use function                          |


### [SimpylFold](https://github.com/tmhedberg/SimpylFold) ###

Proper code folding for Python.

**If `foldmethod` or `foldexpr` are set elsewhere in `.vimrc`, the plugin
won't work!**


### [Tabular](https://github.com/godlygeek/tabular) ###

Line up text.

Use command `:Tabularize` to align lines.

| Command             | Action                            |
| --                  | --                                |
| `:Tabularize /,`    | Align selected lines on commas    |
| `:Tabularize /,/r0` | Right align with 0 padding spaces |

Align may be `l`, `r` or `c`, and padding may be any number.

See the [Vimcast](http://vimcasts.org/episodes/aligning-text-with-tabular-vim/).


### [Tagbar](http://majutsushi.github.io/tagbar/)

Browse tags in source code files.

* Toggle Tagbar: `<Leader>rt`, `F7` or `:TagbarToggle`
* Open and focus Tagbar: `:TagBarOpen`
* Open, focus and auto-close Tagbar: `:TagBarOpenAutoClose`


#### Tagbar commands

The following commands are valid in the Tagbar window:

* Display help: `F1`
* Jump to tag under cursor: `ENTER`
* Jump to tag under cursor, but stay in Taglist window: `p`
* Display prototype of tag on command line: `SPACE`
* Zoom window in/out to full frame: `x`
* Toggle fold: `o`
* Open/close fold: `+, zo`/`-, zc`
* Open/close *all* folds: `*, zR`/`=, zM`
* Close Tagbar window: `q`
* Next/previous root tag: `CTRL+n` and `CTRL+p`
* More help: `:help tagbar-intro`


### CtrlP

Works like CMD+T in TextMate/Sublime Text 2.

| Mapping               | Action                                           |
| --------------------  | ----------------------                           |
| `<Leader>p`           | Toggle `CtrlP`                                   |
| `F5`                  | Purge cache, reload settings                     |
| `CTRL+f` and `CTRL+b` | Cycle between modes                              |
| `CTRL+d`              | Filename only, not full path                     |
| `CTRL+r`              | Regex mode                                       |
| `CTRL+j` and `CTRL+k` | Navigate result list                             |
| `CTRL+p` and `CTRL+n` | Select previous/next string in prompt history    |
| `..`                  | Enter `..` or more dots to ascend directory tree |


### [MarkdownFootnotes](https://github.com/vim-scripts/MarkdownFootnotes) ###

Add markdown footnotes.

| Mapping     | Action               |
| --          | --                   |
| `<Leader>f` | Add footnote         |
| `<Leader>r` | Return from footnote |


### [NERDTree](https://github.com/scrooloose/nerdtree) ###

Navigate the filesystem in `vim`.

Toggle with `F3`.

| Mapping           | Action                                            |
| ----------------- | ------------------------------------------------- |
| `o`  and `<CR>`   | Open files, directories, bookmarks                |
| `go`              | Open selected file but leave cursor in NERDTree   |
| `t`               | Open selected file in new tab                     |
| `T`               | Same, but keep current tab selected               |
| `i`               | Open in split window                              |
| `gi`              | Same, but leave cursor in NERDTree                |
| `s`               | Open in new vsplit                                |
| `gs`              | Same, but leave cursor in NERDTree                |
| `O`               | Recursively open selected directory               |
| `x`               | Close the current node's parent                   |
| `X`               | Recursively close all children of current node    |
| `e`               | Edit the current directory                        |
| `D`               | Delete current bookmark                           |
| `P`               | Jump to root node                                 |
| `p`               | Jump to parent node                               |
| `C`               | Change tree root to selected directory            |
| `u`               | Move the tree root up one directory               |
| `r`               | Recursively refresh current directory             |
| `R`               | Recursively refresh root directory                |
| `m`               | Display NERDTree menu                             |
| `cd`              | Change `CWD` to selected directory                |
| `CD`              | Change tree root to `CWD`                         |
| `I`               | Toggle hidden files                               |
| `F`               | Toggle file display                               |
| `B`               | Toggle bookmark display                           |
| `q`               | Close NERDTree window                             |
| `A`               | Zoom (maximise/minimise) NERDTree window          |
| `?`               | Toggle quick help                                 |


### [NERDCommenter](https://github.com/scrooloose/nerdcommenter) ###

Comment/uncomment.

| Mapping                         | Action                                                                                                              |
| ------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `[count]<Leader>cc`             | Comment out current line or selected text in visual mode                                                            |
| `[count]<Leader>cn`             | Same as `<Leader>cc` but forces nesting                                                                             |
| `[count]<Leader>c<space>`       | Toggle comment state of selected line(s). If first selected line is commented, all selected lines are uncommented   |
| `[count]<Leader>cm`             | Comments given lines using one set of multipart delimiters                                                          |
| `[count]<Leader>ci`             | Toggle comment state of selected line(s) individually                                                               |
| `[count]<Leader>cs`             | Comment out selected lines "sexily"                                                                                 |
| `[count]<Leader>cy`             | Same as `<Leader>cc` except commented line(s) are yanked first                                                      |
| `<Leader>c$`                    | Comment current line from cursor to end of line                                                                     |
| `<Leader>cA`                    | Add comment delimiters to end of line and start insert mode between them                                            |
| `<Leader>ca`                    | Switches to alternative delimiters                                                                                  |
| `<Leader>cl` and `<Leader>cb`   | Same as `<Leader>cc` except delimiters are aligned on the left or both sides respectively                           |
| `<Leader>cu`                    | Uncomment the selected line(s)                                                                                      |
| `<Leader>c<space>`              | Toggle state of selected line(s)                                                                                    |


### [Pytest.vim](https://github.com/alfredodeza/pytest.vim) ###

Run `pytest` test cases from vim.

| Command                                        | Description             |
|--                                              | --                      |
| `:Pytest (class|method|function|file|project)` | Run tests for current X |
| `:Pytest (first|last|next|previous|end)`       | Navigate to errors      |


### [Sensible.vim](https://github.com/tpope/vim-sensible) ###

A collection of sensible defaults for vim. A step up from `nocompatible`.



### [Supertab](https://github.com/ervandew/supertab) ###

Make [YouCompleteMe](#youcompleteme) and [UltiSnips](#ultisnips) play nice.


### [Syntastic](https://github.com/scrooloose/syntastic)

Syntax checking for many filetypes.

By default, Syntastic checks syntax on save. To also check on open:

```vim
let g:syntastic_check_on_open=1
```

To alter mode settings:

```vim
let g:syntastic_mode_map = { 'mode': 'active',
                             'active_filetypes': ['ruby', 'php'],
                             'passive_filetypes': ['puppet'] }
```

| Mapping                | Action                                 |
| ---                    | ---                                    |
| `:SyntasticInfo`       | Show available checkers                |
| `:SyntasticCheck`      | Check current file                     |
| `:SyntasticToggleMode` | Toggle between active and passive mode |
| `:Errors`              | Show detected errors                   |


### [RagTag](https://github.com/tpope/vim-ragtag)

HTML generation. [TL;DR](#ragtag-tldr)


#### Head

| Mapping    | Action                   |
| --         | --                       |
| `CTRL+x !` | Insert DOCTYPE           |
| `CTRL+x #` | Insert meta content-type |
| `CTRL+x $` | Insert JavaScript        |
| `CTRL+x @` | Insert CSS stylesheet    |


#### Body

| Mapping        | Action                                             |
| --             | --                                                 |
| `CTRL+x SPACE` | Turn typed word into tag pair                      |
| `CTRL+x ENTER` | Turn typed word into tag pair with newline between |
| `CTRL+x /`     | Close last opened tag                              |


#### Script language tags

For the following mappings, suppose that you have typed "foo".

| Mapping    | Result           |
| --------   | -----            |
| `CTRL+x =` | `foo<%= \ | %>`  |
| `CTRL+x +` | `<%= foo\ | %>`  |
| `CTRL+x -` | `foo<% \  | %> ` |
| `CTRL+x _` | `<% foo\  | %> ` |
| `CTRL+x '` | `foo<%# \ | %>`  |
| `CTRL+x "` | `<%# foo\ | %>`  |


#### TL;DR [ragtag-tldr] ####

| Mapping        | Action                                                  |
| ------         | ------                                                  |
| `CTRL+x /`     | Close the last open HTML tag                            |
| `CTRL+x SPACE` | Create open/close HTML tags from the typed word         |
| `CTRL+x ENTER` | The same as CTRL+x SPACE but puts a newspace in between |
| `CTRL+x !`     | Insert HTML doctype                                     |
| `CTRL+x @`     | Insert CSS stylesheet                                   |
| `CTRL+x #`     | Insert meta content-type meta tag                       |
| `CTRL+x $`     | Load JavaScript document                                |


#### Links ####
* [catonmat.net tutorial series](http://www.catonmat.net/blog/vim-plugins-ragtag-allml-vim/)


### [Sparkup](https://github.com/rstacruz/sparkup) ###

Generate HTML from CSS selectors.

| Mapping | Action                           | Function name             |
| --      | --                               | --                        |
| `<c-e>` | Execute `sparkup`                | `g:sparkupExecuteMapping` |
| `<c-n>` | Jump to next empty tag/attribute | `g:sparkupNextMapping`    |
|         | Location of `sparkup` executable | `g:sparkup`               |


#### Examples

`div` expands to:

```html
<div></div>
```

`div#header` expands to:

```html
    <div id="header"></div>
```

`div.align-left#header` expands to:

```html
    <div id="header" class="align-left"></div>
```

`div#header + div#footer` expands to:

```html
    <div id="header"></div>
    <div id="footer"></div>
```

`#menu > ul` expands to:

```html
    <div id="menu">
        <ul></ul>
    </div>
```

`#menu > h3 + ul` expands to:

```html
    <div id="menu">
        <h3></h3>
        <ul></ul>
    </div>
```

`#header > h1{Welcome to our site}` expands to:

```html
    <div id="header">
        <h1>Welcome to our site</h1>
    </div>
```

`a[href=index.html]{Home}` expands to:

```html
    <a href="index.html">Home</a>
```

`ul > li*3` expands to:

```html
    <ul>
        <li></li>
        <li></li>
        <li></li>
    </ul>
```

`ul > li.item-$*3` expands to:

```html
    <ul>
        <li class="item-1"></li>
        <li class="item-2"></li>
        <li class="item-3"></li>
    </ul>
```

`ul > li.item-$*3 > strong` expands to:

```html
    <ul>
        <li class="item-1"><strong></strong></li>
        <li class="item-2"><strong></strong></li>
        <li class="item-3"><strong></strong></li>
    </ul>
```

`table > tr*2 > td.name + td*3` expands to:

```html
    <table>
        <tr>
            <td class="name"></td>
            <td></td>
            <td></td>
            <td></td>
        </tr>
        <tr>
            <td class="name"></td>
            <td></td>
            <td></td>
            <td></td>
        </tr>
    </table>
```

`#header > ul > li < p{Footer}` expands to:

```html
    <!-- The < symbol goes back up the parent; i.e., the opposite of >. -->
    <div id="header">
        <ul>
            <li></li>
        </ul>
        <p>Footer</p>
    </div>
```


### [Repeat](https://github.com/tpope/vim-repeat) ###

Repeat.vim remaps `.` in a way that plugins can tap into
it.

The following plugins support repeat.vim:

* [surround.vim](https://github.com/tpope/vim-surround)
* [speeddating.vim](https://github.com/tpope/vim-speeddating)
* [abolish.vim](https://github.com/tpope/vim-abolish)
* [unimpaired.vim](https://github.com/tpope/vim-unimpaired)
* [commentary.vim](https://github.com/tpope/vim-commentary)


### [BufferGator](https://github.com/jeetsukumaran/vim-buffergator) ###

List, navigate and select buffers.

* Open buffer list: `<Leader>b`
* Close buffer list: `<Leader>B` or `q`
* Open tab list: `<Leader>t`
* Close  tab list: `<Leader>T`
* Navigage up/down in lists: `k` or `CTRL+n` or `SPACE` / `j` or `CTRL-p` or `CTRL-SPACE`
* Open selected buffer in previous window: `o` or `ENTER`
* Open selected buffer in new tab: `t`
* Rebuild catalogue: `r`
* Wipe selected buffer: `x`
* Wipe selected buffer with prejudice: `X`


### [Surround.vim](https://github.com/vim-scripts/surround.vim)

Wrap text with quotation marks, brackets, tags etc. [TL;DR](#summary)

Surround.vim is all about "surroundings": parentheses, brackets, quotes, XML tags, and more.  The plugin provides mappings to easily delete, change and add such surroundings in pairs.

It's easiest to explain with examples.  Press `cs"'` inside

    "Hello world!"

to change it to

    'Hello world!'

Now press `cs'<q>` to change it to

    <q>Hello world!</q>

To go full circle, press `cst"` to get

    "Hello world!"

To remove the delimiters entirely, press `ds"`.

    Hello world!

Now with the cursor on "Hello", press `ysiw]` (`iw` is a text object).

    [Hello] world!

Let's make that braces and add some space (use `}` instead of `{` for no
space): `cs]{`

    { Hello } world!

Now wrap the entire line in parentheses with `yssb` or `yss)`.

    ({ Hello } world!)

Revert to the original text: `ds{ds)`

    Hello world!

Emphasize hello: `ysiw<em>`

    <em>Hello</em> world!

Finally, let's try out visual mode. Press a capital V (for linewise visual mode) followed by `S<p class="important">`.

    <p class="important">
      <em>Hello</em> world!
    </p>

This plugin is very powerful for HTML and XML editing, a niche which currently seems underfilled in Vim land.  (As opposed to HTML/XML *inserting*, for which many plugins are available).  Adding, changing, and removing pairs of tags simultaneously is a breeze.

The `.` command will work with `ds`, `cs`, and `yss` if you install
[repeat.vim](https://github.com/tpope/vim-repeat).


#### Summary ####


##### Normal mode

| Mapping | Action                                                                    |
| ------- | ------                                                                    |
| `ds`    | delete a surrounding                                                      |
| `cs`    | change a surrounding                                                      |
| `ys`    | add a surrounding                                                         |
| `yS`    | add a surrounding and place the surrounded text on a new line + indent it |
| `yss`   | add a surrounding to the whole line                                       |
| `ySs`   | add a surrounding to the whole line, place it on a new line + indent it   |
| `ySS`   | same as ySs                                                               |


##### Visual mode

| Mapping | Action                                                   |
| ------- | ------                                                   |
| `s`     | Add a surrounding                                        |
| `S`     | Add a surrounding but place text on new line + indent it |


##### Insert mode

| Mapping            | Action                                |
| -------            | ------------------------------------- |
| `<CTRL+s>`         | Add a surrounding                     |
| `<CTRL+s><CTRL+s>` | Add a new line + surrounding + indent |
| `<CTRL+g>s`        | Same as `<CTRL+s>`                    |
| `<CTRL+g>S`        | Same as `<CTRL+s><CTRL+s>`            |


### [Abolish.vim](https://github.com/tpope/vim-abolish)

Intelligent search and replace that can handle plurals and mixed case.

    :Abolish {despa,sepe}rat{e,es,ed,ing,ely,ion,ions,or}  {despe,sepa}rat{}

replaces:

    :iabbrev  seperation  separation
    :iabbrev desparation desperation
    :iabbrev  seperately  separately
    :iabbrev desparately desperately
    :iabbrev  Seperation  separation
    :iabbrev Desparation Desperation
    :iabbrev  Seperately  Separately
    :iabbrev Desparately Desperately
    :iabbrev  SEPERATION  SEPARATION
    :iabbrev DESPARATION DESPERATION
    :iabbrev  SEPERATELY  SEPARATELY
    :iabbrev DESPARATELY DESPERATELY

#### Substitution

Case- and plural-aware substitution.

    :%Subvert/facilit{y,ies}/building{,s}/g


### [Fugitive.vim](https://github.com/tpope/vim-fugitive/)

VIM `git` integration.

The following commands activate Fugitive. Most actions are performed in the Fugutive window after calling `:Gstatus`.

| Mapping      | Command     | Action     |
| -------      | -------     | ---------- |
| `<Leader>gs` | `:Gstatus`  | git status |
| `<Leader>gc` | `:Gcommit`  | git commit |
| `<Leader>gp` | `:Git push` | git push   |
| None         | `:Git ...`  | git ...    |

Command `:Git ...` is like calling `git` in a terminal, but it changes to the working directory first.

When the Fugitive status window is open, use the following keys:

| Mapping | Action                                     |
| ------  | -----------                                |
| `q`     | close status window                        |
| `R`     | reload status window                       |
| `-`     | add / reset (staged files)                 |
| `C`     | commit                                     |
| `q`     | close status window                        |
| `D`     | `:Gdiff` — diff a Fugitive revision        |
| `p`     | add --patch / reset --patch (staged files) |
| `o`     | `:Gsplit` — `:split` a Fugitive revision   |
| `S`     | `:Gvsplit` — `:vsplit` a Fugitive revision |

For more:

    :help fugitive


### [Tasklist](https://github.com/vim-scripts/TaskList.vim)

Show interactive list of TODO FIXME and similar tags.

| Mapping      | Command     | Action         |
| ------       | ----------- | ------         |
| `<leader>td` | `:TaskList` | Open task list |


### [Gundo](https://github.com/sjl/gundo.vim)

Navigate the `vim` undo tree. Toggle with `:GundoToggle`. I have mapped this to `<F5>`.

| Mapping | Action                                     |
| ------  | -----------                                |
| `p`     | Preview diff of current and selected state |
| `P`     | "Play" to the selected state               |
| `q`     | Close the undo graph                       |


### [Scriptease](https://github.com/tpope/vim-scriptease) ###

A script for making plugins. Contains convenience functions.

* `:PP`: Pretty print.  With no argument, acts as a REPL.
* `:Runtime`: Reload runtime files.  Like `:runtime!`, but it unlets any
  include guards first.
* `:Disarm`: Remove a runtime file's maps, commands, and autocommands,
  effectively disabling it.
* `:Scriptnames`: Load `:scriptnames` into the quickfix list.
* `:Verbose`: Capture the output of a `:verbose` invocation into the preview
  window.
* `:Time`: Measure how long a command takes.
* `:Breakadd`: Like its lowercase cousin, but makes it much easier to set
  breakpoints inside functions.  Also `:Breakdel`.
* `:Vedit`: Edit a file relative the runtime path. For example,
  `:Vedit plugin/scriptease.vim`. Also, `:Vsplit`, `:Vtabedit`, etc.
  Extracted from [pathogen.vim](https://github.com/tpope/vim-pathogen).
* `K`: Look up the `:help` for the VimL construct under the cursor.
* `zS`: Show the active syntax highlighting groups under the cursor.
* `g!`: Eval a motion or selection as VimL and replace it with the result.
  This is handy for doing math, even outside of VimL.  It's so handy, in fact,
  that it probably deserves its own plugin.
* Projections for
  [projectionist.vim](https://github.com/tpope/vim-projectionist).


### [UltiSnips](https://github.com/SirVer/ultisnips) ###

TextExpander-like snippets for `vim`.

| Key         | Action                                  |
| ----------- | --------------------------------------- |
| `<C-]>`     | Expand snippet                          |
| `<C-]>`     | Go to next placeholder in snippet       |
| `<C-[>`     | Go to previous placeholder in snippet   |


### [YouCompleteMe](https://github.com/Valloric/YouCompleteMe) ###

Superb autocompletion for `vim`.


### [TaskPaper](https://github.com/GrzegorzKozub/taskpaper.vim) ###

| Mapping      | Action                             |
| --           | --                                 |
| `<Leader>td` | Mark task done                     |
| `<Leader>tx` | Mark task cancelled                |
| `<Leader>tt` | Mark task as today                 |
| `<Leader>tD` | Archive done tasks                 |
| `<Leader>tX` | Show tasks marked as cancelled     |
| `<Leader>tT` | Show tasks marked as today         |
| `<Leader>t/` | Search for items including keyword |
| `<Leader>ts` | Search for items including tag     |
| `<Leader>tp` | Fold all projects                  |
| `<Leader>t.` | Fold all notes                     |
| `<Leader>tP` | Focus on current project           |
| `<Leader>tj` | Go to next project                 |
| `<Leader>tk` | Go to previous project             |
| `<Leader>tg` | Go to specified project            |
| `<Leader>tm` | Move task to specified project     |

### [Unimpaired](https://github.com/tpope/vim-unimpaired) ###

Pairs of complementary mappings, e.g. encode/decode.

| Mapping                   | Action                                        |
| --                        | --                                            |
| `[f` and `]f`             | next/previous file in directory               |
| `[n` and `]n`             | next/previous SCM conflict marker             |
| `[<space>` and `]<space>` | add newline before/after current line         |
| `[e` and `]e`             | Exchange current line with previous/next line |
| `[a` and `]a`             | `:previous` / `:next`                         |
| `[A` and `]A`             | `:first` / `:last`                            |
| `[b` and `]b`             | `:bprevious` / `:bnext`                       |
| `[B` and `]B`             | `:bfirst` / `:blast`                          |
| `[l` and `]l`             | `:lprevious` / `:lnext`                       |
| `[L` and `]L`             | `:lfirst` / `:llast`                          |
| `[q` and `]q`             | `:cprevious` / `:cnext`                       |
| `[Q` and `]Q`             | `:cfirst` / `:clast`                          |



Run `:help unimpaired` for more information.


### [Vim-go](https://github.com/fatih/vim-go) ###

Golang support for vim.

| Command         | Description                                                      |
| --------------- | ---------------------------------------------------------------- |
| `:GoDef`        | Go to definition                                                 |
| `:GoDoc`        | Look up documentation                                            |
| `:GoImport`     | Automatically import packages                                    |
| `:GoRun`        | Execute current file                                             |
| `:GoTest`       | Run unit tests and see results                                   |
| `:GoPath`       | View or change `GOPATH`                                          |
| `:GoCoverage`   | Create coverage profile                                          |
| `:GoMetaLinter` | Run all linters                                                  |
| `:GoLint`       | Lint the code                                                    |
| `:GoVet`        | Catch static errors                                              |
| `:GoImplements` | Analyse source code                                              |
| `:GoCallees`    | Analyse source code                                              |
| `:GoReferrers`  | Analyse source code                                              |
| `:GoRename`     | Rename identifier                                                |
| `:GoErrCheck`   | Check unchecked errors                                           |
| `:GoPlay`       | Share current code on [play.golang.org](https://play.golang.org) |


#### My mappings ####

| Mapping      | Command                      |
| --           | --                           |
| `<F5>`       | `:GoRun`                     |
| `<C-b>`      | `:GoBuild`                   |
| `K`          | `:GoDoc`                     |
| `<Leader>gd` | `:GoDoc`                     |
| `<Leader>gb` | `:GoDoc` in browser          |
| `<Leader>gv` | `:GoDoc` vertical            |
| `<Leader>gi` | `:GoImplements`              |
| `<Leader>ds` | `:GoDef` in a split          |
| `<Leader>dv` | `:GoDef` in a vertical split |

Run `help vim-go` for more information.


### [Vim-snippets](https://github.com/honza/vim-snippets) ###

A collection of snippets for UltiSnips and snipMate.


### Installing Plugins ###

Use [vim-plug](https://github.com/junegunn/vim-plug) to install plugins
in `~/.vim/plugged`.

Plugins are configured in `~/.vim/vimrc.plugins`.
