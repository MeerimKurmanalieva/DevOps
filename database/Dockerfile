# Use an official MySQL runtime as a parent image
FROM mysql:latest

# Set the environment variables for the database (e.g., root password, database name)
ENV MYSQL_ROOT_PASSWORD=Kurmanalieva93
ENV MYSQL_DATABASE=awsDB

# Expose port 3306
EXPOSE 3306

# Copy SQL initialization script into the container
COPY init.sql /docker-entrypoint-initdb.d/

# Grant privileges to the root user from any host
RUN echo "GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'Kurmanalieva93' WITH GRANT OPTION;" > /docker-entrypoint-initdb.d/init.sql

# Grant privileges to the root user from localhost
RUN echo "GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' IDENTIFIED BY 'Kurmanalieva93' WITH GRANT OPTION;" >> /docker-entrypoint-initdb.d/init.sql
