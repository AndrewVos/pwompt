# pwompt

A PS1 written in go

![screenshot](https://github.com/AndrewVos/pwompt/raw/master/screenshot.png)

## Installation

```
go get github.com/AndrewVos/pwompt
```

Add this to your `~/.bashrc` or wherever:
```
export PWOMPT_CONFIG='battery_charging?("white", "⏚")not_battery_charging?("white", "⌁")battery_percentage("red", "yellow", "green")battery?("white", " ")c("yellow", "[")user("magenta")c("white",":")cwd_short("blue")c("yellow", "]")git?("white", " ")git_branch("magenta")git_dirty?("red", "*")last_exit_failed?("white", " ")last_exit_code("red")git?("white", " ±")not_git?("white", " $")c("white", " ")'
PS1='$(PWOMPT_LAST_EXIT_CODE=$? pwompt)'
```

## Examples

Battery, username, working directory, git, exit code:

```
export PWOMPT_CONFIG='battery_charging?("white", "⏚")not_battery_charging?("white", "⌁")battery_percentage("red", "yellow", "green")battery?("white", " ")c("yellow", "[")user("magenta")c("white",":")cwd_short("blue")c("yellow", "]")git?("white", " ")git_branch("magenta")git_dirty?("red", "*")last_exit_failed?("white", " ")last_exit_code("red")git?("white", " ±")not_git?("white", " $")c("white", " ")'
```

Username, working directory, git, exit code:

```
export PWOMPT_CONFIG='c("yellow", "[")user("magenta")c("white",":")cwd_short("blue")c("yellow", "]")git?("white", " ")git_branch("magenta")git_dirty?("red", "*")last_exit_failed?("white", " ")last_exit_code("red")git?("white", " ±")not_git?("white", " $")c("white", " ")'
```

Username and working directory:

```
export PWOMPT_CONFIG='c("yellow", "[")user("magenta")c("white",":")cwd_short("blue")c("yellow", "]")c("white", " $ ")'
```

Working directory with prompt on new line:

```
export PWOMPT_CONFIG='c("yellow", "[")user("magenta")c("white",":")cwd("blue")c("yellow", "]")c("white", "\n$ ")'
```

## Methods

These methods, if they return true, write custom text out in a certain colour.
Each method has an alternative negative version, for example `something?` and `not_something?`.

```
# current directory is inside a git repository
[not_]git?("colour", "text")

# current directory is inside a dirty git repository
[not_]git_dirty?("colour", "text")

# last command executed had a non-zero exit code
[not_]last_exit_failed?("colour", "text")

# battery exists and acpi is installed
[not_]battery?("colour", "text")

# battery is charging
[not_]battery_charging?("colour", "text")
```

These commands output some text, in any colour:
```
# current working directory
cwd("colour")

# current working directory shortened
cwd_short("colour")

# current git branch
git_branch("colour")

# last command exit code
last_exit_code("colour")
```

Other methods:

```
# write out any text you want
c("colour", "text")

# write out the battery percentage, with different colours for different states
battery_percentage("low-battery-colour", "medium-battery-colour", "high-battery-colour")
```
