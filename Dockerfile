FROM harbor.formovie.net/library/alpine:latest

ENV APPS="go-rest-test"
ENV WORK_DIR=/app/${APPS}

ADD dist/go-rest-test*inux_amd64.tar.gz ${WORK_DIR}/

EXPOSE 80

WORKDIR ${WORK_DIR}

CMD ["sh", "-c", "./go-rest-test"]
