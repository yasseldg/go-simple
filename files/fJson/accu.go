package fJson

import (
	"encoding/json"
	"fmt"

	"github.com/yasseldg/go-simple/files/fAccu"
	"github.com/yasseldg/go-simple/logs/sLog"
)

type InterAccu interface {
	fAccu.Inter
}

type Accu struct {
	fAccu.Inter

	data func() any
}

func NewAccu(file_path string, delete bool, limit int, data func() any) (*Accu, error) {

	accu := new(Accu)

	inter, err := fAccu.New(file_path, delete, limit, accu.save)
	if err != nil {
		return nil, err
	}

	accu.Inter = inter
	accu.data = data

	return accu, nil
}

func (a *Accu) save() error {
	if a.data == nil {
		return nil
	}

	err := a.write()
	if err != nil {
		return err
	}

	// FIX
	// a.data = nil

	return nil
}

func (a *Accu) write() error {
	data, err := json.MarshalIndent(a.data(), "", " ")
	if err != nil {
		return fmt.Errorf("json.MarshalIndent(objects): %s", err)
	}

	file, err := a.OpenFile()
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := file.Write(data)
	if err != nil {
		return fmt.Errorf("file.Write(data) in ( %s ): %s", a.FilePath(), err)
	}

	bb, err := file.WriteString("\n")
	if err != nil {
		return fmt.Errorf("ExportJson: file.WriteString(): path: %s : %s", a.FilePath(), err)

	}

	sLog.Info("Accu Json: %d bytes written successfully in ( %s )", (b + bb), a.FilePath())
	return nil
}
