package filemode

import (
	"fmt"
	"os"
)

type Mode uint32

func (m Mode) String() string {

	v := make([]byte, 10)

	// Type
	switch t := m & S_IFMT; t {
	case 0:
		v[0] = '-'
	case S_IFLNK:
		v[0] = 'l'
	case S_IFREG:
		v[0] = 'r'
	case S_IFDIR:
		v[0] = 'd'
	case S_IFCHR:
		v[0] = 'c'
	case S_IFBLK:
		v[0] = 'b'
	case S_IFIFO:
		v[0] = 'f'
	case S_IFSOCK:
		v[0] = 's'
	default:
		v[0] = '?'
	}

	// User
	if IsSet(m, ReadUser) {
		v[1] = 'r'
	} else {
		v[1] = '-'
	}

	if IsSet(m, WriteUser) {
		v[2] = 'w'
	} else {
		v[2] = '-'
	}

	if IsSet(m, ReadUser) {
		v[3] = 'x'
	} else {
		v[3] = '-'
	}

	// Group
	if IsSet(m, ReadGroup) {
		v[4] = 'r'
	} else {
		v[4] = '-'
	}

	if IsSet(m, WriteGroup) {
		v[5] = 'w'
	} else {
		v[5] = '-'
	}

	if IsSet(m, ExecGroup) {
		v[6] = 'x'
	} else {
		v[6] = '-'
	}

	// Other
	if IsSet(m, ReadOther) {
		v[7] = 'r'
	} else {
		v[7] = '-'
	}

	if IsSet(m, WriteOther) {
		v[8] = 'w'
	} else {
		v[8] = '-'
	}

	if IsSet(m, ExecOther) {
		v[9] = 'x'
	} else {
		v[9] = '-'
	}

	return string(v)
}

// From linux/stat.h
// See more: https://www.gnu.org/software/libc/manual/html_node/Permission-Bits.html
// See more: https://www.gnu.org/software/libc/manual/html_node/Testing-File-Type.html

const (
	S_IRUSR Mode = 0x00400 // Read permission bit for the owner of the file.
	S_IWUSR Mode = 0x00200 // Write permission bit for the owner of the file.
	S_IXUSR Mode = 0x00100 // Execute (for ordinary files) or search (for directories) permission bit for the owner of the file.
	S_IRWXU Mode = 0x00700 // This is equivalent to (S_IRUSR | S_IWUSR | S_IXUSR).

	S_IRGRP Mode = 0x00040 // Read permission bit for the group owner of the file.
	S_IWGRP Mode = 0x00020 // Write permission bit for the group owner of the file.
	S_IXGRP Mode = 0x00010 // Execute or search permission bit for the group owner of the file.
	S_IRWXG Mode = 0x00070 // This is equivalent to (S_IRGRP | S_IWGRP | S_IXGRP).

	S_IROTH Mode = 0x00004 // Read permission bit for other users.
	S_IWOTH Mode = 0x00002 // Write permission bit for other users.
	S_IXOTH Mode = 0x00001 // Execute or search permission bit for other users.
	S_IRWXO Mode = 0x00007 // This is equivalent to (S_IROTH | S_IWOTH | S_IXOTH).

	S_IFMT   Mode = 00170000 // This is a bit mask used to extract the file type code from a mode value.
	S_IFSOCK Mode = 0140000  // This is the file type constant of a socket.
	S_IFLNK  Mode = 0120000  // This is the file type constant of a symbolic link.
	S_IFREG  Mode = 0100000  // This is the file type constant of a regular file.
	S_IFBLK  Mode = 0060000  // This is the file type constant of a block-oriented device file.
	S_IFDIR  Mode = 0040000  // This is the file type constant of a directory file.
	S_IFCHR  Mode = 0020000  // This is the file type constant of a character-oriented device file.
	S_IFIFO  Mode = 0010000  // This is the file type constant of a FIFO or pipe.
	S_ISUID  Mode = 0004000  // This is the set-user-ID on execute bit.
	S_ISGID  Mode = 0002000  // This is the set-group-ID on execute bit.
	S_ISVTX  Mode = 0001000  // This is the sticky bit.

	// Easy to use aliases

	ReadUser  Mode = S_IRUSR // Read user
	WriteUser Mode = S_IWUSR // Write user
	ExecUser  Mode = S_IXUSR // Execute user

	ReadGroup  Mode = S_IRGRP // Read group
	WriteGroup Mode = S_IWGRP // Write group
	ExecGroup  Mode = S_IXGRP // Execute group

	ReadOther  Mode = S_IROTH // Read other
	WriteOther Mode = S_IWOTH // Write other
	ExecOther  Mode = S_IXOTH // Execute other
)

// Set sets m bit in mode.
func Set(mode, m Mode) Mode {

	return mode | m

}

// Unset unsets m bit in mode.
func Unset(mode, m Mode) Mode {

	return mode ^ (mode & m)

}

// IsSet check whether m bit is set in mode.
func IsSet(mode, m Mode) bool {

	return mode&m == m

}

// IsLnk returns whether the file is a symbolic link.
func IsLnk(m Mode) bool {

	// S_ISLNK(m) (((m) & S_IFMT) == S_IFLNK)

	return m&S_IFMT == S_IFLNK
}

// IsReg returns whether the file is a regular file.
func IsReg(m Mode) bool {

	// S_ISREG(m) (((m) & S_IFMT) == S_IFREG)

	return m&S_IFMT == S_IFREG
}

// IsDir returns whether the file is a directory.
func IsDir(m Mode) bool {

	// S_ISDIR(m) (((m) & S_IFMT) == S_IFDIR)

	return m&S_IFMT == S_IFDIR
}

// IsChr returns whether the file is a character special file (a device like a terminal).
func IsChr(m Mode) bool {

	// S_ISCHR(m) (((m) & S_IFMT) == S_IFCHR)

	return m&S_IFMT == S_IFCHR
}

// IsBlk returns whether the file is a block special file (a device like a disk).
func IsBlk(m Mode) bool {

	// S_ISBLK(m) (((m) & S_IFMT) == S_IFBLK)

	return m&S_IFMT == S_IFBLK
}

// IsFifo returns whether the file is a FIFO special file, or a pipe.
func IsFifo(m Mode) bool {

	// S_ISFIFO(m) (((m) & S_IFMT) == S_IFIFO)

	return m&S_IFMT == S_IFIFO
}

// IsSock returns  whether the file is a socket.
func IsSock(m Mode) bool {

	// S_ISSOCK(m) (((m) & S_IFMT) == S_IFSOCK)

	return m&S_IFMT == S_IFSOCK
}

// GetFile returns the file mode of file.
func GetFile(file *os.File) (Mode, error) {

	s, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf("failed to get stat: %w", err)
	}

	return Mode(s.Mode()), nil
}

// SetFile sets the mode bit m in file.
func SetFile(file *os.File, m Mode) error {

	mode, err := GetFile(file)
	if err != nil {
		return fmt.Errorf("failed to get mode: %w", err)
	}

	return file.Chmod(os.FileMode(Set(m, mode)))
}

// UnsetFile unsets the mode bit m in file.
func UnsetFile(file *os.File, m Mode) error {

	mode, err := GetFile(file)
	if err != nil {
		return fmt.Errorf("failed to get mode: %w", err)
	}

	return file.Chmod(os.FileMode(Unset(mode, m)))
}

// IsSetFile checks whether the mode bit m is set in file.
func IsSetFile(file *os.File, m Mode) (bool, error) {

	mode, err := GetFile(file)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return IsSet(mode, m), nil
}

// IsLnkFile returns whether the file is a symbolic link.
func IsLnkFile(file *os.File) (bool, error) {

	m, err := GetFile(file)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFLNK, nil
}

// IsRegFile returns whether the file is a regular file.
func IsRegFile(file *os.File) (bool, error) {

	m, err := GetFile(file)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFREG, nil
}

// IsDirFile returns whether the file is a directory.
func IsDirFile(file *os.File) (bool, error) {

	m, err := GetFile(file)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFDIR, nil
}

// IsChrFile returns whether the file is a character special file (a device like a terminal).
func IsChrFile(file *os.File) (bool, error) {

	m, err := GetFile(file)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFCHR, nil
}

// IsBlkFile returns whether the file is a block special file (a device like a disk).
func IsBlkFile(file *os.File) (bool, error) {

	m, err := GetFile(file)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFBLK, nil
}

// IsFifoFile returns whether the file is a FIFO special file, or a pipe.
func IsFifoFile(file *os.File) (bool, error) {

	m, err := GetFile(file)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFIFO, nil
}

// IsSockFile returns  whether the file is a socket.
func IsSockFile(file *os.File) (bool, error) {

	m, err := GetFile(file)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFSOCK, nil
}

// GetPath returns the mode of the file in path.
// This function (and every *Path functions) does not follows symbolic link.
func GetPath(path string) (Mode, error) {

	stat, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("failed to stat: %w", err)
	}

	return Mode(stat.Mode()), nil
}

// SetPath sets the mode bit m for file in path.
func SetPath(path string, m Mode) error {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

	return SetFile(file, m)
}

// UnsetPath unsets the mode bit m for file in path.
func UnsetPath(path string, m Mode) error {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

	return UnsetFile(file, m)
}

// GetPath returns the mode of the file in path.
// This function does not follow symbolic link.
func IsSetPath(path string, m Mode) (bool, error) {

	stat, err := os.Lstat(path)
	if err != nil {
		return false, fmt.Errorf("failed to stat: %w", err)
	}

	return IsSet(Mode(stat.Mode()), m), nil
}

// IsLnkPath returns whether the file in path is a symbolic link.
func IsLnkPath(path string) (bool, error) {

	m, err := GetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFLNK, nil
}

// IsRegPath returns whether the file in path is a regular file.
func IsRegPath(path string) (bool, error) {

	m, err := GetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFREG, nil
}

// IsDirPath returns whether the file in path is a directory.
func IsDirPath(path string) (bool, error) {

	m, err := GetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFDIR, nil
}

// IsChrPath returns whether the file in path is a character special file (a device like a terminal).
func IsChrPath(path string) (bool, error) {

	m, err := GetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFCHR, nil
}

// IsBlkPath returns whether the file in path is a block special file (a device like a disk).
func IsBlkPath(path string) (bool, error) {

	m, err := GetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFBLK, nil
}

// IsFifoPath returns whether the file in path is a FIFO special file, or a pipe.
func IsFifoPath(path string) (bool, error) {

	m, err := GetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFIFO, nil
}

// IsSockPath returns whether the file in path is a socket.
func IsSockPath(path string) (bool, error) {

	m, err := GetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFSOCK, nil
}
