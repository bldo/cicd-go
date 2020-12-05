FROM alpine

EXPOSE 80
COPY cicd-go /usr/local/bin/

ENTRYPOINT [ "cicd-go" ]
