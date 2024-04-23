@echo off

set version="v0.1.20"
set msg="optimize Http Server"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m "'release %version%'" && git push --tags