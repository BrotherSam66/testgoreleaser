FROM scratch
# FROM alpine
COPY mybin /app
ENTRYPOINT ["/app"]