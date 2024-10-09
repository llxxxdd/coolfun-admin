FROM golang:1.22.2-alpine

###############################################################################
#                                INSTALLATION
###############################################################################
# Set project path
ENV WORKDIR /var/www/coolfun
# Add the application executable and set the execution permission
ADD ./coolfun-admin   $WORKDIR/server
RUN chmod +x $WORKDIR/server

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
# Set the environment variables
ENV IS_PROD true
# Run the application
CMD ./server