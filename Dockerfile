##
## Dockerfile for purppura
##
## As our application is a standalone golang binary deployment should
## be pretty simple, however note that we rely upon an external
## MySQL database for storing state.
##
## That means that we need docker-compose to launch/link it up.
##
## Also note that we need to execute an external binary when an
## alert becomes raised.
##
## Building:
##
##    docker build -t purppura:latest .
##
## Launching:
##
##    docker-compose up [-d]
##
## Once launched add the first user, which we'll :
##
##     $ docker-compose exec purppura /app/purppura add-user
##     Enter Username: steve
##     Enter Password: kemp
##
##

##
## Builder-image
##
FROM golang:1.14 AS builder

# Create a working directory
WORKDIR /go/src/app

# Copy our source and build it.
COPY .  .

# Install
RUN go install -v ./...



##
## Runtime image
##
FROM debian:buster

#
# Install dependencies for my notifier
#
RUN apt-get update && apt-get install -y libjson-perl libwww-perl && apt-get clean

#
#
# Use our new workdir with just the copied assets
#
WORKDIR /app

#
# Copy from the build-image to our runtime image.
#
COPY --from=builder /go/bin/purppura /app/
COPY --from=builder /go/src/app/purppura.sql /app/purppura.sql

#
# Expose the port
#
EXPOSE 8080


#
# Ensure the external notifier is discoverable from /srv/bin,
# it is assumed the user will mount it there.
#
ENV PATH="/srv/bin:${PATH}"


#
# Run the application
#
CMD ["/app/purppura", "serve", "-host", "0.0.0.0" ]
