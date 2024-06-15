package gostrings

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestStrings(t *testing.T) {
	// Read assets/main binary file
	f, err := os.Open("assets/main")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expected := []string{
		"/lib64/ld-linux-x86-64.so.2",
		"__cxa_finalize",
		"__libc_start_main",
		"puts",
		"libc.so.6",
		"GLIBC_2.2.5",
		"GLIBC_2.34",
		"_ITM_deregisterTMCloneTable",
		"__gmon_start__",
		"_ITM_registerTMCloneTable",
		"PTE1",
		"u+UH",
		"Hello, World!",
		":*3$\"",
		"GCC: (Ubuntu 11.4.0-1ubuntu1~22.04) 11.4.0",
		"Scrt1.o",
		"__abi_tag",
		"crtstuff.c",
		"deregister_tm_clones",
		"__do_global_dtors_aux",
		"completed.0",
		"__do_global_dtors_aux_fini_array_entry",
		"frame_dummy",
		"__frame_dummy_init_array_entry",
		"main.c",
		"__FRAME_END__",
		"_DYNAMIC",
		"__GNU_EH_FRAME_HDR",
		"_GLOBAL_OFFSET_TABLE_",
		"__libc_start_main@GLIBC_2.34",
		"_ITM_deregisterTMCloneTable",
		"puts@GLIBC_2.2.5",
		"_edata",
		"_fini",
		"__data_start",
		"__gmon_start__",
		"__dso_handle",
		"_IO_stdin_used",
		"_end",
		"__bss_start",
		"main",
		"__TMC_END__",
		"_ITM_registerTMCloneTable",
		"__cxa_finalize@GLIBC_2.2.5",
		"_init",
		".symtab",
		".strtab",
		".shstrtab",
		".interp",
		".note.gnu.property",
		".note.gnu.build-id",
		".note.ABI-tag",
		".gnu.hash",
		".dynsym",
		".dynstr",
		".gnu.version",
		".gnu.version_r",
		".rela.dyn",
		".rela.plt",
		".init",
		".plt.got",
		".plt.sec",
		".text",
		".fini",
		".rodata",
		".eh_frame_hdr",
		".eh_frame",
		".init_array",
		".fini_array",
		".dynamic",
		".data",
		".bss",
		".comment",
	}
	result, err := GetStrings(f, 4, 1024)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
