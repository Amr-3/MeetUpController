FROM ubuntu:18.04
COPY ./ /bin/
RUN ls -l /bin/
RUN ["/bin/test.exe"]
