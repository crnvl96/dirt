package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsGitRepo(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "dirt_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// With this setup we end with the following file structure:
	//
	// dir_test/

	// Test non-git directory
	if isGitRepo(tempDir) {
		t.Error("Expected false for non-git directory")
	}

	// Create .git directory
	gitDir := filepath.Join(tempDir, ".git")
	err = os.Mkdir(gitDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// With this setup we end with the following file structure:
	//
	// dir_test/
	//   .git/

	// Test git directory
	if !isGitRepo(tempDir) {
		t.Error("Expected true for git directory")
	}
}

func TestWalkDir(t *testing.T) {
	// Create a temporary directory structure
	tempDir, err := os.MkdirTemp("", "dirt_test_walk")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create subdirs
	subDir1 := filepath.Join(tempDir, "subdir1")
	err = os.Mkdir(subDir1, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	subDir2 := filepath.Join(subDir1, "subdir2")
	err = os.Mkdir(subDir2, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// Create a git repo in subdir2
	gitDir := filepath.Join(subDir2, ".git")
	err = os.Mkdir(gitDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// With this setup we end with the following file structure:
	//
	// dir_test_walk/
	//   subdir1/
	//     subdir2/
	//       .git/

	var repos []repoStatus
	walkDir(tempDir, 0, 2, &repos)

	// Only subdir2 is a repo

	if len(repos) != 1 {
		t.Errorf("Expected 1 repo, got %d", len(repos))
	}

	if repos[0].Path != subDir2 {
		t.Errorf("Expected path %s, got %s", subDir2, repos[0].Path)
	}
}

func TestWalkDirStopsAtGitRepo(t *testing.T) {
	// Test that walkDir stops searching subdirectories when a git repo is found
	tempDir, err := os.MkdirTemp("", "dirt_test_stop")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create tempSubDir as git repo
	tempSubDir := filepath.Join(tempDir, "tempSubDir")
	err = os.Mkdir(tempSubDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}
	gitDir1 := filepath.Join(tempSubDir, ".git")
	err = os.Mkdir(gitDir1, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// Create tempSubDir2 inside tempSubDir, also as git repo
	tempSubDir2 := filepath.Join(tempSubDir, "tempSubDir2")
	err = os.Mkdir(tempSubDir2, 0o755)
	if err != nil {
		t.Fatal(err)
	}
	gitDir2 := filepath.Join(tempSubDir2, ".git")
	err = os.Mkdir(gitDir2, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// With this setup we end with the following file structure:
	//
	// dir_test_stop/
	//   tempSubDir/
	//     .git/
	//     tempSubDir2/
	//       .git/

	var repos []repoStatus
	walkDir(tempDir, 0, 2, &repos)

	// Should find only tempSubDir, not tempSubDir2, since
	// tempSubDir contains tempSubDir2 and tempSubDir is already a git repo
	if len(repos) != 1 {
		t.Errorf("Expected 1 repo, got %d", len(repos))
	}

	if repos[0].Path != tempSubDir {
		t.Errorf("Expected path %s, got %s", tempSubDir, repos[0].Path)
	}
}

func TestWalkDirFindsNestedGitRepo(t *testing.T) {
	// Test that walkDir finds nested git repo when parent is not git
	tempDir, err := os.MkdirTemp("", "dirt_test_nested")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create tempSubDir, not git
	tempSubDir := filepath.Join(tempDir, "tempSubDir")
	err = os.Mkdir(tempSubDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// Create tempSubDir2 as git repo
	tempSubDir2 := filepath.Join(tempSubDir, "tempSubDir2")
	err = os.Mkdir(tempSubDir2, 0o755)
	if err != nil {
		t.Fatal(err)
	}
	gitDir := filepath.Join(tempSubDir2, ".git")
	err = os.Mkdir(gitDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// With this setup we end with the following file structure:
	//
	// dir_test_nested/
	//   tempSubDir/
	//     tempSubDir2/
	//       .git/

	var repos []repoStatus
	walkDir(tempDir, 0, 2, &repos)

	// Should find tempSubDir2, since tempSubDir is not a git repo
	if len(repos) != 1 {
		t.Errorf("Expected 1 repo, got %d", len(repos))
	}

	if repos[0].Path != tempSubDir2 {
		t.Errorf("Expected path %s, got %s", tempSubDir2, repos[0].Path)
	}
}

func TestGetVCSInfos(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "dirt_test_vcs")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a git repo
	gitDir := filepath.Join(tempDir, ".git")
	err = os.Mkdir(gitDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// With this setup we end with the following file structure:
	//
	// dir_test_vcs/
	//   .git/

	repos := getVCSInfos([]string{tempDir})

	if len(repos) != 1 {
		t.Errorf("Expected 1 repo, got %d", len(repos))
	}

	if repos[0].Path != tempDir {
		t.Errorf("Expected path %s, got %s", tempDir, repos[0].Path)
	}
}
