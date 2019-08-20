FROM docker.elastic.co/logstash/logstash:7.3.0

EXPOSE 80

RUN rm -f /usr/share/logstash/pipeline/logstash.conf

COPY pipeline/ /usr/share/logstash/pipeline/
COPY config/ /usr/share/logstash/config/
