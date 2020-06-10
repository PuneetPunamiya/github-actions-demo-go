package hello

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. CI Runs yo" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions. CI Runs yo", result)
	}
}
