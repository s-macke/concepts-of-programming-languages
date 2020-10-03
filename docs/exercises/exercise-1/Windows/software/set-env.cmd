@echo off

type ascii-art.txt

set SEU_HOME=G:

set GOROOT=%SEU_HOME%\software\goroot
set GOPATH=%SEU_HOME%\codebase
set GOBIN=%SEU_HOME%\codebase\bin
set PATH=%SEU_HOME%\software\goroot\bin;%PATH%
