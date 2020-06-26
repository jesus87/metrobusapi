FROM alpine:latest
COPY entry.sh .
COPY build/scheduler  /etc/periodic/15min/scheduler
RUN chmod 755 /entry.sh
CMD ["./entry.sh"]