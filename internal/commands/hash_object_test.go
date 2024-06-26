package commands_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/Luisgustavom1/build-your-own-git/internal/commands"
	test_utils "github.com/Luisgustavom1/build-your-own-git/internal/commands/tests/utils"
	"github.com/stretchr/testify/require"
)

func TestHashObject(t *testing.T) {
	testCases := []struct {
		name         string
		flag         string
		filePath     string
		expectedHash string
	}{
		{
			name:         "with flag -w",
			flag:         "-w",
			filePath:     "../fixtures/hello-world.txt",
			expectedHash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := test_utils.GitInitSetup(t)
			require.NoError(t, err)

			args := []string{"mygit", "hash-object", tc.flag, tc.filePath}

			res, err := commands.HashObject(args)
			require.NoError(t, err)
			require.Equal(t, fmt.Sprintln(tc.expectedHash), res)

			objectPath := path.Join(".git/objects", tc.expectedHash[:2])
			_, err = os.Stat(objectPath)
			require.False(t, os.IsNotExist(err), fmt.Sprintf("dir %s should be created", objectPath))
		})
	}

	t.Run("invalid arguments", func(t *testing.T) {
		dir, err := os.MkdirTemp("", "test")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(dir)

		args := []string{"mygit", "hash-object"}

		_, err = commands.HashObject(args)
		require.EqualErrorf(t, err, "usage: mygit hash-object <object>\n", "Invalid args")
	})
}
