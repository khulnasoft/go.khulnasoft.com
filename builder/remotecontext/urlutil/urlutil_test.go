package urlutil // import "go.khulnasoft.com/builder/remotecontext/urlutil"

import "testing"

var (
	gitUrls = []string{
		"git://go.khulnasoft.com",
		"git@github.com:docker/docker.git",
		"git@bitbucket.org:atlassianlabs/atlassian-docker.git",
		"https://go.khulnasoft.com.git",
		"http://go.khulnasoft.com.git",
		"http://go.khulnasoft.com.git#branch",
		"http://go.khulnasoft.com.git#:dir",
	}
	incompleteGitUrls = []string{
		"go.khulnasoft.com",
	}
	invalidGitUrls = []string{
		"http://go.khulnasoft.com.git:#branch",
		"https://github.com/docker/dgit",
	}
)

func TestIsGIT(t *testing.T) {
	for _, url := range gitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range incompleteGitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range invalidGitUrls {
		if IsGitURL(url) {
			t.Fatalf("%q should not be detected as valid Git prefix", url)
		}
	}
}
