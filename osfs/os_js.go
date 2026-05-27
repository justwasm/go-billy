//go:build js

package osfs

func (f *file) Lock() error {
	f.m.Lock()
	defer f.m.Unlock()
	return nil
}

func (f *file) Unlock() error {
	f.m.Lock()
	defer f.m.Unlock()
	return nil
}

func (f *file) Sync() error {
	return f.File.Sync()
}

func umask(m int) func() {
	return func() {}
}

func (f *file) Fd() (uintptr, bool) {
	return f.File.Fd(), true
}