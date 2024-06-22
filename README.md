# [`mast`](https://owncloud.gwdg.de/index.php/s/ff7rGTo3Pz9B4VL)
## Description
Explore matching statistics.
## Author
[Bernhard Haubold](http://guanine.evolbio.mpg.de/), `haubold@evolbio.mpg.de`
## Make the Program

```
make
```

The program `mast` is now in directory `bin`. To apply `mast` to a
small [query](mast/q1.fasta) and [reference](mast/r1.fasta),

```
./bin/mast mast/q1.fasta mast/r1.fasta
```

This yields

```
# Query: q1; reference: r1
# Pos   m_s     m_l
1       5       5
2       4       4
3       3       3
4       2       2
5       2       1
6       5       0
7       4       4
8       3       3
9       2       2
10      1       1
```

To run the [query](mast/q2.fasta) and [reference](mast/r2.fasta) from
the [Fastms](https://github.com/odenas/indexed_ms) website,

```
./bin/mast mast/q2.fasta mast/r2.fasta
```

which yields

```
# Query: q2; reference: r2
# Pos   m_s     m_l
1       3       3
2       4       2
3       3       1
4       2       0
5       3       3
6       2       2
7       3       1
8       3       0
9       4       4
10      3       3
11      2       2
12      1       1
```
