FROM golang AS builder

WORKDIR /build

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Retrieve application dependencies
# This allows the container build to reuse cached dependencies
COPY go.* ./
RUN go mod download

# Compile
COPY . .
RUN go install cmd/server.go

FROM ubuntu 

WORKDIR /runtime

COPY --from=builder /go/bin/server ./
COPY --from=builder /go/bin/migrate ./

COPY migrations migrations

EXPOSE 8081

CMD ./migrate -path migrations -database $DB_CONNECTION_URL up && \
    ./server