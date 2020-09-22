package slugify

import (
	"testing"
)

func TestDanish(t *testing.T) {
	s := Danish("Alefarm Brewing A/S vil sende særligt og unikt øl på Nasdaq First North Growth Market Denmark")

	if s != "alefarm-brewing-a-s-vil-sende-saerligt-og-unikt-oel-paa-nasdaq-first-north-growth-market-denmark" {
		t.Errorf("Unexpected slugified version of Danish string")
	}
}
