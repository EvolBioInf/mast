nw = $(shell which noweb)
all: tutorial.sh
tutorial.sh: tutorial.org
	if [ $(nw) != "" ]; then\
		bash ../scripts/org2nw tutorial.org | notangle -Rtut.sh > tutorial.sh;\
	fi
clean:
	rm tutorial.sh

