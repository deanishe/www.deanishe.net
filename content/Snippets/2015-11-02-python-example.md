title: Python Example
date: 2015-11-02
author: deanishe
email: deanishe@deanishe.net
tags: python, tagging
---

```python
#!/usr/bin/env python
# encoding: utf-8
# vim: ft=python
#
# tag.py
#
# Copyright © 2015 Dean Jackson <deanishe@deanishe.net>
#
# MIT Licence. See http://opensource.org/licenses/MIT
#
# Created on 2015-10-07

# ---------------------------------------------------------
# dP                dP
# 88                88
# 88d888b. .d8888b. 88 88d888b.
# 88'  `88 88ooood8 88 88'  `88
# 88    88 88.  ... 88 88.  .88
# dP    dP `88888P' dP 88Y888P'
#                      88
#                      dP
# ---------------------------------------------------------
"""tags [options] [<file>...]

List, edit and filter by OS X filesystem tags.

If no files are specified as arguments, filepaths will be read
from STDIN.


Listing tags:

    If no tags are specified with -a/-r/-t, the current tags will
    be printed to STDOUT.

    Each file will be printed on a separate line followed by a tab
    and its tags (also tab-separated).


Adding/removing tags:

    Use -a/-r/-t to add, remove and toggle tags respectively.
    Toggle means that if a tag is present, it will be removed.
    If it is absent, it will be added.

    Note: To add/remove/toggle multiple tags, specify multiple
    -a/-r/-t options, e.g.:

        tags -a important -a todo some/file/path

    If you use -n/--nothing to test the script (tags will not be
    altered), the files will be printed to STDOUT with their tags
    as when listing tags, but tags to be added will be preceded
    by '+', tags to be removed by '-' and unchanged tags as-is.


Filtering by tag(s):

    Use -f/--filter to filter input files by tags. Specify tags that
    must be included with -i/--include and tags that must be excluded
    with -x/--exclude. Use -I/--invert to print non-matching files.

    If no files match, the script will exit with status 1, so you
    can use it to test for the presence/absence of tags.


Usage:
    tags [-v|-q|-d] [-n] [-a <tag>]... [-r <tag>]... [-t <tag>]... [<file>...]
    tags [-v|-q|-d] [-n] -p <tag1> <tag2> [<file>...]
    tags [-v|-q|-d] -f [-I] [-i <tag>]... [-x <tag>]... [<file>...]
    tags (-h|-V)


Options:
    -a, --add=<tag>           Tag to add.
    -r, --remove=<tag>        Tag to remove.
    -t, --toggle=<tag>        Remove tag if present, else add it.
    -p, --pair <tag1> <tag2>  Toggle tags as a pair. If <tag1> is lacking,
                              it is added and <tag2> removed. If <tag1>
                              is present, <tag1> is removed and <tag2>
                              added.
    -n, --nothing             Don't actually write tags.
    -f, --filter              Run in filter mode.
    -i, --include=<tag>       When filtering, match files containing tag.
    -x, --exclude=<tag>       When filtering, match files without tag.
    -I, --invert              When filtering, invert results (i.e. print
                              results that *don't* match).
    --version                 Show version number and exit.
    -h, --help                Show this message and exit.
    -q, --quiet               Only show errors.
    -v, --verbose             Show info messages.
    -d, --debug               Show debugging message.

"""

# ---------------------------------------------------------
# oo                                         dP
#                                            88
# dP 88d8b.d8b. 88d888b. .d8888b. 88d888b. d8888P .d8888b.
# 88 88'`88'`88 88'  `88 88'  `88 88'  `88   88   Y8ooooo.
# 88 88  88  88 88.  .88 88.  .88 88         88         88
# dP dP  dP  dP 88Y888P' `88888P' dP         dP   `88888P'
#               88
#               dP
# ---------------------------------------------------------

from __future__ import print_function, unicode_literals, absolute_import


import logging
import logging.handlers
import os
import sys

from docopt import docopt
# This script only runs on OS X and requires the PyObjC library
import Foundation


__version__ = '1.0'
__author__ = 'deanishe@deanishe.net'


NSURLTagNamesKey = 'NSURLTagNamesKey'


# ---------------------------------------------------------
# dP                            oo
# 88
# 88 .d8888b. .d8888b. .d8888b. dP 88d888b. .d8888b.
# 88 88'  `88 88'  `88 88'  `88 88 88'  `88 88'  `88
# 88 88.  .88 88.  .88 88.  .88 88 88    88 88.  .88
# dP `88888P' `8888P88 `8888P88 dP dP    dP `8888P88
#                  .88      .88                  .88
#              d8888P   d8888P               d8888P
# ---------------------------------------------------------

DEFAULT_LOG_LEVEL = logging.WARNING
LOGPATH = os.path.expanduser('~/Library/Logs/MyScripts.log')
LOGSIZE = 1024 * 1024 * 1  # 1 megabyte

# Configured by init_logging()
log = logging.getLogger('tag')
log.addHandler(logging.NullHandler())


class TechnicolorFormatter(logging.Formatter):
    """Intelligent and pretty log formatting.

    Colourise output to a TTY and prepend logging level name to
    levels other than INFO.

    """

    BLACK, RED, GREEN, YELLOW, BLUE, MAGENTA, CYAN, WHITE = range(8)

    RESET = '\033[0m'
    COLOUR_BASE = '\033[1;{:d}m'
    BOLD = '\033[1m'

    LEVEL_COLOURS = {
        logging.DEBUG:    BLUE,
        logging.INFO:     WHITE,
        logging.WARNING:  YELLOW,
        logging.ERROR:    MAGENTA,
        logging.CRITICAL: RED,
    }

    def __init__(self, fmt=None, datefmt=None, technicolor=True):
        """Create new Formatter.

        Args:
            fmt (str): A `logging.Formatter` format string.
            datefmt (str): `strftime` format string.
            technicolor (bool): Colourise TTY output?

        """

        logging.Formatter.__init__(self, fmt, datefmt)
        self.technicolor = technicolor
        self._isatty = sys.stderr.isatty()

    def format(self, record):
        """Extend `logging.Formatter.format()`.

        Prepend log level for levels other than INFO.
        Colourise level names for TTY output.

        """

        # Output `INFO` messages without level name.
        # The idea is to treat them as normal status messages.
        if record.levelno == logging.INFO:
            msg = logging.Formatter.format(self, record)
            return msg

        # Other levels have their level name colourised if
        # the destination is a TTY.
        if self.technicolor and self._isatty:
            colour = self.LEVEL_COLOURS[record.levelno]
            bold = (False, True)[record.levelno > logging.INFO]
            levelname = self.colourise('{:9s}'.format(record.levelname),
                                       colour, bold)
        else:
            levelname = '{:9s}'.format(record.levelname)

        return (levelname + logging.Formatter.format(self, record))

    def colourise(self, text, colour, bold=False):
        """Surround `text` with terminal colours."""

        colour = self.COLOUR_BASE.format(colour + 30)
        output = []
        if bold:
            output.append(self.BOLD)
        output.append(colour)
        output.append(text)
        output.append(self.RESET)
        return ''.join(output)


def init_logging():
    """Set up logging handlers, add and configure global `log`"""

    # logfile
    logfile = logging.handlers.RotatingFileHandler(LOGPATH,
                                                   maxBytes=LOGSIZE,
                                                   backupCount=1)
    formatter = logging.Formatter(
        '%(asctime)s %(levelname)-8s [%(name)-12s] %(message)s',
        datefmt="%d/%m %H:%M:%S")
    logfile.setFormatter(formatter)
    logfile.setLevel(logging.DEBUG)

    # console output
    console = logging.StreamHandler()
    formatter = TechnicolorFormatter('%(message)s')
    console.setFormatter(formatter)
    console.setLevel(logging.DEBUG)

    log.addHandler(logfile)
    log.addHandler(console)


# ---------------------------------------------------------
#   dP
#   88
# d8888P .d8888b. .d8888b. .d8888b.
#   88   88'  `88 88'  `88 Y8ooooo.
#   88   88.  .88 88.  .88       88
#   dP   `88888P8 `8888P88 `88888P'
#                      .88
#                  d8888P
# ---------------------------------------------------------

def get_tags(filepath):
    """Return `set` of tags for `filepath`.

    Args:
        filepath (unicode): File/directory whose tags to retrieve.

    """

    url = Foundation.NSURL.fileURLWithPath_(filepath)
    metadata, error = url.resourceValuesForKeys_error_([NSURLTagNamesKey],
                                                       None)
    if not metadata:
        return set()
    if NSURLTagNamesKey not in metadata:
        return set()

    return set(metadata[NSURLTagNamesKey])


def set_tags(filepath, tags):
    """Set tags for `filepath`.

    Args:
        filepath (unicode): File/directory whose tags to set.
        tags (sequence): Tags to set for `filepath`.

    Raises:
        OSError: Raised if call to Foundation API to set tags fails.

    """

    tags = sorted(set(tags))

    url = Foundation.NSURL.fileURLWithPath_(filepath)
    result, error = url.setResourceValue_forKey_error_(tags, NSURLTagNamesKey,
                                                       None)
    if not result:
        raise OSError('Could not set tags',
                      unicode(error).encode('ascii', 'ignore'))


def add_tags(filepath, tags):
    """Add `tags` to existing `filepath` tags.

    Args:
        filepath (unicode): File/directory to add tags to.
        tags (sequence): Tags to add to `filepath`.

    Raises:
        OSError: Raised if call to Foundation API to set tags fails.

    """

    tags = set(tags)

    current_tags = get_tags(filepath)
    new_tags = tags - current_tags

    if len(new_tags):
        log.debug('Adding tags %r to %s', new_tags, filepath)
        set_tags(filepath, current_tags | new_tags)


def remove_tags(filepath, tags):
    """Remove `tags` from current `filepath` tags.

    Args:
        filepath (unicode): File/directory to remove tags from.
        tags (sequence): Tags to remove.

    Raises:
        OSError: Raised if call to Foundation API to set tags fails.

    """

    tags = set(tags)

    current_tags = get_tags(filepath)
    new_tags = current_tags - tags

    if new_tags:
        set_tags(filepath, new_tags | current_tags)


# ---------------------------------------------------------
# dP                dP
# 88                88
# 88d888b. .d8888b. 88 88d888b. .d8888b. 88d888b. .d8888b.
# 88'  `88 88ooood8 88 88'  `88 88ooood8 88'  `88 Y8ooooo.
# 88    88 88.  ... 88 88.  .88 88.  ... 88             88
# dP    dP `88888P' dP 88Y888P' `88888P' dP       `88888P'
#                      88
#                      dP
# ---------------------------------------------------------

def read_stdin():
    """Read filepaths from STDIN.

    Returns:
        filepaths (list): List of `unicode` filepaths.

    """

    paths = []
    for line in sys.stdin.readlines():
        line = line.strip()
        if not line:
            continue
        paths.append(os.path.abspath(line.decode('utf-8')))

    return paths


def format_tags(filepath, tags):
    """Return tab-separated string list of `filepath` and `tags`.

    Args:
        filepath (unicode): Filepath
        tags (list): Tags for `filepath`

    Returns:
        UTF-8 encoded string.

    """

    return '\t'.join([filepath] + tags).encode('utf-8')


def print_tags(filepath, tags):
    """Print `filepath` and `tags`.

    Args:
        filepath (unicode): Filepath
        tags (list): Tags for `filepath`.

    """

    print(format_tags(filepath, tags))


def filter_files(filepaths, include_tags, exclude_tags, invert=False):
    """Yield items in `filepaths` matching tags.

    Args:
        filepaths (list): Filepaths to filter.
        include_tags (set): Tags that must be present.
        exclude_tags (set): Tags that must not be present.
        invert (bool, optional): If True, yield results that don't match.

    Yields:
        filepaths (unicode)

    """

    for filepath in filepaths:

        tags = get_tags(filepath)

        log.debug('filepath : %s  tags: %s include: %s  exclude: %s',
                  filepath, tags, include_tags, exclude_tags)

        if tags >= include_tags and not tags & exclude_tags:
            if not invert:
                yield filepath

        elif invert:
            yield filepath


# ---------------------------------------------------------
#                      dP   oo
#                     88
# .d8888b. .d8888b. d8888P dP .d8888b. 88d888b. .d8888b.
# 88'  `88 88'  `""   88   88 88'  `88 88'  `88 Y8ooooo.
# 88.  .88 88.  ...   88   88 88.  .88 88    88       88
# `88888P8 `88888P'   dP   dP `88888P' dP    dP `88888P'
# ---------------------------------------------------------

def do_list_files(args):
    """Print list of filepaths and their tags to STDOUT.

    Args:
        args (dict): Parsed (and unicodified) `docopt` args.

    """

    for filepath in args['<file>']:
        if not os.path.exists(filepath):
            log.error('File does not exists: %s', filepath)
            continue

        tags = sorted(get_tags(filepath))
        print_tags(filepath, tags)

    return 0


def do_filter_files(args):
    """Print list of (non-)matching filepaths to STDOUT.

    Args:
        args (dict): Parsed (and unicodified) `docopt` args.

    """

    to_include = set(args.get('--include'))
    to_exclude = set(args.get('--exclude'))

    if not len(to_include | to_exclude):
        log.critical('No tags to include or exclude specified.')
        return 1

    if not len(args.get('<file>')):
        log.critical('No filepaths specified.')
        return 1

    common = to_include & to_exclude

    if len(common):
        log.critical(
            '%s specified as include *and* exclude.', ', '.join(
                ['`{}`'.format(t) for t in common]))
        return 1

    count = 0
    for filepath in filter_files(args.get('<file>'),
                                 to_include, to_exclude,
                                 args.get('--invert')):

        print(filepath.encode('utf-8'))
        count += 1

    if not count:
        return 1

    return 0


def do_tag_files(args):
    """Add/remove tags from files.

    Args:
        args (dict): Parsed (and unicodified) `docopt` args.

    """

    to_add = set(args.get('--add'))
    to_remove = set(args.get('--remove'))
    to_toggle = set(args.get('--toggle'))

    for filepath in args['<file>']:
        current = get_tags(filepath)

        for tag in to_toggle:
            if tag in current:
                to_remove.add(tag)
            else:
                to_add.add(tag)

        updated = (current - to_remove) | to_add

        added = updated - current
        removed = current - updated
        unchanged = current & updated

        if current != updated:
            tag_status = ([(t, '+') for t in added] +
                          [(t, '-') for t in removed] +
                          [(t, '') for t in unchanged])
            tag_status.sort()
            tag_status = ['{}{}'.format(t[1], t[0]) for t in tag_status]

            if args.get('--nothing'):
                print_tags(filepath, tag_status)
            else:
                log.info(format_tags(filepath, tag_status))
                set_tags(filepath, list(updated))

        elif args.get('--nothing'):
            print_tags(filepath, sorted(current))

    return 0


def do_toggle_pair(args):
    """Toggle pair of tags on/off.

    Args:
        args (dict): Parsed (and unicodified) `docopt` args.

    Returns:
        int: Exit status.
    """

    tag1 = args.get('--pair')
    tag2 = args.get('<tag2>')

    for filepath in args['<file>']:
        current = get_tags(filepath)
        updated = set(current)
        if tag1 not in current:
            updated.add(tag1)
            updated.discard(tag2)
            # log.debug('Added `%s`, removed `%s` : %r', tag1, tag2, filepath)
        else:
            updated.add(tag2)
            updated.discard(tag1)
            # log.debug('Added `%s`, removed `%s` : %r', tag2, tag1, filepath)

        if updated != current:
            tag_status = ([(t, '+') for t in updated.difference(current)] +
                          [(t, '-') for t in current.difference(updated)] +
                          [(t, '') for t in current.intersection(updated)])
            tag_status.sort()
            tag_status = ['{}{}'.format(t[1], t[0]) for t in tag_status]

            if args.get('--nothing'):
                print_tags(filepath, tag_status)

            else:
                log.info(format_tags(filepath, tag_status))
                set_tags(filepath, list(updated))


def main():
    """Run command-line script."""

    init_logging()

    args = docopt(__doc__, version=__version__)

    if args.get('--verbose'):
        log.setLevel(logging.INFO)
    elif args.get('--quiet'):
        log.setLevel(logging.ERROR)
    elif args.get('--debug'):
        log.setLevel(logging.DEBUG)
    else:
        log.setLevel(DEFAULT_LOG_LEVEL)
    log.debug("Set log level to %s" %
              logging.getLevelName(log.level))

    # Decode args
    for k, v in args.items():
        if isinstance(v, list):
            args[k] = [s.decode('utf-8') for s in v]

    log.debug('args : %s', args)

    # If no tags have been specified, just list file tags
    list_mode = not len(args.get('--add') +
                        args.get('--remove') +
                        args.get('--toggle'))

    if not args.get('<file>'):
        args['<file>'] = read_stdin()

    if not args['<file>']:
        log.error('No input files specified.')
        return 0

    if args.get('--filter'):
        return do_filter_files(args)

    elif args.get('--pair'):
        return do_toggle_pair(args)

    elif list_mode:
        return do_list_files(args)

    else:
        return do_tag_files(args)

if __name__ == '__main__':
    sys.exit(main())
```