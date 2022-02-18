FROM alpine

ADD ./bin/photoservice /app/bin/photoservice
RUN ls
RUN pwd
COPY ./dist ./app/bin/web
RUN chmod +x /app/bin/photoservice
WORKDIR /app
RUN apk --no-cache add ca-certificates
CMD ["/app/bin/photoservice", "-d=true", "-log=false"]