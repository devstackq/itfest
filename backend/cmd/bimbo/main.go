package main

import (
	"log"
	"time"

	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

// app.Start()
/*
	plans:
	mind map - tasks
	db - models Db

	auth
	search
	download/
	upload -documents -> db -> history

	crud - document - pattern froms ? -> history
	access document -> history
	backup db
	crypto - docs
	pdf ->
	history -> douments(user)

	history -> uplaodd -> open access -> signature

	1:
*/

func main() {
	c := creator.New()
	c.SetPageMargins(50, 50, 100, 70)

	helvetica, _ := model.NewStandard14Font("Helvetica")
	helveticaBold, _ := model.NewStandard14Font("Helvetica-Bold")

	p := c.NewParagraph("UniDoc")
	p.SetFont(helvetica)
	p.SetFontSize(48)
	p.SetMargins(15, 0, 150, 0)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(p)

	p = c.NewParagraph("Example Page")
	p.SetFont(helveticaBold)
	p.SetFontSize(30)
	p.SetMargins(15, 0, 0, 0)
	p.SetColor(creator.ColorRGBFrom8bit(45, 148, 215))
	c.Draw(p)

	t := time.Now().UTC()
	dateStr := t.Format("1 Jan, 2006 15:04")

	p = c.NewParagraph(dateStr)
	p.SetFont(helveticaBold)
	p.SetFontSize(12)
	p.SetMargins(15, 0, 5, 60)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(p)

	loremTxt := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt" +
		"ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
		"aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore" +
		"eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt " +
		"mollit anim id est laborum."

	p = c.NewParagraph(loremTxt)
	p.SetFontSize(16)
	p.SetColor(creator.ColorBlack)
	p.SetLineHeight(1.5)
	p.SetMargins(0, 0, 5, 0)
	p.SetTextAlignment(creator.TextAlignmentJustify)
	c.Draw(p)

	err := c.WriteToFile("report.pdf")
	if err != nil {
		log.Println("Write file error:", err)
	}
}
