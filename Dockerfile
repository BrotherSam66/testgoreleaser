FROM scratch
COPY mybin /app
ENTRYPOINT ["/app"]