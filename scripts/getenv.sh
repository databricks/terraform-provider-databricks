# This script is intended to be `source`d (e.g. `source getenv.sh`)
# and it will parse the .env file and export the variables defined within it
# TODO - handle commented lines :-)

while read line; do
  NAME=$(echo "$line" | awk '{split($1,parts, "="); print parts[1]}')
  VALUE=$(echo "$line" | awk '{split($1,parts, "="); print parts[2]}')
  export $NAME=$VALUE
done <.env