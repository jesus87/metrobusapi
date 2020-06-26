FROM alpine:latest
RUN apk --no-cache add tzdata
WORKDIR /root
COPY build/ .
EXPOSE ${CONTAINER_PORT}
CMD ["./graphql"]