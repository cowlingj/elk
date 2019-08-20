FROM docker.elastic.co/beats/filebeat:7.3.0
USER root
RUN rm /usr/share/filebeat/filebeat.reference.yml
COPY --chown=root:filebeat filebeat.yml /usr/share/filebeat/filebeat.yml
USER filebeat
RUN ./filebeat modules enable logstash
ENTRYPOINT [ "./filebeat", "-e" ]

