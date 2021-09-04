package editor

import (
	"io"
	"os"
	"os/exec"

	"github.com/keesvv/gotag/pkg/parser"
	"github.com/keesvv/gotag/pkg/tagger"
)

type Editor struct {
	Exec string
}

const TMP_PATH = "/tmp/gotag"

func GetPreferred() *Editor {
	editorEnv, _ := os.LookupEnv("EDITOR")

	if editorEnv == "" {
		editorEnv = "vim"
	}

	return &Editor{
		Exec: editorEnv,
	}
}

func (edt *Editor) GetTempBuffer() (*os.File, error) {
	if err := os.Mkdir(TMP_PATH, 0700); !os.IsExist(err) {
		return nil, err
	}

	return os.CreateTemp(TMP_PATH, "buf-*.yml")
}

func (edt *Editor) WriteDefaults(fname string, buf *os.File, tagger *tagger.Tagger) error {
	p := parser.Parser{}
	b, err := p.MarshalTag(fname, tagger.Tag)

	if err != nil {
		return err
	}

	_, err = buf.Write(b)
	return err
}

func (edt *Editor) Edit(buf *os.File) ([]byte, error) {
	cmd := exec.Command(edt.Exec, buf.Name())
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Start()

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	defer buf.Close()
	defer os.Remove(buf.Name())

	buf.Seek(0, 0)

	return io.ReadAll(buf)
}
