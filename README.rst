=============
What is this?
=============

This is a docker wrapper around https://github.com/apuigsech/seekret. The original library no longer builds with the git functionality because it requires an older version of https://github.com/libgit2/git2go. We did not need it, so we removed it and retained the directory scan functionality. We then wrapped it into a Docker image to use as part of CI. 

************
seekret
************

Go library and command line to seek for secrets on various sources.


************
Command Line
************

Description
===========

``seekret`` inspect a directory to seek for secrets. It can be used to prevent that secrets are
published in exposed locations.


Installing seekret
==================

``seekret`` can be directly installed by using go get.

::

    go get github.com/cds-snc/seekret/cmd/seekret


The requirements for a success installation are:

 * golang >= 1.6

Usage
=====

General Options
~~~~~~~~~~~~~~~

::

    NAME:
       seekret - seek for secrets on various sources.

    USAGE:
       seekret [global options] command [command options] [arguments...]

    VERSION:
       0.0.1

    AUTHOR(S):
       Albert Puigsech Galicia <albert@puigsech.com>

    COMMANDS:
       seek:
         dir    seek for seecrets on a directory.

    GLOBAL OPTIONS:
       --exception FILE, -x FILE    load exceptions from FILE.
       --rules PATH         PATH with rules. [$SEEKRET_RULES_PATH]
       --format value, -f value specify the output format. (default: "human")
       --known FILE, -k FILE  load known secrets from FILE.
       --workers value, -w value  number of workers used for the inspection (default: 4)
       --help, -h           show help
       --version, -v        print the version


``-x, --exception``

``--rules``

``-f, --format``

``-k, --known``

``-w, --workers``


Options for Dir
~~~~~~~~~~~~~~~

::

    NAME:
       seekret dir - seek for seecrets on a directory.

    USAGE:
       seekret dir [command options] [arguments...]

    CATEGORY:
       seek

    OPTIONS:
       --recursive, -r
       --hidden


``-r, --recursive``

``-h, --hidden``



Examples
========

Scan all files (including hidden) in a local folder::

    seekret dir --recursive --hidden /path/to/dir
