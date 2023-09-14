package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		file, err := os.CreateTemp(os.TempDir(), "go-template.yaml")
		require.NoError(t, err)
		defer os.Remove(file.Name())

		_, err = file.WriteString("listen_addr: 127.0.0.1:8080")
		require.NoError(t, err)

		err = file.Close()
		require.NoError(t, err)

		conf, err := ReadConfig(file.Name())
		require.NoError(t, err)

		want := Config{ListenAddr: "127.0.0.1:8080"}
		require.Equal(t, want, conf)
	})

	t.Run("empty listen address", func(t *testing.T) {
		file, err := os.CreateTemp(os.TempDir(), "go-template.yaml")
		require.NoError(t, err)
		defer os.Remove(file.Name())

		err = file.Close()
		require.NoError(t, err)

		conf, err := ReadConfig(file.Name())
		require.NoError(t, err)

		want := Config{ListenAddr: "0.0.0.0:8080"}
		require.Equal(t, want, conf)
	})

	t.Run("invalid yaml", func(t *testing.T) {
		file, err := os.CreateTemp(os.TempDir(), "go-template.yaml")
		require.NoError(t, err)
		defer os.Remove(file.Name())

		_, err = file.WriteString("hello: world: !")
		require.NoError(t, err)

		err = file.Close()
		require.NoError(t, err)

		_, err = ReadConfig(file.Name())
		require.ErrorContains(t, err, "unmarshal file")
	})

	t.Run("non-existing config file", func(t *testing.T) {
		_, err := ReadConfig("not-exists.yaml")
		require.ErrorContains(t, err, "no such file or directory")
	})
}
