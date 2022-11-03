git add .
# testcommit: your commite message
git commit -m "testactions"
# vXX.XX.XX: your version. your tag
git tag -a v0.3.88 -m "test release"
# master: your branch name 
git push origin master
# release and push to github & docker.io
goreleaser release  --rm-dist
# only test realase
# goreleaser --snapshot --skip-publish --rm-dist