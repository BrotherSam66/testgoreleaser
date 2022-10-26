git add .
git commit -m "test"
git push origin master
git tag -a v0.3.17 -m "test release"
goreleaser release  --rm-dist
# goreleaser --snapshot --skip-publish --rm-dist