@echo off

set version="v0.1.23"
set msg="optimize http server v6"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m "'release %version%'" && git push --tags