FROM gcr.io/jenkinsxio/jx-cli-base:0.0.21

ENTRYPOINT ["jx-tap"]

COPY ./build/linux/jx-tap /usr/bin/jx-tap