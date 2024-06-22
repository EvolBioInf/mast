stan -t 1 -n 1 -o
mast targets/t1.fasta neighbors/n1.fasta > mast.out
head mast.out
tail -n +3 mast.out |
    awk '{print $1, $2, "ms"}' >  mast.dat
tail -n +3 mast.out |
    awk '{print $1, $3, "ml"}' >> mast.dat
plotLine -x "Position" -y "Length" mast.dat
