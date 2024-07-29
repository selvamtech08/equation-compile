package compile

import (
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
)

// plain tex document lines
const texSource = `\documentclass[class=minimal,preview,varwidth,border=5pt]{standalone}
\usepackage{amsmath,amssymb}
\usepackage{amsfonts,amsthm}

[[PREAM]]

\begin{document}

[[EQU]]

\end{document}`

// save the equation as temporally file
func SaveAsFile(input string, preamble string, mode string) (string, error) {
	randNumber := "./data/" + strconv.Itoa(rand.Intn(100))
	os.MkdirAll(randNumber, 0666)
	fileName := path.Join(randNumber, "equation")
	file, err := os.Create(fileName + ".tex")
	if err != nil {
		return "", err
	}
	defer file.Close()
	equData := strings.Replace(texSource, "[[PREAM]]", preamble, 1)
	if strings.ToLower(mode) == "on" {
		equData = strings.Replace(equData, "[[EQU]]", `$ `+input+` $`, 1)
	} else {
		equData = strings.Replace(equData, "[[EQU]]", `$$ `+input+` $$`, 1)
	}
	_, err = file.WriteString(equData)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
