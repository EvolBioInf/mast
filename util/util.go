// The package util summarizes functions for the programs in the mast repo.
package util

import (
	"github.com/evolbioinf/clio"
	"os"
)

var version, date string

//  The function Version takes as argument the name of the program  and prints a version string before it exits.
func Version(name string) {
	a := "Bernhard Haubold"
	e := "haubold@evolbio.mpg.de"
	l := "Gnu General Public License, " +
		"https://www.gnu.org/licenses/gpl.html"
	clio.PrintInfo(name, version, date, a, e, l)
	os.Exit(0)
}
