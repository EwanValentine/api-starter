FROM centurylink/ca-certs
WORKDIR /app
COPY api /app/
ENTRYPOINT ["./api"]
