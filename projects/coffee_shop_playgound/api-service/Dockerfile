FROM golang:latest as build

WORKDIR /usr/src/

ARG pat

RUN git config --global url."https://pat:$pat@dev.azure.com/mdtcrhf/CRHFMLifeInfrastructure/_git/clctl".insteadOf "https://dev.azure.com/mdtchrf/CRHFMLifeInfrastructure/clctl"

ENV GOPRIVATE="dev.azure.com"

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o clctl-api

FROM scratch

COPY --from=build /usr/src/clctl-api .

EXPOSE 8080

CMD ["./clctl-api"]