package slugify

import (
	"testing"
)

func TestEnglish(t *testing.T) {
	s := Danish("Alefarm Brewing A/S is intending to introduce unique and flavorful beers to the Nasdaq First North Growth Market Denmark")

	if s != "alefarm-brewing-a-s-is-intending-to-introduce-unique-and-flavorful-beers-to-the-nasdaq-first-north-growth-market-denmark" {
		t.Errorf("Unexpected slugified version of English string")
	}
}
