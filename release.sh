
set -xe

ci "type conversion"

git push

newVersion=`nextVersion`

git tag -a $newVersion -m "release $newVersion"
git push --tags

echo "release $newVersion successfully!"