@echo off

set version="v0.1.24"
set msg="enhance file mod"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m "'release %version%'" && git push --tags