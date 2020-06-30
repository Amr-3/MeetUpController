FROM centos:7
COPY ./ /bin
RUN /bin/sh -c "/bin/meetup.exe"
