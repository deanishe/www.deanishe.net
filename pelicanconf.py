#!/usr/bin/env python
# -*- coding: utf-8 -*- #
from datetime import datetime

AUTHOR = 'deanishe'
SITENAME = 'www.deanishe.net'
SITEURL = ''

DEFAULT_PAGINATION = 10

TIMEZONE = 'Europe/Berlin'

DEFAULT_LANG = 'en'
LOCALE = 'en_GB.UTF-8'
DATE_FORMATS = {
    'en': ('en_GB.UTF-8', '%d/%m/%Y'),
}

CURRENT_YEAR = datetime.now().year

# Tell Pelican not to parse .html files
READERS = {'html': None}

SUMMARY_MAX_LENGTH = 100


#                     dP   dP
#                     88   88
# 88d888b. .d8888b. d8888P 88d888b. .d8888b.
# 88'  `88 88'  `88   88   88'  `88 Y8ooooo.
# 88.  .88 88.  .88   88   88    88       88
# 88Y888P' `88888P8   dP   dP    dP `88888P'
# 88
# dP

DIRECT_TEMPLATES = [
    'index',
    'posts',
    'categories',
    'authors',
    'archives',
    'tags',
]
PAGINATED_DIRECT_TEMPLATES = ['posts', 'index']

PATH = 'content'
PAGE_PATHS = ['pages']
ARTICLE_EXCLUDES = ['pages']

USE_FOLDER_AS_CATEGORY = True

#                   dP
#                   88
# dP    dP 88d888b. 88 .d8888b.
# 88    88 88'  `88 88 Y8ooooo.
# 88.  .88 88       88       88
# `88888P' dP       dP `88888P'

# News & stats
# _ARTICLE_BASE = '{category}/{date:%Y}/{date:%m}/{date:%d}/{slug}/'
# ARTICLE_URL = '/' + _ARTICLE_BASE
# ARTICLE_URL = '{category}/{date:%Y}/{date:%m}/{date:%d}/{slug}/'
# ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}/'
# ARTICLE_SAVE_AS = ARTICLE_URL + 'index.html'

ARTICLE_URL = 'p/{date:%Y}/{date:%m}/{date:%d}/{slug}/'
ARTICLE_SAVE_AS = ARTICLE_URL + 'index.html'

# Author(s)
AUTHOR_URL = 'n/{slug}/'
AUTHOR_SAVE_AS = AUTHOR_URL + 'index.html'
AUTHORS_SAVE_AS = 'n/index.html'

# Categories (i.e. Neuigkeiten)
# _CATEGORY_BASE = '{slug}/'
# CATEGORY_URL = '/' + _CATEGORY_BASE
CATEGORY_URL = 'c/{slug}/'
CATEGORY_SAVE_AS = CATEGORY_URL + 'index.html'
CATEGORIES_SAVE_AS = 'c/index.html'

# Drafts
DRAFT_URL = 'd/{date:%Y%m%d}/{slug}.html'
DRAFT_SAVE_AS = DRAFT_URL

# List of all pages
INDEX_SAVE_AS = 'index.html'

# Static pages
PAGE_URL = 's/{slug}/'
PAGE_SAVE_AS = PAGE_URL + 'index.html'

# Tags
TAG_URL = 't/{slug}/'
TAG_SAVE_AS = TAG_URL + 'index.html'
TAGS_SAVE_AS = 't/index.html'

# Archives
YEAR_ARCHIVE_SAVE_AS = 'p/{date:%Y}/index.html'
MONTH_ARCHIVE_SAVE_AS = 'p/{date:%Y}/{date:%m}/index.html'
DAY_ARCHIVE_SAVE_AS = 'p/{date:%Y}/{date:%m}/{date:%d}/index.html'
ARCHIVES_SAVE_AS = 'p/index.html'


# .8888b                         dP
# 88   "                         88
# 88aaa  .d8888b. .d8888b. .d888b88 .d8888b.
# 88     88ooood8 88ooood8 88'  `88 Y8ooooo.
# 88     88.  ... 88.  ... 88.  .88       88
# dP     `88888P' `88888P' `88888P8 `88888P'

# Feed generation is usually not desired when developing
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None


#            dP              dP   oo
#            88              88
# .d8888b. d8888P .d8888b. d8888P dP .d8888b.
# Y8ooooo.   88   88'  `88   88   88 88'  `""
#       88   88   88.  .88   88   88 88.  ...
# `88888P'   dP   `88888P8   dP   dP `88888P'

STATIC_PATHS = [
    'images',
    'files',
    'extra',
]

EXTRA_PATH_METADATA = {}


#   dP   dP
#   88   88
# d8888P 88d888b. .d8888b. 88d8b.d8b. .d8888b.
#   88   88'  `88 88ooood8 88'`88'`88 88ooood8
#   88   88    88 88.  ... 88  88  88 88.  ...
#   dP   dP    dP `88888P' dP  dP  dP `88888P'

THEME = 'themes/default'
# THEME = 'simple'
ASSET_SOURCE_PATHS = ['static']
SHOW_AUTHOR = True

MAIN_MENU = [
    # ('Home', '/'),
    # ('Alfred Workflows', '/alfred-workflows/'),
]


#          dP                   oo
#          88
# 88d888b. 88 dP    dP .d8888b. dP 88d888b. .d8888b.
# 88'  `88 88 88    88 88'  `88 88 88'  `88 Y8ooooo.
# 88.  .88 88 88.  .88 88.  .88 88 88    88       88
# 88Y888P' dP `88888P' `8888P88 dP dP    dP `88888P'
# 88                        .88
# dP                    d8888P

PLUGIN_PATHS = ['pelican-plugins']
PLUGINS = [
    'assets',
    'sitemap',
    'extract_toc',
]

JINJA_EXTENSIONS = ['jinja2.ext.loopcontrols']

# Markdown
MD_EXTENSIONS = [
    'codehilite(css_class=highlight, linenums=True)',
    'extra',
    'smarty',
    'meta',
    # 'headerid',
    'toc',
]


#          oo   dP
#               88
# .d8888b. dP d8888P .d8888b. 88d8b.d8b. .d8888b. 88d888b.
# Y8ooooo. 88   88   88ooood8 88'`88'`88 88'  `88 88'  `88
#       88 88   88   88.  ... 88  88  88 88.  .88 88.  .88
# `88888P' dP   dP   `88888P' dP  dP  dP `88888P8 88Y888P'
#                                                 88
#                                                 dP

# https://github.com/getpelican/pelican-plugins/tree/master/sitemap
SITEMAP = {
    'format': 'xml',
    'priorities': {
        'articles': 0.5,
        'indexes': 0.5,
        'pages': 0.5,
    },
    'changefreqs': {
        'articles': 'weekly',
        'indexes': 'weekly',
        'pages': 'monthly',
    }
}


#       dP
#       88
# .d888b88 .d8888b. dP   .dP
# 88'  `88 88ooood8 88   d8'
# 88.  .88 88.  ... 88 .88'
# `88888P8 `88888P' 8888P'

LOAD_CONTENT_CACHE = False
DEVMODE = True
RELATIVE_URLS = True
GENTIME = datetime.now()

# Blogroll
LINKS = (('Pelican', 'http://getpelican.com/'),
         ('Python.org', 'http://python.org/'),
         ('Jinja2', 'http://jinja.pocoo.org/'),
         ('You can modify those links in your config file', '#'),)

# Social widget
SOCIAL = (('You can add links in your config file', '#'),
          ('Another social link', '#'),)


# Uncomment following line if you want document-relative URLs when developing
#RELATIVE_URLS = True
