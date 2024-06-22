package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/esa"
	"github.com/evolbioinf/fasta"
	"github.com/evolbioinf/mast/util"
	"io"
	"log"
	"os"
)

func parse(r io.Reader, args ...interface{}) {
	qseqs := args[0].([]*fasta.Sequence)
	sc := fasta.NewScanner(r)
	for sc.ScanSequence() {
		rseq := sc.Sequence()
		for _, qseq := range qseqs {
			q := qseq.Data()
			r := rseq.Data()
			ml := make([]int, len(q))
			esa := esa.MakeEsa(r)
			i := 0
			for i < len(q) {
				l := esa.MatchPref(q[i:]).L
				ml[i] = l
				i += l + 1
			}
			for i := 1; i < len(ml); i++ {
				if ml[i] < ml[i-1]-1 {
					ml[i] = ml[i-1] - 1
				}
			}
			ms := make([]int, len(q))
			for i := 0; i < len(q); i++ {
				ms[i] = esa.MatchPref(q[i:]).L
			}
			fmt.Printf("# Query: %s; reference: %s\n",
				qseq.Header(), rseq.Header())
			fmt.Printf("# i\tms[i]\tml[i]\n")
			for i := 0; i < len(ml); i++ {
				fmt.Printf("%d\t%d\t%d\n", i+1, ms[i], ml[i])
			}
		}
	}
}
func main() {
	clio.PrepLog("mast")
	u := "mast [option]... q.fasta [r1.fasta]..."
	p := "Calculate matching statistics and " +
		"match lengths for q w.r.t. r"
	e := "mast q.fasta r.fasta"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	flag.Parse()
	if *optV {
		util.Version("mast")
	}
	files := flag.Args()
	if len(files) < 1 {
		m := "please provide the " +
			"name of a query file"
		log.Fatal(m)
	}
	qseqs := make([]*fasta.Sequence, 0)
	file, err := os.Open(files[0])
	if err != nil {
		log.Fatal(err)
	}
	scanner := fasta.NewScanner(file)
	for scanner.ScanSequence() {
		qseqs = append(qseqs, scanner.Sequence())
	}
	refSeqs := files[1:]
	clio.ParseFiles(refSeqs, parse, qseqs)
}
