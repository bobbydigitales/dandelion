#!/bin/bash

# Set the service type
service_type="_http._tcp"

# Begin script
echo "Starting script to search for $service_type services..."

# Browse for devices advertising the service
dns-sd -B $service_type | while read -r timestamp junk type instance more_junk; do
  # Filter out non-relevant lines
  if [[ ! $type == $service_type || -z $instance ]]; then
    continue
  fi

  # Log the extracted instance name
  echo "Found instance: $instance"

  # Look up the instance details to get the hostname
  hostname=$(dns-sd -L "$instance" $service_type 2>/dev/null | grep -Eo '[^ ]+\.local' | head -n 1)
  
  if [ -z "$hostname" ]; then
    echo "Hostname not found for instance: $instance"
  else
    echo "Hostname for $instance is: $hostname"
  fi
done

echo "Script completed."
