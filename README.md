Description
===========

This allows MySQL Shell (mysqlsh) to read the username and password from `~/.my.cnf`, similar to
what `mysql` does by default.

Usage
=====

    mysqlsh --credential-store-helper=mycnf --sql root@localhost

Or you can put something like this in `~/.mysqlsh/options.json`:

    {
        "credentialStore.helper": "mycnf",
        "defaultMode": "sql"
    }

And then:

    mysqlsh root@localhost

Installation
============

From binary: [Download](https://github.com/dveeden/mysql-secret-store-mycnf/releases) and place in the same directory as mysqlsh

From source:

    go get github.com/go-ini/ini
    go build
    mv mysql-secret-store-mycnf $(dirname $(which mysqlsh))

Related
=======

* [MySQL Bug #93003](https://bugs.mysql.com/bug.php?id=93003)

