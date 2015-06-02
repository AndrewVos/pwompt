# pwompt

A PS1 written in go

## Installation

```
go get github.com/AndrewVos/pwompt
echo PS1=\'$\(PWOMPT_LAST_EXIT_CODE=\$? pwompt\)\' >> ~/.bashrc
```

## Examples

```
PS1='$(PWOMPT_LAST_EXIT_CODE=$? pwompt)' >> ~/.bashrc
⌁87% [~/.../pwompt] master* ± 

PS1='$(PWOMPT_LAST_EXIT_CODE=$? pwompt -shorten-path=false)' >> ~/.bashrc
⌁87% [~/gopath/src/github.com/AndrewVos/pwompt] master ± 

PS1='$(PWOMPT_LAST_EXIT_CODE=$? pwompt -show-battery=false)' >> ~/.bashrc
[~/.../pwompt] master* 130 ± 
```
