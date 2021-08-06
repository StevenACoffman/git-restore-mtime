# git-restore-mtime
*Restore original modification time of files based on the date of the most recent commit that modified them*

Git, unlike other version control systems, does not preserve the original timestamp of committed files. Whenever repositories are cloned, or branches/files are checked out, file timestamps are reset to the current date. While this behavior has its justifications (notably when using `make` to compile software), sometimes it is desirable to restore the original modification date of a file (for example, when generating release tarballs). As git does not provide any way to do that, `git-restore-mtime` tries to workaround this limitation.

For more information and background, see http://stackoverflow.com/a/13284229/624066

This Go source code was originally from [Steve Penny](https://stackoverflow.com/users/1002260/steven-penny) from https://stackoverflow.com/a/62683363/2904179 
(and I am a different Steve 😀).