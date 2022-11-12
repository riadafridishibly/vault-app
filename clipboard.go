package main

import "golang.design/x/clipboard"

type ClipboardService struct{}

func NewClipboardService() (*ClipboardService, error) {
	err := clipboard.Init()
	if err != nil {
		return nil, err
	}
	return &ClipboardService{}, nil
}

func (clip *ClipboardService) Copy(value string) {
	clipboard.Write(clipboard.FmtText, []byte(value))
}

func (clip *ClipboardService) Paste() string {
	return string(clipboard.Read(clipboard.FmtText))
}
