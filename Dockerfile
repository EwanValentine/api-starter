FROM centurylink/ca-certs
WORKDIR /app
COPY api /app/
COPY config.json /api/config.json
ENTRYPOINT ["./api"]
