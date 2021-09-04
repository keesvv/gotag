package editor

import (
	"io"
	"os"
	"os/exec"

	id3 "github.com/bogem/id3v2"
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

	return os.CreateTemp(TMP_PATH, "buf-*")
}

func (edt *Editor) WriteDefaults(buf *os.File, tag *id3.Tag) error {
	_, err := buf.Write([]byte(tag.Artist()))
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
