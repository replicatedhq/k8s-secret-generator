FROM alpine:3.8

COPY ./bin/k8s-secret-generator /bin/k8s-secret-generator

ENTRYPOINT ["/bin/k8s-secret-generator"]
