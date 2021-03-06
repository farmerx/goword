// Copyright 2021 - 2021 The goword Authors. All rights reserved. Use of
// this source code is governed by a MIT license that can be found in
// the LICENSE file.
//
// Package goword providing a set of functions that allow you to write to
// and read from DOCX files. Supports reading and writing
// wordprocessing documents generated by Microsoft Word™ 2007 and later.
// This library needs Go version 1.15 or later.

package goword

import "encoding/xml"

// Properties specifies to an OOXML document properties such as the
// template used, the number of pages and words, and the application name and
// version.
type Properties struct {
	XMLName              xml.Name `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Properties"`
	Template             string
	Manager              string
	Company              string
	Pages                int
	Words                int
	Characters           int
	PresentationFormat   string
	Lines                int
	Paragraphs           int
	Slides               int
	Notes                int
	TotalTime            int
	HiddenSlides         int
	MMClips              int
	ScaleCrop            bool
	HeadingPairs         *vectorVariant
	TitlesOfParts        *vectorLpstr
	LinksUpToDate        bool
	CharactersWithSpaces int
	SharedDoc            bool
	HyperlinkBase        string
	HLinks               *vectorVariant
	HyperlinksChanged    bool
	DigSig               *digSig
	Application          string
	AppVersion           string
	DocSecurity          int
}

// vectorVariant specifies the set of hyperlinks that were in this
// document when last saved.
type vectorVariant struct {
	Content string `xml:",innerxml"`
}

// vectorLpstr ...
type vectorLpstr struct {
	Content string `xml:",innerxml"`
}

// digSig contains the signature of a digitally signed document.
type digSig struct {
	Content string `xml:",innerxml"`
}
