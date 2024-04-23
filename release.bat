@echo off

set version="v0.1.22"
set msg="optimize file mod"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m "'release %version%'" && git push --tags