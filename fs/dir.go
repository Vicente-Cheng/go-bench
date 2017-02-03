package dir

import (
	"os"
	"time"
)

// In some cases, umask would affect the folder permissions.
// We should reset the umask to ensure our permissions.
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
func Dir_Traverse(dirpath string) error {
	return nil
}
