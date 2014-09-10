package the_platinum_searcher

import (
	"testing"
)

func TestIgnorePatterns(t *testing.T) {

	patterns := IgnorePatterns("files/ignore", []string{"ignore.txt"}, -1)

	if !patterns[0].Match("pattern1", 0) {
		t.Errorf("It should be match %s", "pattern1")
	}
	if !patterns[0].Match("pattern2", 0) {
		t.Errorf("It should be match %s", "pattern2")
	}
}

func TestGitIgnoreMatcher(t *testing.T) {
	ignorePattern := "/ignoreme.txt"
	gitMatcher := gitIgnoreMatcher{[]string{ignorePattern}, 1}

	ignoreThis := "ignoreme.txt"
	depth := 2
	if gitMatcher.Match(ignoreThis, depth) {
		t.Errorf(
			"Git ignore pattern \"%s\" should not match \"%s\" on level %d",
			ignorePattern,
			ignoreThis,
			depth,
		)
	}

	depth = 0
	if gitMatcher.Match(ignoreThis, depth) {
		t.Errorf(
			"Git ignore pattern \"%s\" should not match \"%s\" on level %d",
			ignorePattern,
			ignoreThis,
			depth,
		)
	}

	depth = 1
	if !gitMatcher.Match(ignoreThis, depth) {
		t.Errorf(
			"Git ignore pattern \"%s\" should match \"%s\" on level %d",
			ignorePattern,
			ignoreThis,
			depth,
		)
	}
}

func TestGitIgnoreDir(t *testing.T) {
	gitMatcher := gitIgnoreMatcher{[]string{"node_modules"}, 1}

	if !gitMatcher.Match("node_modules/", 1) {
		t.Error(`Git ignore pattern: "node_modules" should match dir: "node_modules/"`)
	}

	gitMatcher = gitIgnoreMatcher{[]string{"node_modules/"}, 1}
	if gitMatcher.Match("node_modules", 1) {
		t.Error(`Git ignore pattern: "node_modules/" should not match file: "node_modules"`)
	}

	if !gitMatcher.Match("node_modules/", 1) {
		t.Error(`Git ignore pattern: "node_modules/" should match dir: "node_modules/"`)
	}
}