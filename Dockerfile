FROM scratch

ADD ./sdr-app /sdr-app

CMD ["/sdr-app"]