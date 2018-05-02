# cmd build  =>  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o learning-assistain .
FROM ubuntu

COPY learning-assistain /

#создаем папки
#RUN mkdir -p /database
#RUN mkdir -p /public
#RUN mkdir -p /templates
#RUN mkdir -p /view

COPY ./database/learnig.db.data /database/
COPY ./public/ /public/
COPY ./templates/ /templates/
COPY ./view/ /view/

#ADD ca-certificates.crt /etc/ssl/certs/
VOLUME /database
EXPOSE 4000
CMD ["/learning-assistain"]