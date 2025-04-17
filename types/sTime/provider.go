package sTime

import (
	"fmt"
	"time"
	_ "time/tzdata" // Importar para embeber la base de datos de zonas horarias
)

type Provider struct {
	comma rune

	layout      string
	layout_date string
	layout_time string

	location  *time.Location
	time_diff time.Duration
}

func NewProvider() *Provider {
	return &Provider{
		location:  time.UTC,
		time_diff: 0,
	}
}

func (p *Provider) SetComma(comma rune) {
	p.comma = comma
}

func (p *Provider) SetLayout(layout string) {
	p.layout = layout
}

func (p *Provider) SetLayoutDate(layout string) {
	p.layout_date = layout
}

func (p *Provider) SetLayoutTime(layout string) {
	p.layout_time = layout
}

func (p *Provider) SetLocation(name string) error {
	// Cargar la ubicación de Nueva York (soporta cambios automáticos entre EST y EDT)
	location, err := time.LoadLocation(name)
	if err != nil {
		return fmt.Errorf("time.LoadLocation(): %s", err)
	}

	p.location = location

	return nil
}

func (p *Provider) SetTimeDiff(diff time.Duration) {
	p.time_diff = diff
}

func (p *Provider) GetUTC(str string) (int64, error) {

	t, err := time.ParseInLocation(p.layout, str, p.location)
	if err != nil {
		return 0, fmt.Errorf("time.ParseInLocation( %s )", str)
	}

	return t.Add(p.time_diff).UTC().Unix(), nil
}

func (p *Provider) Comma() rune {
	return p.comma
}

func (p *Provider) LayoutDate() string {
	return p.layout_date
}

func (p *Provider) LayoutTime() string {
	return p.layout_time
}
