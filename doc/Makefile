src = mast.tex tutorial.tex util.tex
txt = mastDoc
date = $(shell bash ../scripts/getDate.sh)
version = $(shell bash ../scripts/getVersion.sh)

all: $(src) $(txt).pdf
%.tex:	../*/%.org $(txt).tex
	awk -f ../scripts/preWeave.awk $< | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > $@
$(txt).pdf: $(src) ref.bib mlAlg.tex intro.tex ../tutorial/mast.ps
	echo $(date) | tr '_' ' ' > date.txt
	echo $(version) | tr '-' ' ' | awk '{printf "%s", $$1; if ($$2) printf "-%s", $$2; printf "\n"}' > version.txt
	latex $(txt)
	bibtex $(txt)
	latex $(txt)
	latex $(txt)
	dvips $(txt) -o -q
	ps2pdf -dALLOWPSTRANSPARENCY $(txt).ps
publish: $(txt).pdf
	if [ -d /mnt/c/Users/haubold/ownCloud\ -\ haubold@evolbio.mpg.de@owncloud.gwdg.de/docs/ ]; then \
		cp $(txt).pdf /mnt/c/Users/haubold/ownCloud\ -\ haubold@evolbio.mpg.de@owncloud.gwdg.de/docs/; \
	fi
clean:
	rm -f $(src) *.pdf *.aux *.bbl *.blg *.dvi *.log *.toc
