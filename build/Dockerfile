FROM public.ecr.aws/docker/library/golang:1.24.5

RUN apt update && apt upgrade -y && apt install -y fakeroot shellcheck zip

ADD build.sh /build/

CMD [ "bash", "-x", "/build/build.sh" ]
