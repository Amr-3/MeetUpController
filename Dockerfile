FROM ubuntu:18.04
COPY ./ /bin/
RUN ls -l
RUN ["/bin/test.exe"]
