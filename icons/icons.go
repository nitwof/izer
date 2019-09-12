package icons

import "path/filepath"

//go:generate gomplate -d map=map.json -f map.go.tmpl -o map.go

// Get returns icon by filetype
func Get(filename string) string {
	ext := filepath.Ext(filename)
	if icon, ok := iconsMap[ext]; ok {
		return icon
	}
	return "\uE612"
}
