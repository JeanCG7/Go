package protein

import (
	"errors"
	"fmt"
)

var (
	ErrStop        = errors.New("need to terminate the translation")
	ErrInvalidBase = errors.New("invalid codon")
	proteins       = map[interface{}][]string{
		"Methionine":    {"AUG"},
		"Tryptophan":    {"UGG"},
		"Cysteine":      {"UGU", "UGC"},
		"Phenylalanine": {"UUU", "UUC"},
		"Leucine":       {"UUA", "UUG"},
		"Tyrosine":      {"UAU", "UAC"},
		"Serine":        {"UCU", "UCC", "UCA", "UCG"},
		ErrStop:         {"UAA", "UAG", "UGA"},
		ErrInvalidBase:  {"ABC"},
	}
)

func FromCodon(codonInput string) (string, error) {
	for protein, codons := range proteins {
		for _, codon := range codons {

			if codon == codonInput {
				if protein == ErrStop {
					return "", ErrStop
				}

				if protein == ErrInvalidBase {
					return "", ErrInvalidBase
				}

				return fmt.Sprintf("%v", protein), nil
			}
		}
	}
	return "", nil
}

func FromRNA(rna string) ([]string, error) {
	proteinsResult := make([]string, 0)
	codonsInput := splitString(rna, 3)
	for i := 0; i < len(codonsInput); i++ {
		if len(codonsInput[i]) != 3 {
			return nil, ErrInvalidBase
		}

		for protein, codons := range proteins {
			for _, codon := range codons {
				if codon == codonsInput[i] {
					if protein == ErrStop {
						return nil, ErrStop
					}

					proteinsResult = append(proteinsResult, fmt.Sprintf("%s", protein))
				}
			}
		}
	}
	return proteinsResult, nil
}

func splitString(word string, quantityPerIndex int) []string {
	runes := []rune(word)
	result := make([]string, 0)
	tmpString := ""
	for i := range runes {
		if i%quantityPerIndex == 0 && i != 0 {
			result = append(result, tmpString)
			tmpString = ""

			if len(runes) == i {
				return result
			}
		}

		tmpString += string(runes[i])
	}
	result = append(result, tmpString)
	return result
}
