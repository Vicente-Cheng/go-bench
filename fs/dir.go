package fs

import (
	"os"
	"time"
)

//
func Dir_Create(dirpath string) (error, time.Duration) {
	t_start := time.Now()
	err := os.MkdirAll(dirpath, 0777)
	t_end := time.Now()
	t_dur := t_end.Sub(t_start)
	return err, t_dur
}

//
func Dir_Remove(dirpath string) (error, time.Duration) {
	t_start := time.Now()
	err := os.RemoveAll(dirpath)
	t_end := time.Now()
	t_dur := t_end.Sub(t_start)
	return err, t_dur
}

//
func Dir_Change(dirpath string) (error, time.Duration) {
	t_start := time.Now()
	err := os.Chdir(dirpath)
	t_end := time.Now()
	t_dur := t_end.Sub(t_start)
	return err, t_dur
}
