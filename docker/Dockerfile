FROM alpine:3.1 
# Copy in compiled output from ember app and go binaries
COPY .rootfs/ /
# Default command points to database running on host machine
CMD PORT=80 ./commish --conn='WebClient@tcp(10.0.2.2:3306)/auth' > /run/commish_log.txt 2> /run/commish_errors.txt
EXPOSE 8081
