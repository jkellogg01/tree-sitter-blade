package tree_sitter_blade_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_blade "github.com/jkellogg01/tree-sitter-blade/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_blade.Language())
	if language == nil {
		t.Errorf("Error loading Blade grammar")
	}
}
