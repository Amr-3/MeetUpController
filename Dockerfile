FROM centos:7
COPY ./ /bin
CMD ["/bin/test.exe"]
