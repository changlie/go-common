@echo off

set version="v0.1.19"
set msg="optimize Date & Http Server"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m "'release %version%'" && git push --tags