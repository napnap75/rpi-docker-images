#!/bin/bash

# First check if the gandi cli is properly configured, otherwise configure it

while true ; do
  # Get my current IP
  my_ip=$(curl -s https://api.ipify.org)

  # Get my currently registered IP
  current_record=$(curl -s -H"X-Api-Key: $GANDI_API_KEY" https://dns.api.gandi.net/api/v5/domains/$GANDI_DOMAIN/records | jq -c '.[] | select(.rrset_name == "'$GANDI_HOST'") | select(.rrset_type == "A")')
  current_ip=$(echo $current_record | jq -r '.rrset_values[0]')

  # Check if both IP addresses are correct
  if [[ "$my_ip" =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ && "$current_ip" =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]] ; then
    # If they do not match, change it (and keep the TTL and TYPE)
    if [[ "$my_ip" != "$current_ip" ]]; then
      echo "Updating $GANDI_HOST.$GANDI_DOMAIN record with IP $my_ip"
      current_ttl=$(echo $current_record | jq -r '.rrset_ttl')
      curl -s -X PUT -H "Content-Type: application/json" -H "X-Api-Key: $GANDI_API_KEY" -d '{"rrset_ttl": '$current_ttl', "rrset_values":["'$my_ip'"]}' https://dns.api.gandi.net/api/v5/domains/$GANDI_DOMAIN/records/$GANDI_HOST/A

      # If the update was OK
      if [[ $? == 0 ]] ; then
        # Send a notification to Slack
        if [[ "$SLACK_URL" != "" ]] ; then
          curl -o /dev/null -s -X POST -d "payload={\"username\": \"gandi\", \"icon_emoji\": \":dart:\", \"text\": \"New IP $my_ip for host $GANDI_HOST.$GANDI_DOMAIN\"}" $SLACK_URL
        fi

        # Send a notification to Healthchecks
        if [[ "$HEALTHCHECKS_URL" != "" ]] ; then
          curl -o /dev/null -s -m 10 --retry 5 $HEALTHCHECKS_URL
        fi
      fi
    else
      # Send a notification to Healthchecks
      if [[ "$HEALTHCHECKS_URL" != "" ]] ; then
        curl -o /dev/null -s -m 10 --retry 5 $HEALTHCHECKS_URL
      fi
    fi
  fi

  # Wait 5 minutes
  sleep 300
done
