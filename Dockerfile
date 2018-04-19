FROM alpine:3.7
ADD ./static /static
ADD ./timescale /timescale
ADD ./conf.yaml /conf.yaml
RUN addgroup -S app && adduser -S -G app app 
USER app
CMD ["/timescale"]
