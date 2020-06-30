FROM centos:7
COPY ./ /bin
RUN /bin/meetup.exe
