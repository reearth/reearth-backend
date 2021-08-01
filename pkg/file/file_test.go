package file

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReaders(t *testing.T) {
	zf, err := os.Open("testdata/test.zip")
	assert.NoError(t, err)
	defer func() {
		_ = zf.Close()
	}()
	zr, err := ZipReaderFrom(zf, 1024)
	assert.NoError(t, err)

	tf, err := os.Open("testdata/test.tar.gz")
	assert.NoError(t, err)
	defer func() {
		_ = tf.Close()
	}()
	tr, err := TarReaderFromTarGz(tf)
	assert.NoError(t, err)

	files := map[string]string{
		"reearth.json": "{\n  \"reearth\": \"Re:Earth\"\n}\n",
		"index.js":     "console.log(\"hello world\");\n",
		"test/foo.bar": "test\n",
	}

	testCases := []struct {
		Name    string
		Archive Archive
		Files   []string
	}{
		{
			Name:    "zip",
			Archive: zr,
			Files:   []string{"test/foo.bar", "index.js", "reearth.json"},
		},
		{
			Name:    "tar",
			Archive: tr,
			Files:   []string{"test/foo.bar", "index.js", "reearth.json"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(tt *testing.T) {
			// tt.Parallel() cannot be used
			assert := assert.New(tt)

			for i, f := range tc.Files {
				n, err := tc.Archive.Next()
				assert.NoError(err)
				assert.Equal(f, n.Path, "file %d in %s", i, tc.Name)
				assert.Equal(int64(len(files[f])), n.Size, "file %d in %s", i, tc.Name)
				assert.Equal("", n.ContentType, "file %d in %s", i, tc.Name)

				fc, err := io.ReadAll(n.Content)
				assert.NoError(err)
				assert.Equal(files[f], string(fc))

				assert.NoError(n.Content.Close())
			}

			n, err := tc.Archive.Next()
			assert.Nil(err)
			assert.Nil(n)

			n, err = tc.Archive.Next()
			assert.Nil(err)
			assert.Nil(n)
		})
	}
}
