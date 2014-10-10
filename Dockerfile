FROM busybox

ADD server /usr/bin/server

EXPOSE 4040

RUN chmod +x /usr/bin/server

ENTRYPOINT /usr/bin/server
