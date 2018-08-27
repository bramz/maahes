FROM alpine:latest

LABEL maintainer="Brock Ramsey <brockramz@gmail.com>"
LABEL description="Container for maahes discord bot"

CMD ["./bin/start.sh"]