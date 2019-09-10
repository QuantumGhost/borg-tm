package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/quantumghost/borg-tm/consts"
	"github.com/quantumghost/borg-tm/internal"
)

func main() {
	var borgArgs, mountpoint, lockFile string
	flag.StringVar(&borgArgs, "borg-args", "", "argument passed to `borg create`")
	flag.StringVar(&mountpoint, "mountpoint", "/tmp/snapshot",
		"mountpoint for snapshot, should be kept the same across backups")
	flag.StringVar(&lockFile, "lock-file", "/var/run/borg.lock", "lock file for borg-tm")
	var printVersion bool
	flag.BoolVar(&printVersion, "V", false, "print version")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s

This program must be run as root.

Environment variables:
- BORG_REPO: repository to backup to
- BORG_PASSPHRASE: passphrase for borg repository

Arguments:
`, os.Args[0])

		flag.PrintDefaults()
	}
	flag.Parse()
	if printVersion {
		consts.PrintVersion()
		os.Exit(0)
	}

	repo := os.Getenv("BORG_REPO")
	if repo == "" {
		log.Fatalln("BORG_REPO not specified")
	}
	if pass := os.Getenv("BORG_PASSPHRASE"); pass == "" {
		log.Fatalln("BORG_PASSPHRASE not specified")
	}
	parts := strings.Split(borgArgs, " ")
	args := make([]string, 0, len(parts))
	for _, v := range parts {
		if v != "" {
			args = append(args, v)
		}
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancelFn := context.WithCancel(context.Background())
	go func() {
		<-sig
		cancelFn()
	}()

	if os.Getuid() != 0 {
		log.Fatalln("requires root privileges.")
	}
	backup := internal.NewBackup(mountpoint, lockFile, args)
	err := backup.Run(ctx)
	if err != nil {
		log.Fatalf("error while backup: %+v\n", err)
	}
}
