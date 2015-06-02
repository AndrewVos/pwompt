# pwompt

A PS1 written in go

![screenshot](https://github.com/AndrewVos/pwompt/raw/master/screenshot.png)

## Installation

```
go get github.com/AndrewVos/pwompt
```

Add this to your `~/.bashrc` or wherever:
```
export PWOMPT_CONFIG='battery_charging?("white", "⏚")battery_discharging?("white", "⌁")battery_percentage("red", "yellow", "green")battery?("white", " ")c("yellow", "[")cwd_short("blue")c("yellow", "] ")git_branch("red")git_dirty?("red", "* ")last_exit_code("magenta")last_exit_failed?("white", " ")git?("white", "±")not_git?("white", "$")c("white", " ")'
PS1='$(PWOMPT_LAST_EXIT_CODE=$? pwompt)'
```

## Arguments

```
PS1='$(PWOMPT_LAST_EXIT_CODE=$? pwompt)' >> ~/.bashrc
⌁87% [~/.../pwompt] master* ± 

PS1='$(PWOMPT_LAST_EXIT_CODE=$? pwompt -shorten-path=false)' >> ~/.bashrc
⌁87% [~/gopath/src/github.com/AndrewVos/pwompt] master ± 

PS1='$(PWOMPT_LAST_EXIT_CODE=$? pwompt -show-battery=false)' >> ~/.bashrc
[~/.../pwompt] master* 130 ± 
```

```
battery_charging?("white", "⏚")battery_discharging?("white", "⌁")battery_percentage("red", "yellow", "green")battery?("white", " ")c("yellow", "[")cwd("yellow")c("yellow", "] ")git_branch("red")git_dirty("red")git?("white", " ± ")!git("white", " $ ")
```
