package handler

import (
	"app/internal"
	"path"
)

func GetControllerPathCreator(base string) func(string) string {
	return func(subpath string) string {
		return path.Join(internal.APIBasePath, base, subpath)
	}
}
