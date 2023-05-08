FROM gcr.io/distroless/static

# Copy the binary that goreleaser built
COPY metadata-api /metadata-api

# Run the web service on container startup.
ENTRYPOINT ["/metadata-api"]
CMD ["serve"]
