# React Component Creator

Create React components with boilerplate without the hassle.

# Installation

The installation is very simple, just add the rcc or rcc.exe in the build folder to your PATH
depending on your OS and you're set!

## Linux and Mac PATH

On linux and mac you execute the following command to add it to your path:

> export PATH="$PATH:/path/to/the/repo/build/rcc"

## Windows Path

On Windows you must set the path variable by searching the **Edit the system environment variables**
section on your pc, opening it, then clicking **Environment Variables** on the right bottom corner,
clicking Edit in the **System Variables** section after choosing the **Path** variable
and then adding the full path there, via clicking New in the right top corner,
without the exe file itself like:

> C/path/to/the/repo/build

Warning: You may need to restart your PC after changing the PATH variable.

# Usage

Just run the command line tool with the command **rcc "Component Name Here"** and you're set!

To disable TypeScript, run with the argument **--no-ts**

The default test suite is Cypress, you can change it to jest by using the **--use-jest** flag.
