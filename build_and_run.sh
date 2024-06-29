RED=`tput setaf 1`
GREEN=`tput setaf 2`
NC=`tput sgr0`
BROWN=`tput setaf 5`

exec_command() {
	echo "Executing: $2	"
	sleep 2
	$2
	if [ $? -ne 0 ]; then
		echo "${RD} $1 Failed"; exit 1
	fi
	echo "${BROWN} $1 ${GREEN} SUCCESS ${NC}"
}

exec_command "Getting Code from Git" "git clone https://github.com/skoyya/cronparser.git"
exec_command "Changing to workspace" "cd cronparser"
export GOPATH=`pwd`
exec_command "unit-test praser" "go test -v parser"
sleep 1
go run main.go
