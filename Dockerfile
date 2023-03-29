FROM golang:latest
RUN go env -w GO111MODULE=auto
RUN go get github.com/emersion/go-imap
RUN go get github.com/emersion/go-message
RUN go get github.com/lib/pq
RUN go get gopkg.in/gomail.v2
RUN go get github.com/emersion/go-sasl
ADD mgp2 .
CMD ["go","run","main.go"]