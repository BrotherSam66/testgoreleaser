git add .
git commit -m "test"
git push origin master
git tag -a v0.3.10 -m "test release"
goreleaser release  --rm-dist