@echo off

set version="v0.1.17"
set msg="release script test v5"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m "'release %version%'" && git push --tags