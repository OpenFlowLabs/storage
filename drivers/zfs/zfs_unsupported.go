// +build !linux,!freebsd,!illumos

package zfs

func checkRootdirFs(rootdir string) error {
	return nil
}

func getMountpoint(id string) string {
	return id
}
