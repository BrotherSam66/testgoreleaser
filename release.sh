git add .
git commit -m "test"
git push origin master
git tag -a v0.3.11 -m "test release"
goreleaser release  --rm-dist