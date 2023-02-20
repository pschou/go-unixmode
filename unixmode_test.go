package unixmode_test

import (
	"fmt"
	"log"
	"os"

	"github.com/pschou/go-unixmode"
)

func ExampleString() {
	m := unixmode.UnixMode(041755)
	fmt.Printf("mode: %q\n", m.String())

	// Output:
	// mode: "drwxr-xr-t "
}

func ExamplePermString() {
	m := unixmode.UnixMode(0755)
	fmt.Printf("mode: %q\n", m.PermString())

	// Output:
	// mode: "rwxr-xr-x"
}

func ExamplePermString_Directory() {
	m := unixmode.UnixMode(040755)
	fmt.Printf("mode: %q\n", m.PermString())

	// Output:
	// mode: "rwxr-xr-x"
}

func ExamplePermString_StickyBits() {
	m := unixmode.UnixMode(06755)
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

func ExampleString_Socket() {
	if stat, err := os.Lstat("/dev/log"); err != nil {
		log.Fatal("Could not stat socket", err)
	} else {
		m := stat.Mode()
		fmt.Printf("mode: %q\n", unixmode.FileModeString(m))
	}

	// Output:
	// mode: "srw-rw-rw- "
}

func ExampleString_Chmod() {
	m := unixmode.UnixMode(02644 | unixmode.ModeRegular)
	fmt.Printf("mode: %q\n", m.PermString())
	unixmode.Chmod("t", m)

	// Output:
	// mode: "rw-r-Sr--"
}

func ExampleString_RawSUDO() {
	if stat, err := os.Lstat("/usr/bin/sudo"); err != nil {
		log.Fatal("Could not stat character device", err)
	} else {
		m := stat.Mode()
		fmt.Printf("mode: %c %04o\n", unixmode.FileModeTypeLetter(m), unixmode.FileModePerm(m))
	}

	// Output:
	// mode: - 4111
}

func ExampleString_RawCharacter() {
	if stat, err := os.Lstat("/dev/lp0"); err != nil {
		log.Fatal("Could not stat character device", err)
	} else {
		m := stat.Mode()
		fmt.Printf("mode: %c %04o\n", unixmode.FileModeTypeLetter(m), unixmode.FileModePerm(m))
	}

	// Output:
	// mode: c 0660
}

func ExampleString_Character() {
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
		fmt.Printf("mode: %04o\n", *m)
	}
	// Output:
	// mode: 0770
}

func ExampleParse_DirectoryMixedRaw() {
	if m, err := unixmode.Parse("d-w-r-S-wT"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("mode: %04o\n", *m)
	}
	// Output:
	// mode: 43242
}

func ExampleParse_Mixed() {
	if m, err := unixmode.Parse("-w-r-S-wT"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("mode: %04o\n", m.Perm())
	}
	// Output:
	// mode: 3242
}

func ExampleParse_Invalid() {
	if m, err := unixmode.Parse("drwSrwSrwS "); err != nil {
		fmt.Println("Err:", err)
	} else {
		fmt.Printf("mode: %04o\n", uint32(*m))
	}
	// Output:
	// Err: Invalid 'S' at position 9
}
