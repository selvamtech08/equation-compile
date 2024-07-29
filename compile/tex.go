package compile

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image/png"
	"log"
	"os"
	"os/exec"
	"path"
)

const PNGExec = `pdftoppm`

// convert pdf to png and return png as bytes
func convertPNG(dir, fileName string) ([]byte, error) {
	cmd := exec.Command(PNGExec, `-f`, `1`, `-l`, `2`, `-r`, `300`, `-png`, fileName+".pdf", fileName)
	cmd.Dir = dir
	log.Printf("compile: %s\n", cmd.String())
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(path.Join(dir, fileName+"-1.png"))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// compile the tex file and return the png output as base64 encode string
func Run(file string) (string, error) {
	log.Println("file:", file)
	var stdout bytes.Buffer
	dir, fileName := path.Split(file)
	cmd := exec.Command("pdflatex", fileName+".tex")
	cmd.Dir = dir
	cmd.Stdout = &stdout
	log.Printf("compile: %s\n", cmd.String())

	// remove temp file dir once operation completed
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			log.Printf("failed to remove equ dir %s\n", err.Error())
		}
	}()

	if err := cmd.Run(); err != nil {
		return stdout.String(), errors.New("compile error")
	}

	pngbytes, err := convertPNG(dir, fileName)
	if err != nil {
		return "", err
	}

	img, err := png.Decode(bytes.NewReader(pngbytes))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", err
	}

	data := base64.StdEncoding.EncodeToString(buf.Bytes())
	return data, nil
}
