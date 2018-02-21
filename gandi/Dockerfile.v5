FROM napnap75/rpi-alpine-base:latest

# Add dependencies
RUN apk add --no-cache curl jq

# Add and run the update script
ADD updatedns.sh.v5 /usr/bin/updatedns.sh
RUN chmod +x /usr/bin/updatedns.sh
CMD /usr/bin/updatedns.sh
