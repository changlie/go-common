@echo off

set version="v0.1.18"
set msg="ProgramDir add"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m "'release %version%'" && git push --tags