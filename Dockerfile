FROM test.cargo.io/library/alpine:3.7

EXPOSE 80
COPY cicd-go /usr/local/bin/

ENTRYPOINT [ "cicd-go" ]
