package fCsv

import (
	"encoding/csv"
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/files/fAccu"
	"github.com/yasseldg/go-simple/logs/sLog"
)

type Accu struct {
	fAccu.Inter

	mu sync.Mutex

	data [][]string
}

func NewAccu(file_path string, delete bool, limit int) (*Accu, error) {

	accu := new(Accu)

	inter, err := fAccu.New(file_path, delete, limit, accu.save)
	if err != nil {
		return nil, err
	}

	accu.Inter = inter

	return accu, nil
}

func (a *Accu) AddHeader(header []string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.IsNew() {
		a.data = append([][]string{header}, a.data...)

		a.SetNew(false)
	}
}

func (a *Accu) AddData(d []string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.data = append(a.data, d)

	a.Increase()

	if len(a.data) >= a.Limit() {
		a.Save()
	}
}

//  private methods

func (a *Accu) save() error {
	if len(a.data) == 0 {
		return fmt.Errorf("no data to save in ( %s )", a.FilePath())
	}

	err := a.write()
	if err != nil {
		return err
	}

	a.data = [][]string{}

	return nil
}

func (a *Accu) write() error {
	f, err := a.OpenFile()
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	err = writer.WriteAll(a.data)
	if err != nil {
		return fmt.Errorf("csvWriter.WriteAll(data) in ( %s ): %s", a.FilePath(), err)
	}

	sLog.Info("Accu Csv: %d lines written successfully in ( %s )", len(a.data), a.FilePath())
	return nil
}
