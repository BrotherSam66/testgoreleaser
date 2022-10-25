FROM scratch
# FROM alpine
COPY zexample /app
ENTRYPOINT ["/app"]