# pwompt

A PS1 written in go

## Installation

```
go get github.com/AndrewVos/pwompt
echo PS1=\'$\(PWOMPT_LAST_EXIT_CODE=\$? pwompt\)\' >> ~/.bashrc
```
