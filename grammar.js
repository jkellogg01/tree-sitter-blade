/**
 * @file a tree-sitter parser for the blade templating language
 * @author Joshua Kellogg
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "blade",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
