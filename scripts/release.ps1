$tag = $args[0]

if (-Not ($tag -match '^v[0-9].[0-9].[0-9]')) {
	Write-Output "Bad version number"
	exit 1
}

git checkout -b release-$tag
git tag $tag
git push -u origin release-$tag
git push -u origin $tag
git checkout main
git merge release-$tag
