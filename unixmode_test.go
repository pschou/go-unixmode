package unixmode_test

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/pschou/go-unixmode"
)

func ExampleMode_String() {
	m := unixmode.Mode(041755)
	fmt.Printf("mode: %q\n", m.String())

	// Output:
	// mode: "drwxr-xr-t "
}

func ExampleMode_IsDir() {
	m, _ := unixmode.Parse("dr-s-wSr-T")
	fmt.Printf("is dir: %v\n", m.IsDir())

	// Output:
	// is dir: true
}

func ExampleMode_IsRegular() {
	m, _ := unixmode.Parse("-r-s-wSr-T")
	fmt.Printf("is regular: %v\n", m.IsRegular())

	// Output:
	// is regular: true
}

func ExampleMode_Perm_parse() {
	m, _ := unixmode.Parse("r-s-wSr-T")
	fmt.Printf("mode: %04o\n", m.Perm())

	// Output:
	// mode: 7524
}

func ExampleMode_Perm() {
	m := unixmode.Mode(041755)
	fmt.Printf("mode: %04o\n", m.Perm())

	// Output:
	// mode: 1755
}

func ExampleMode_PermString() {
	m := unixmode.Mode(0755)
	fmt.Printf("mode: %q\n", m.PermString())

	// Output:
	// mode: "rwxr-xr-x"
}

func ExampleMode_PermString_directory() {
	m := unixmode.Mode(040755)
	fmt.Printf("mode: %q\n", m.PermString())

	// Output:
	// mode: "rwxr-xr-x"
}

func ExampleMode_PermString_setUidBits() {
	m := unixmode.Mode(06755)
	fmt.Printf("mode: %q\n", m.PermString())

	// Output:
	// mode: "rwsr-sr-x"
}

func ExampleFileModeString() {
	stat, _ := os.Lstat("/tmp")
	m := stat.Mode()
	fmt.Printf("mode: %q\n", unixmode.FileModeString(m))

	// Output:
	// mode: "drwxrwxrwt "
}

func ExampleFileModePermString() {
	stat, _ := os.Lstat("/tmp")
	m := stat.Mode()
	fmt.Printf("mode: %q\n", unixmode.FileModePermString(m))

	// Output:
	// mode: "rwxrwxrwt"
}

func ExampleFileModeString_socket() {
	if stat, err := os.Lstat("/dev/log"); err != nil {
		log.Fatal("Could not stat socket", err)
	} else {
		m := stat.Mode()
		fmt.Printf("mode: %q\n", unixmode.FileModeString(m))
	}

	// Output:
	// mode: "srw-rw-rw- "
}

func ExampleMode_Chmod() {
	m := unixmode.Mode(02644 | unixmode.ModeRegular)
	fmt.Printf("mode: %q\n", m.PermString())
	unixmode.Chmod("t", m)

	// Output:
	// mode: "rw-r-Sr--"
}

func ExampleFileModePerm_sUDO() {
	if stat, err := os.Lstat("/usr/bin/sudo"); err != nil {
		log.Fatal("Could not stat character device", err)
	} else {
		m := stat.Mode()
		fmt.Printf("mode: %c %04o\n", unixmode.FileModeTypeLetter(m), unixmode.FileModePerm(m))
	}

	// Output:
	// mode: - 4111
}

func ExampleFileModeTypeLetter() {
	if stat, err := os.Lstat("/dev/null"); err != nil {
		log.Fatal("Could not stat character device", err)
	} else {
		m := stat.Mode()
		fmt.Printf("mode: %c\n", unixmode.FileModeTypeLetter(m))
	}

	// Output:
	// mode: c
}

func ExampleFileModePerm_rawCharacter() {
	if stat, err := os.Lstat("/dev/lp0"); err != nil {
		log.Fatal("Could not stat character device", err)
	} else {
		m := stat.Mode()
		fmt.Printf("mode: %c %04o\n", unixmode.FileModeTypeLetter(m), unixmode.FileModePerm(m))
	}

	// Output:
	// mode: c 0660
}

func ExampleFileModeString_character() {
	if stat, err := os.Lstat("/dev/lp0"); err != nil {
		log.Fatal("Could not stat character device", err)
	} else {
		m := stat.Mode()
		fmt.Printf("mode: %q\n", unixmode.FileModeString(m))
	}

	// Output:
	// mode: "crw-rw---- "
}

func ExampleParse() {
	if m, err := unixmode.Parse("rwxrwx---"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("mode: %04o\n", m)
	}
	// Output:
	// mode: 0770
}

func ExampleParse_directoryMixedRaw() {
	if m, err := unixmode.Parse("d-w-r-S-wT"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("mode: %04o\n", m)
	}
	// Output:
	// mode: 43242
}

func ExampleParse_mixed() {
	if m, err := unixmode.Parse("-w-r-S-wT"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("mode: %04o\n", m.Perm())
	}
	// Output:
	// mode: 3242
}

func ExampleParse_invalid() {
	if m, err := unixmode.Parse("drwSrwSrwS "); err != nil {
		fmt.Println("Err:", err)
	} else {
		fmt.Printf("mode: %04o\n", m)
	}
	// Output:
	// Err: Invalid 'S' at position 9
}

func ExampleNew() {
	fm := fs.FileMode(0777 | fs.ModeSetuid)
	fmt.Printf("Full unix mode: %07o\n", unixmode.New(fm))
	fmt.Printf("Unix perms: %05o\n", unixmode.New(fm).Perm())
	// Output:
	// Full unix mode: 0104777
	// Unix perms: 04777
}

func ExampleMode_FileMode() {
	fm, _ := unixmode.Parse("dr-sr-srwx")
	fmt.Printf("Unix mode: %q\n", fm.PermString())
	fmt.Printf("Go FileMode: %q\n", fm.FileMode().Perm())
	// Output:
	// Unix mode: "r-sr-srwx"
	// Go FileMode: "-r-xr-xrwx"

	// Note that Go FileMode does not show suid bits
}

func ExampleModeFilemodeMode() {
	modes := []unixmode.Mode{
		unixmode.ModeNamedPipe,
		unixmode.ModeCharDevice,
		unixmode.ModeDir,
		unixmode.ModeDevice,
		unixmode.ModeRegular,
		unixmode.ModeSymlink,
		unixmode.ModeSocket,
		02777 | unixmode.ModeRegular,
		04525 | unixmode.ModeRegular,
		01331 | unixmode.ModeRegular,
	}
	for _, m := range modes {
		fi := m.FileMode()
		m2 := unixmode.New(fi)
		if m != m2 {
			fmt.Printf("mode: %07o -> % 12o -> %07o\n", m, fi, m2)
		}
	}
	// Output:
}
func ExampleFilemodeModeFilemode() {
	modes := []fs.FileMode{
		fs.ModeNamedPipe,
		fs.ModeCharDevice | fs.ModeDevice,
		fs.ModeDir,
		fs.ModeDevice,
		0,
		fs.ModeSymlink,
		fs.ModeSocket,
		0777 | fs.ModeSetgid,
		0525 | fs.ModeSetuid,
		0331 | fs.ModeSticky,
	}
	for _, fi := range modes {
		m := unixmode.New(fi)
		fi2 := m.FileMode()
		if fi != fi2 {
			fmt.Printf("mode: % 12o -> %07o -> % 12o\n", fi, m, fi2)
		}
	}
	// Output:
}
