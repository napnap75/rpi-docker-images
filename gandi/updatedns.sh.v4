#!/bin/bash

# First check if the gandi cli is properly configured, otherwise configure it
if ! (gandi config get api.key >> /dev/null && gandi config get api.host >> /dev/null); then
  gandi config set api.host "https://rpc.gandi.net/xmlrpc/"
  gandi config set api.key "$GANDI_API_KEY"
fi

while true ; do
  # Get my current IP
  my_ip=$(curl -s https://api.ipify.org)

  # Get my currently registered IP
  current_record=$(gandi record list -f json $GANDI_DOMAIN | jq -c '.[] | select(.name == "'$GANDI_HOST'")')
  current_ip=$(echo $current_record | jq -r '.value')

  # If they do not match, change it (and keep the TTL and TYPE)
  if [[ "$my_ip" != "$current_ip" ]]; then
    echo "Updating $GANDI_HOST.$GANDI_DOMAIN record with IP $my_ip"
    host_string="$GANDI_HOST $(echo $current_record | jq '.ttl') $(echo $current_record | jq -r '.type')"
    gandi record update -r "$host_string $current_ip" --new-record "$host_string $my_ip" $GANDI_DOMAIN
  fi

  # Wait 5 minutes
  sleep 300
done
