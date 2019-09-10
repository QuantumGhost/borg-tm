# Borg TimeMachine

Use [borg]() and APFS snapshot to backup your Mac.

## FAQ

Q: Why not use Time Machine directly?

TL;DR: Because Time Machine is slow and problematic.

Time Machine is terriable slow. It seems to take forever for backup my 256G MacBook, even if I have 
1G Ethernet attached.

There are many cases that TM is always keeping backup but never succeed. 
There are cases that TM backups cannot be restored. I don't want to put my trust on such a unstable 
software.

Q: How does this works?

Thanks to APFS, we can easily take snapshot for filesystem now.

This command works by first using `tmutil localsnapshot` to create a snapshot of your root filesystem.
Then use `mount` to mount it to a path, and backup it with `borg`. Due to borg's deduplication mechaniusm,
it's recommended that you mount snapshot to the same path every time.

I have written a [bash script](https://gist.github.com/QuantumGhost/1aae8eb8527c9d522fe2a57f214f6ee5) to do
basically the same thing. You may consult it if you want to know more. **Please note due to the lack of `flock(1)`,
this bash script can be run in parallel and cause problems.**
