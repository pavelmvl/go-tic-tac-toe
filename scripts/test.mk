
PKGS := $(wildcard internal/*)

test: ${PKGS}
	for module in ${<}; do echo "=== test package $< ==="; go test -v $${module}/*.go; done
