git add .
# testcommit: your commite message
git commit -m "testcommit"
# master: your branch name 
git push origin master
# vXX.XX.XX: your version. your tag
git tag -a v0.3.47 -m "test release"
# release and push to github & docker.io
goreleaser release  --rm-dist
# only test realase
# goreleaser --snapshot --skip-publish --rm-dist