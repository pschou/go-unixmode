package unixmode_test

import (
	"fmt"
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

func ExampleMode_PermString_stickyBits() {
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
