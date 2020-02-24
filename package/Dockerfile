FROM ubuntu:18.04
RUN apt-get update && apt-get install -y git curl ca-certificates unzip xz-utils && \
    mkdir -p /var/lib/kustomize && \
    chown -R /var/lib/kustomize /usr/local/bin
WORKDIR /var/lib/kustomize

ARG KUSTOMIZE_BRANCH=master

RUN git clone -b ${KUSTOMIZE_BRANCH} --single-branch https://github.com/kubernetes-sigs/kustomize /usr/lib/kustomize

COPY run.sh /run.sh

RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
