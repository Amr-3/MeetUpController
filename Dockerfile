FROM ubuntu:18.04
COPY ./ /bin
RUN ["cd /bin"]
RUN ["test.exe"]
