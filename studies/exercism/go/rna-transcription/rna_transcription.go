// Package strand contains simple functions to handle DNA and RNA strands
package strand

// ToRNA returns an RNA strand given an DNA strand
func ToRNA(dna string) string {
	var rna string

	for i := range dna {
		switch dna[i] {
		case 'G':
			rna = rna + "C"
		case 'C':
			rna = rna + "G"
		case 'T':
			rna = rna + "A"
		case 'A':
			rna = rna + "U"
		}
	}

	return rna
}
