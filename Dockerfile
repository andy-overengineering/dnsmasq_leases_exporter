# build stage
FROM golang:1.16-buster AS build-stage
ADD . /src
ENV CGO_ENABLED=0
WORKDIR /src
RUN go build -o dnsmasq_leases_exporter

# final stage
FROM scratch
WORKDIR /app
COPY --from=build-stage /src/dnsmasq_leases_exporter /app/
USER 65534
ENTRYPOINT ["/app/dnsmasq_leases_exporter"]
