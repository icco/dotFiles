#!/usr/bin/env bash

# Given an email, get the associated gravatar and put it in the users ~/.face file.

EMAIL='nat@natwelch.com'

# Size in pixels you want, must be less than 512
SIZE='256'
HASH=`echo -n $EMAIL | awk '{print tolower($0)}' | tr -d '\n ' | md5sum --text | tr -d '\- '`
URL="http://www.gravatar.com/avatar/$HASH?s=$SIZE&d=404"

# Alright, grab the file, store it.
curl -s $URL > ~/.face

# A test for nat@natwelch.com
# echo "http://www.gravatar.com/avatar/229e3746f6f5100c1d7d5d7a8a5b82d5?s=200&d=404"
