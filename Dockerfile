FROM centos:7.6.1810
RUN mkdir -p /opt/tweek
COPY txterv8 /opt/tweek

WORKDIR /opt/tweek
ENTRYPOINT ["./txterv8"]

EXPOSE 80
