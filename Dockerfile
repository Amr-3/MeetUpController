FROM ubuntu:18.04
COPY ./ /bin/
RUN ["/bin/test.exe"]
