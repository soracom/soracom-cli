FROM golang:1.16.5-buster

RUN apt update && apt upgrade -y && apt install -y fakeroot shellcheck zip

ADD build.sh /build/

CMD [ "bash", "-x", "/build/build.sh" ]
