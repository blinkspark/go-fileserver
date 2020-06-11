FROM scratch

COPY build/fileserver /bin/fileserver
WORKDIR /public
ENTRYPOINT [ "/bin/fileserver", "-addr", "0.0.0.0:8080" ]
CMD [ "-path", "/public" ]
