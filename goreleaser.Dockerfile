FROM scratch
COPY bookmarks /
COPY .env /
COPY data /data

EXPOSE 8080

ENTRYPOINT ["/bookmarks"]