FROM scratch

COPY --chown=65536:65536 --chmod=0777 bin/pixie /

# set default ports for rootless use
ENV TFTP :6969
ENV HTTP :8080
ENV SCRIPTS /scripts

EXPOSE 6969/udp
EXPOSE 8080/tcp
VOLUME /scripts

CMD ["/pixie"]
