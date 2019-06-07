FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o httptester

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/httptester /app/
ENTRYPOINT [ "./httptester" ]
CMD [ "-port", "8080" ]
