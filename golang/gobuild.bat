@echo off

setlocal

if exist build.bat goto ok
echo build.bat must be run from its folder
goto end

:ok

set OLDGOPATH=%GOPATH%
rem set GOPATH=%~dp0
set GOPATH=%cd%

rem set GOPATH

gofmt -w src

go build main

set GOPATH=OLDGOPATH

:end
echo finished
