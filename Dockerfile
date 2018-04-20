FROM alpine:3.7
ADD ./static /static
ADD ./timescale /timescale
RUN addgroup -S app && adduser -S -G app app 
ENV PORT=3000
USER app
CMD ["/timescale"]
