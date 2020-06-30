FROM ubuntu:18.04
COPY ./ /bin
CMD ["cd /bin"]
CMD ["test.exe"]
