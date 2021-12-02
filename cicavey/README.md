# Using gotip (aka 1.18 for generics!)

See https://sebastian-holstein.de/post/2021-11-08-go-1.18-features/

Assuming you already have go installed,

Preemptively append `~/go/bin` to your PATH which will expose "gotip" after we install it.

    go install golang.org/dl/gotip@latest
    gotip download

From here switching Go env should be pretty easy in vscode `Crtl+Shift+p` then `goen` will offer a command to chose environment.
