package main

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/urfave/cli/v2"
)

func isZipFile(inputFile string) (bool, error) {
	_, err := os.Stat(inputFile)
	if err != nil {
		return false, err
	}

	inputFileDiscriptor, err := os.Open(inputFile)
	if err != nil {
		return false, err
	}
	defer inputFileDiscriptor.Close()

	buff := make([]byte, 512)
	_, err = inputFileDiscriptor.Read(buff)
	if err != nil {
		return false, err
	}

	filetype := http.DetectContentType(buff)
	return filetype == "application/zip", nil
}

func findCoverFile(inputFile string) (*zip.File, error) {
	isZip, err := isZipFile(inputFile)
	if err != nil {
		return nil, err
	}

	if !isZip {
		return nil, errors.New("This is not epub file")
	}

	inputFileReader, err := zip.OpenReader(inputFile)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	for _, file := range inputFileReader.File {
		log.Print(file.Name)
		match, _ := regexp.MatchString(`.+(Cover|cover)(.+)?\.(jpeg|jpg|png)$`, file.Name)
		if match {
			return file, nil
		}
	}

	return nil, errors.New("Epub has no cover")
}

func extractCover(ctx *cli.Context) (bool, error) {
	inputFile := ctx.Args().Get(0)
	outputFile := ctx.Args().Get(1)

	if len(inputFile) == 0 {
		return true, cli.Exit("No input file given", 86)
	}

	if len(outputFile) == 0 {
		return true, cli.Exit("No output file given", 86)
	}

	coverFile, err := findCoverFile(inputFile)
	if err != nil {
		log.Fatal(err, inputFile)
		cli.Exit(err, 1)
	}

	srcFile, err := coverFile.Open()
	if err != nil {
		log.Fatal(err, inputFile)
		cli.Exit(err, 1)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err, inputFile)
		cli.Exit(err, 1)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		cli.Exit(err, 1)
	}
	return false, nil
}

func main() {
	app := &cli.App{
		Name:                 "epub-cover",
		HelpName:             "",
		Usage:                "Extract cover image for epub file",
		UsageText:            "",
		ArgsUsage:            "",
		Version:              "",
		Description:          "",
		DefaultCommand:       "",
		Commands:             []*cli.Command{},
		Flags:                []cli.Flag{},
		EnableBashCompletion: false,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          false,
		BashComplete: func(*cli.Context) {
		},
		Before: func(*cli.Context) error {
			return nil
		},
		After: func(*cli.Context) error {
			return nil
		},
		Action: func(ctx *cli.Context) error {
			_, err := extractCover(ctx)

			return err
		},
		CommandNotFound: func(*cli.Context, string) {
		},
		OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		InvalidFlagAccessHandler: func(*cli.Context, string) {
		},
		Compiled:  time.Time{},
		Authors:   []*cli.Author{},
		Copyright: "",
		Reader:    nil,
		Writer:    nil,
		ErrWriter: nil,
		ExitErrHandler: func(cCtx *cli.Context, err error) {
		},
		Metadata: map[string]interface{}{},
		ExtraInfo: func() map[string]string {
			return nil
		},
		CustomAppHelpTemplate:     "",
		SliceFlagSeparator:        "",
		DisableSliceFlagSeparator: false,
		UseShortOptionHandling:    false,
		Suggest:                   false,
		AllowExtFlags:             false,
		SkipFlagParsing:           false,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
