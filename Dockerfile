FROM ghcr.io/jenkins-x/jx-boot:latest

ENTRYPOINT ["jx-tap"]

COPY ./build/linux/jx-tap /usr/bin/jx-tap