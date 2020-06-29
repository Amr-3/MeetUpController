FROM golang
RUN mkdir -p ~/meetup
COPY . ~/meetup
WORKDIR ~/meetup
RUN cd ~/meetup
CMD ["go mod init"]
CMD ["go get -v -t -d ./..."]
CMD ["go build -v ."]
