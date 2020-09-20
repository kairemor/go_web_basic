FROM golang:1.13 AS Builder


WORKDIR /myapp

COPY . ./
RUN useradd api 
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o api 

FROM scratch 
COPY --from=Builder /myapp/api /usr/bin/api 
COPY --from=Builder /etc/passwd /etc

USER api 
EXPOSE 4000 

ENTRYPOINT ["/usr/bin/api"]