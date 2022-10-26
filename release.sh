git add .
git commit -m "test"
git push origin master
git tag -a v0.3.16 -m "test release"
# goreleaser release  --rm-dist
goreleaser --snapshot --skip-publish --rm-dist