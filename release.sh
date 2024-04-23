
set -xe

msg=""
newVersion="v0.1.13"

ci "type conversion"
git push


git tag -a $newVersion -m "release $newVersion"
git push --tags

echo "release $newVersion successfully!"


# git tag -a v0.1.13 -m "release " && git push --tags
# git add . && git commit . -m "" && git push && git tag -a v0.1.14 -m "release v0.1.14" && git push --tags