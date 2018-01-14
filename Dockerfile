FROM scratch
COPY thea /thea
EXPOSE  80 443 1323
CMD ["/thea"]
