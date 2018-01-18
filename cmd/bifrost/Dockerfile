FROM scratch
MAINTAINER Alex Peters <info@alexanderpeters.de>

COPY /pkg/database/migrations/ /migrations/
COPY /cmd/bifrost/bifrost /
COPY /certs/cacert.pem /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/bifrost"]
CMD [""]