
set -xe

ci "type conversion"

git push

newVersion=`nextVersion`

git tag -a $newVersion -m "release $newVersion"
git push --tags

echo "release $newVersion successfully!"


# git tag -a v0.1.13 -m "release v0.1.13" && git push --tags

# git add . && git commit . -m "optimize file mod" && git tag -a v0.1.13 -m "release v0.1.13" && git push --tags