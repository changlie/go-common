@echo off

set version="v0.1.27"
set msg="optimize json; http client; types"

git add . && git commit . -m '%msg%' && git push && git tag -a %version% -m "'release %version%'" && git push --tags