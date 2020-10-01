package main

import (
	"fmt"
	"log"

	libseccomp "github.com/seccomp/libseccomp-golang"
)

func SeccompInit() error {
	syscalls := []string{
		"write",      // 1
		"nanosleep",  // 35
		"exit_group", // 231
	}

	// Some newer syscalls aren't universally available to lookup by string.  Add them manually
	// Note these will be specific to your architecture, if you compile for multiple architectures
	// then adding newer syscalls will be tricker.
	syscallsByNumber := []libseccomp.ScmpSyscall{
		// 318, //getrandom
		// 332, //statx
	}

	// ActLog is the most concervative action.  Your proram will run as normal
	// Requires libseccomp API level 3 or later, so may not be available on old versions of Linux
	filter, err := libseccomp.NewFilter(libseccomp.ActLog)

	// As you can confidence in your system call list you may want to switch to returning EPERM or ActKill
	// instead of ActLog
	//filter, err := libseccomp.NewFilter(libseccomp.ActErrno.SetReturnCode(int16(syscall.EPERM)))
	//filter, err := libseccomp.NewFilter(libseccomp.ActKill)

	if err != nil {
		return err
	}

	for _, element := range syscalls {
		syscallID, err := libseccomp.GetSyscallFromName(element)
		if err != nil {
			log.Printf("Could not whitelist syscall:%v", element)
			continue
		}
		err = filter.AddRule(syscallID, libseccomp.ActAllow)
		if err != nil {
			log.Printf("Could not whitelist syscall rule:%v", element)
		}
	}

	for _, syscallID := range syscallsByNumber {
		err = filter.AddRule(syscallID, libseccomp.ActAllow)
		if err != nil {
			log.Printf("Could not whitelist syscall rule:%v", syscallID)
		}
	}

	err = filter.Load()
	if err != nil {
		return fmt.Errorf("could not load seccomp: %w", err)
	}

	return nil

}
