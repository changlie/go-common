
ci "type conversion"

git push

newVersion=`newVersion`

git tag -a $newVersion -m "release $newVersion"
git push --tags

echo "release $newVersion successfully!"