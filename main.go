package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type Spell struct {
	Slug          string `json:"slug,omitempty"`
	Name          string `json:"name,omitempty"`
	Desc          string `json:"desc,omitempty"`
	HigherLevel   string `json:"higher_level,omitempty"`
	Range         string `json:"range,omitempty"`
	Components    string `json:"components,omitempty"`
	Material      string `json:"material,omitempty"`
	Ritual        string `json:"ritual,omitempty"`
	Duration      string `json:"duration,omitempty"`
	Concentration string `json:"concentration,omitempty"`
	CastingTime   string `json:"casting_time,omitempty"`
	LevelInt      int    `json:"level_int,omitempty"`
	AttackType    string `json:"attack_type,omitempty"`
}

func (s Spell) Print() {
	var titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("3")).
		// Background(lipgloss.Color("0")).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingLeft(4).
		PaddingRight(4).
		Width(22)
	var descStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("0")).
		Width(60).
		Bold(true)

	r, _ := regexp.Compile("\\d+")
	m := r.FindString(s.Range)
	rangeFeet, _ := strconv.Atoi(m)
	rangeStr := fmt.Sprintf("%s", s.Range)
	if rangeFeet != 0 {
		rangeM := float32(rangeFeet) * 0.3

		rangeStr = fmt.Sprintf("%s/ %.0f m", s.Range, rangeM)
	} else {
	}
	infoStr := fmt.Sprintf("Level: %d\nRange: %s\nComponents: %s\nDuration: %s\nConcentration: %s\nRitual: %s\nCastingTime: %s", s.LevelInt, rangeStr, s.Components, s.Duration, s.Concentration, s.Ritual, s.CastingTime)

	var subStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("6")).Bold(true)

	spellStr := fmt.Sprintf("%s\n%s\n%s\n\n%s\n%s", titleStyle.Render(s.Name), subStyle.Render("Basic Info"), descStyle.Render(infoStr), subStyle.Render("Description"), descStyle.Render(s.Desc))
	fmt.Println(spellStr)
}
func main() {

	l := len(os.Args)
	if l == 1 {
		println("err")
	} else {
		f, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}

		var spell Spell
		json.Unmarshal(f, &spell)

		spell.Print()
	}
}
