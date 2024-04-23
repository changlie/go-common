@echo off

set version="v0.1.16"
set msg="release script test v4"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m 'release %version%' && git push --tags