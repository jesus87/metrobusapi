FROM alpine:latest
COPY build/scheduler  /etc/periodic/15min/scheduler
CMD ["/usr/sbin/crond", "-f"]