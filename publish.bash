#!/bin/bash

PELICAN=$HOME/.pyenv/versions/www.deanishe.net/bin/pelican
ROOT=$HOME/Sites/www.deanishe.net

CONTENTDIR=$ROOT/content
OUTPUTDIR=$ROOT/output
SETTINGS=$ROOT/pelicanconf.py

LOGFILE=$ROOT/publish.log
# PIDFILE=$ROOT/publish.pid


PELICANOPTS=--debug

PELICAN "$CONTENTDIR" -o "$OUTPUTDIR" -s "$SETTINGS" $PELICANOPTS 2>&1 | tee -a "$LOGFILE"

exit ${PIPESTATUS[0]}
