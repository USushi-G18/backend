FROM golang AS builder

WORKDIR /build

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Retrieve application dependencies
# This allows the container build to reuse cached dependencies
COPY go.* ./
RUN go mod download

# Compile
COPY . .
RUN go install cmd/admin/admin.go

FROM ubuntu 

WORKDIR /runtime

COPY --from=builder /go/bin/admin ./
COPY --from=builder /go/bin/migrate ./
COPY --from=builder /build/secrets ./secrets

COPY migrations migrations

EXPOSE 8081

CMD ./migrate -path migrations -database $DB_CONNECTION_URL up && \
    ./admin