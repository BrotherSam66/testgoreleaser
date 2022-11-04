FROM scratch
# FROM alpine
COPY guanceexample /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]