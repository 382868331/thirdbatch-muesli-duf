package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixDuf001SourceContract(t *testing.T) {
    source, err := os.ReadFile("mounts_linux.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if !sawSep && len(all) > mountinfoOptionalFields {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if !sawSep || len(all) > mountinfoOptionalFields {") {
        t.Fatalf("mutated source contract is still present")
    }
}
