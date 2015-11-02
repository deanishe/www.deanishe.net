title: Bash Example
date: 2015-11-01
author: deanishe
email: deanishe@deanishe.net
tags: bash
---

An example of a big, fat bash script posted inline.

```sh
# vim: set filetype=sh:

export LANG=en_GB.UTF-8
export LC_ALL=en_GB.UTF-8

#-----------------------------------------------------------------------------
# Source other configuration files
#-----------------------------------------------------------------------------

PLATFORM=$(una
HOSTNAME=$(hostname -s | tr '[:upper:]' '[:lower:]')

if [[ $PLATFORM == 'Darwin' ]]; then
  PLATFORM=osx
elif [[ $PLATFORM == 'Linux' ]]; then
    PLATFORM=linux
fi

# Get all source files applicable to this machine/platform, sort them
# by filename, then source them
sourcefiles=()
# sourceroot=~/.dotfiles/source.bash
sourceroot=~/.dotfiles/lib/source.d

sourcedirs=()
sourcedirs+=("${sourceroot}")
sourcedirs+=("${sourceroot}/${PLATFORM}")
sourcedirs+=("${sourceroot}/${HOSTNAME}")

# Append suffixes to names to sort *.sh files before *.bash
function makeSortName() {
  local fn="$1"
  local p="$2"  # Priority
  local bn
  local x
  local sn
  local sfx
  x="${fn##*.}"
  bn="${fn%.*}"
  if [[ "$x" = "sh" ]]; then
    sfx="-100"
  elif [[ "$x" = "bash" ]]; then
    sfx="-200"
  else
    return
  fi
  echo "${bn}${sfx}${p}"
}


# Build array of 'filename   filepath' entries
i=1
for dirpath in "${sourcedirs[@]}"; do
  # echo $dirpath
  if [[ -d "${dirpath}" ]]; then
    for filename in $(ls "${dirpath}"); do
      # Only load .sh and .bash files. Give .sh files a higher priority, so
      # a bash-specific .bash one of the same name can override it.
      name="$( makeSortName "${filename}" "-$i" )"
      if [[ -z "$name" ]]; then
        continue
      fi
      sourcefiles+=("${name} ${dirpath}/${filename}")
    done
  fi
  i=`expr $i + 1`
done

# echo "${#sourcefiles[@]} sourcefiles"
# printf '%s\n' "${sourcefiles[@]}"

# Sort the list then remove filenames
sourcefiles=( $(
for entry in "${sourcefiles[@]}"; do
    echo "${entry}"
done | sort | awk '{print $2}' ))

# Source the files
for filepath in "${sourcefiles[@]}"; do
  # echo "sourcing ${filepath} ..."
  source "${filepath}"
done

unset filename i name sourcefiles sourceroot sourcedirs
unset PLATFORM HOSTNAME

# If not running interactively, don't do anything else
[ -z "$PS1" ] && return

#-----------------------------------------------------------------------------
# Terminal/shell settings
#-----------------------------------------------------------------------------

# check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize


#-----------------------------------------------------------------------------
# History
#-----------------------------------------------------------------------------

# Store 10000 commands in bash history
export HISTFILESIZE=10000
export HISTSIZE=10000
# don't overwrite GNU Midnight Commander's setting of `ignorespace'.
HISTCONTROL=$HISTCONTROL${HISTCONTROL+,}ignoredups
# ... or force ignoredups and ignorespace
HISTCONTROL=ignoreboth

# append to the history file, don't overwrite it
shopt -s histappend


#-----------------------------------------------------------------------------
# Program setup
#-----------------------------------------------------------------------------

#
# Setup Grep
#
export GREP_OPTIONS=

# Ignore certain directory patterns
export GREP_OPTIONS="--exclude-dir=.svn $GREP_OPTIONS"
export GREP_OPTIONS="--exclude-dir=.git $GREP_OPTIONS"
export GREP_OPTIONS="--exclude-dir=CVS $GREP_OPTIONS"
export GREP_OPTIONS="--exclude=\*.svn\* $GREP_OPTIONS"
export GREP_OPTIONS="--exclude=\*.git\* $GREP_OPTIONS"



[ -f ~/.fzf.bash ] && source ~/.fzf.bash

```