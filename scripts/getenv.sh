# This script is intended to be `source`d (e.g. `source getenv.sh`)
# and it will parse the .env file and export the variables defined within it
# TODO - handle commented lines :-)

process_line(){
  line=$1
  if [[ $line == "" ]]; then
    return
  fi
  line_start=${line:0:1} # Test if the line starts with a comment and skip
  if [[ $line_start == "#" ]]; then
    return
  fi
  NAME=$(echo "$line" | IFS= awk '{split($1,parts, "="); print parts[1]}')
  VALUE=$(echo "$line" | IFS= awk -F\" '{split($1,parts, "="); print parts[2]}')
  export $NAME="$VALUE"
}

while IFS= read line; do # `IFS= ` removes the default space separator
  process_line "$line"
done <.env
process_line "$line" # handle a last line without a newline char
