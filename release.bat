@echo off

set version="v0.1.15"
set msg="release script test v2"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m 'release %version%' && git push --tags