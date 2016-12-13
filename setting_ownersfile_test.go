package main

import (
	"testing"
)

func TestOwnersFileToRepoInfo1(t *testing.T) {
	o := OwnersFile{
		IsDisabledMergeAutomatically: false,
		ShouldDeleteMerged:           true,
	}

	ok, info := o.ToRepoInfo()
	if !ok {
		t.Fatal("should be success to convert from OwnersFile")
	}

	if !info.ShouldMergeAutomatically {
		t.Fatal("ShouldMergeAutomatically: should be true")
	}

	if !info.ShouldDeleteMerged {
		t.Fatal("ShouldDeleteMerged: should be true")
	}
}

func TestOwnersFileToRepoInfo2(t *testing.T) {
	o := OwnersFile{
		IsDisabledMergeAutomatically: true,
		ShouldDeleteMerged:           true,
	}

	ok, info := o.ToRepoInfo()
	if !ok {
		t.Fatal("should be success to convert from OwnersFile")
	}

	if info.ShouldMergeAutomatically {
		t.Fatal("ShouldMergeAutomatically: should be false")
	}

	if !info.ShouldDeleteMerged {
		t.Fatal("ShouldDeleteMerged: should be true")
	}
}