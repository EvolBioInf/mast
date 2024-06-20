src = annotate.tex shuffle.tex ego.tex sumEgo.tex intro.tex util.tex
date = $(shell bash ../scripts/getDate.sh)
version = $(shell bash ../scripts/getVersion.sh)

all: $(src) gin.pdf 
%.tex:	../*/%.org gin.tex
	awk -f ../scripts/preWeave.awk $< | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > $@
gin.pdf: $(src) ../annotate/iv.tex ../ego/egoTab[123].tex
	echo $(date) | tr '_' ' ' > date.txt
	echo $(version) | tr '-' ' ' | awk '{printf "%s", $$1; if ($$2) printf "-%s", $$2; printf "\n"}' > version.txt
	latex gin
	latex gin
	dvips gin -o -q
	ps2pdf -dALLOWPSTRANSPARENCY gin.ps
	if [ -d /mnt/c/Users/haubold/ownCloud\ -\ haubold@evolbio.mpg.de@owncloud.gwdg.de/docs/ ]; then \
		cp gin.pdf /mnt/c/Users/haubold/ownCloud\ -\ haubold@evolbio.mpg.de@owncloud.gwdg.de/docs/; \
	fi
clean:
	rm -f $(src) *.pdf *.aux *.bbl *.blg *.dvi *.log *.toc
