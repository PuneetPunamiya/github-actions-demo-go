package hello

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Hello()
	if result != "Hello GitHub Actions. CI Runs yo" {
		t.Errorf("Hello() = %s; want Hello GitHub Actions. CI Runs yo", result)
	}
}
