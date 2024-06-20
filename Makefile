progs = mast
packs = util

all:
	test -d bin || mkdir bin
	for pack in $(packs); do \
		make -C $$pack; \
	done
	for prog in $(progs); do \
		make -C $$prog; \
		cp $$prog/$$prog bin; \
	done
tangle:
	for pack in $(packs) $(progs); do \
		make tangle -C $$pack; \
	done
.PHONY: doc test
doc:
	make -C doc
clean:
	for prog in $(progs) $(packs) doc; do \
		make clean -C $$prog; \
	done
test: data
	echo test
	for prog in $(packs) $(progs); do \
		make test -C $$prog; \
	done
