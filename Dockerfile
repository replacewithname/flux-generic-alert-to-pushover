FROM debian:11.3-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY flux-generic-alert-to-pushover /usr/local/bin

CMD ["flux-generic-alert-to-pushover"]