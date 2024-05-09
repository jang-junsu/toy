FROM mysql:8.0

ENV MYSQL_PASSWORD=pass
ENV MYSQL_ROOT_PASSWORD=pass

VOLUME /var/lib/mysql

COPY ./my.conf /etc/mysql/conf.d/my.conf
