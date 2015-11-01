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


# ---------------------------------------------------------
# Paths and URL formatting
# ---------------------------------------------------------

DIRECT_TEMPLATES = [
    'posts',
    'categories',
    'authors',
    'archives',
    'tags',
]
PAGINATED_DIRECT_TEMPLATES = ['posts']

PATH = 'content'
PAGE_PATHS = ['pages']
ARTICLE_EXCLUDES = ['pages']

USE_FOLDER_AS_CATEGORY = True

# News & stats
# _ARTICLE_BASE = '{category}/{date:%Y}/{date:%m}/{date:%d}/{slug}/'
# ARTICLE_URL = '/' + _ARTICLE_BASE
# ARTICLE_URL = '{category}/{date:%Y}/{date:%m}/{date:%d}/{slug}/'
# ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}/'
# ARTICLE_SAVE_AS = ARTICLE_URL + 'index.html'

ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}.html'
ARTICLE_SAVE_AS = ARTICLE_URL

# Static pages
PAGE_URL = '{slug}/'
PAGE_SAVE_AS = '{slug}/index.html'

# Categories (i.e. Neuigkeiten)
# _CATEGORY_BASE = '{slug}/'
# CATEGORY_URL = '/' + _CATEGORY_BASE
CATEGORY_URL = '{slug}/'
CATEGORY_SAVE_AS = CATEGORY_URL + 'index.html'


# ---------------------------------------------------------
# Feeds
# ---------------------------------------------------------

# Feed generation is usually not desired when developing
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None


# ---------------------------------------------------------
# Static files
# ---------------------------------------------------------

STATIC_PATHS = [
    'images',
    'files',
    'extra',
]

EXTRA_PATH_METADATA = {}


# ---------------------------------------------------------
# Theme
# ---------------------------------------------------------

THEME = 'themes/default'
# THEME = 'simple'
ASSET_SOURCE_PATHS = ['static']
SHOW_AUTHOR = False


# ---------------------------------------------------------
# Plugins
# ---------------------------------------------------------

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

# ---------------------------------------------------------
# Sitemap
# ---------------------------------------------------------

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

# ---------------------------------------------------------
# Development only
# ---------------------------------------------------------

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
