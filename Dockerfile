FROM alpine:3.4
RUN apk --no-cache add ca-certificates && update-ca-certificates
ADD ./holayoga-dialogflow-service /usr/local/bin
ENTRYPOINT ["/usr/local/bin/holayoga-dialogflow-service"]