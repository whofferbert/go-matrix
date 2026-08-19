#!/bin/bash
# by William Hofferbert

#
# Environmental Variables - Default Settings
#

script_path=$( cd "$(dirname "$0")" ; pwd -P )
script_name=$(basename $0)
self=$script_path/$script_name

bold=$(tput bold)
reg=$(tput sgr0)
name_bold="${bold}${script_name}${reg}"

buildscript_basedir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

#output_binary=${buildscript_basedir##*/}
output_binary="matrix"

cgo=0

#
# Functions
#

helps () {
fold -s -w $(tput cols) << EOF

  This script simply runs the proper `go build` command to construct the final
  executable from the project source.


Usage

  $name_bold [options] ${bold}[REQUIRED ARG]${reg}

Options:

  --cgo        -c   Use "cgo" with libmusl for static binary.

  --output     -o   [name]
                    The name of the binary to build.

  --help       -h   Display this help text.

EOF
exit 0
}

handle_args () {
  OPTS=$(getopt -o ho:c --long "help,output:,cgo"  -- "$@")
  #OPTS=$(getopt -o hv: --long "help,value:"  -- "$@")
  eval set -- "$OPTS"
  while true
    do
      case "$1" in
        -h|--help) helps; shift;;
        -o|--output) output_binary=$2; shift 2;;
		-c|--cgo) cgo=1; shift;;
        --) shift; break;;
      esac
  done
  #first_arg_leftover=$1
}

err () {
  echo "$1" >&2
  exit 2
}

warn () {
  echo "$1" >&2
}

cleanup () {
  warn "Signal received ..."
  exit 1
}

sanity () {
  #[ -z "$first_arg_leftover" ] && err "$script_name requires argument _!"
  echo -en ''
}

#
# main
#

main () {
  # pass arguments to the arg handler
  handle_args $@
  sanity
  trap cleanup SIGHUP SIGINT SIGKILL
  # more logical steps here...

  if [ -e "src" ] ; then
    cd src
	output_binary="../$output_binary"
  fi

  buildcmd=""
  if [ $cgo -eq 1 ] ; then
    buildcmd="CGO_ENABLED=1 CC=musl-gcc go build -o $output_binary --ldflags '-w -s -linkmode external -extldflags=-static' *.go"
  else 
    buildcmd="CGO_ENABLED=0 go build -o $output_binary --ldflags '-w -s -extldflags=-static' *.go"
  fi
  echo "RUNNING: $buildcmd"
  sh -c "$buildcmd"
  
}

# pass all script arguments to main
main $@
