FROM centurylink/ca-certs
WORKDIR /app
COPY api /app/
COPY config.json /app/config.json
ENTRYPOINT ["./api"]
