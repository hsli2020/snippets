#! /bin/bash

COLOR_RED="\033[0;31m"
COLOR_YELLOW="\033[0;33m"
COLOR_GREEN="\033[0;32m"
COLOR_OCHRE="\033[38;5;95m"
COLOR_BLUE="\033[0;34m"
COLOR_WHITE="\033[0;37m"
COLOR_RESET="\033[0m"


function git_color {
  local git_status="$(git status 2> /dev/null)"

  if [[ ! $git_status =~ "working directory clean" ]]; then
    echo -e $COLOR_RED
  elif [[ $git_status =~ "Your branch is ahead of" ]]; then
    echo -e $COLOR_YELLOW
  elif [[ $git_status =~ "nothing to commit" ]]; then
    echo -e $COLOR_GREEN
  else
    echo -e $COLOR_OCHRE
  fi
}


function git_branch {
  local git_status="$(git status 2> /dev/null)"
  local on_branch="On branch ([^${IFS}]*)"
  local on_commit="HEAD detached at ([^${IFS}]*)"

  if [[ $git_status =~ $on_branch ]]; then
    local branch=${BASH_REMATCH[1]}
    echo "($branch)"
  elif [[ $git_status =~ $on_commit ]]; then
    local commit=${BASH_REMATCH[1]}
    echo "($commit)"
  fi
}


PS1="\[$WHITE\]\n[\W]"          # basename of pwd
PS1+="\[\$(git_color)\]"        # colors git status
PS1+="\$(git_branch)"           # prints current branch
PS1+="\[$BLUE\]\$\[$RESET\] "   # '#' for root, else '$'
export PS1

#-------------------------------------------------------------------------------

Black="\[\033[0;30m\]"
Red="\[\033[0;31m\]"
Green="\[\033[0;32m\]"
Yellow="\[\033[0;33m\]"
Blue="\[\033[0;34m\]"
Purple="\[\033[0;35m\]"
Cyan="\[\033[0;36m\]"
White="\[\033[0;37m\]"

function _git_prompt() {
  local ansi="$White"
  local status="`git status -unormal 2>&1`"
  if ! echo "$status" | grep -E 'Not a git repo' > /dev/null;
  then
    if echo "$status" | grep -E 'nothing to commit' > /dev/null;
    then
      local ansi=$Yellow
    elif echo "$status" | grep -E 'Changes to be committed' > /dev/null; 
    then
      local ansi=$Green
    elif echo "$status" | grep -E 'nothing added to commit but untracked files present' > /dev/null; 
    then
      local ansi=$Cyan
    else
      local ansi=$Red
    fi
  fi
  echo -n "$Yellow\w $ansi"'$(git_branch)'"\033[0m\n\$ "
}

function _prompt_command() {
  PS1="`_git_prompt`"
}
PROMPT_COMMAND=_prompt_command

