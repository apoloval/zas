## Welcome to ZAS

ZAS is yet another Zilog Z-80 assembler. It's purpose is to run on modern computers to generate code for retrocomputers (MSX, Spectrum, Amstrad, etc).

## Why another assembler?

Mainly because:
* I think writing an assembler is fun.
* Many existing assemblers are not so recent. Most of them were written in late 90s or early 2000s. Surprisingly, most of them only work on Windows. Some others work on Unix, but have SIGSEGV failures. Some others simply works, but are too simple and really limited.
* I'd like to explore and experiment new ways to write ASM code for retrocomputers. Things like:
    * The ability to write unit tests for your Z-80 code and execute them in a virtual CPU.
    * Have some high-level abstractions that leave you between ASM and C. Close enough to the CPU to ensure the performance but removing some tedious ASM boilerplate.
    * Have some static analysis to prevent subtle errors. Things like unset arguments for a subroutine, modification of registers that were smashed by some previous call, etc. 

## How can I use ZAS?

This is WIP. There is no build system or installer you can use to try it out, but...

This is open source, and it's written in Go. You can checkout this repository and build the sources. And, if you have any question (or even better: if you want to contribute!!!), please contact me.
