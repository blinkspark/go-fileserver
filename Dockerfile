FROM scrach

COPY buid/fileserver /bin/fileserver
WORKDIR /public
ENTRYPOINT [ "/bin/fileserver" ]
CMD [ "-path", "/public" ]
