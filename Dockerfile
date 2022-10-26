FROM scratch
# FROM alpine
COPY guanceexample /app
ENTRYPOINT ["/app"]