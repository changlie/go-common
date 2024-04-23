@echo off

set version="v0.1.21"
set msg="optimize Http Server v4"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m "'release %version%'" && git push --tags