git add .
git commit -m "test"
git push origin master
git tag -a v0.1.10 -m "First release"
goreleaser release  --rm-dist