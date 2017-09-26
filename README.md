`$ dev`
=======

A general-purpose command runner; kept as minimal as possible.

**Use cases**
  - Setup a development environment.
  - Run tests after saves.
  - Trigger deployments after builds.
  - Document your development workflow in a readable/executable format.

Installation
------------

```
$ go get -u github.com/ef2k/dev/...
```

Usage
-----

#### Step 1 - Create a configuration file

```sh
$ dev init     # creates dev.yaml
```

#### Step 2 - Configure commands

```yaml
# dev.yaml

watch: . # watch all files

ignore:
  - huge-files

tasks:
  - cmd: go install ./cmd/cli
    title: "Install the cli utility after every file change"

  - cmd: go test -v ../...
    title: "Run tests whenever a test file changes"

  - cmd: ./web_app
    title: "Run the web application with environment variables set."
    env:
      APP_ENV: development
      PORT: 3000
      DB: "sql://user:pass@host/db"

bg-tasks:
  - cmd: markdown-preview {}
    title: "Preview markdown whenever a markdown file changes"
    watch: *.md

  - cmd: say "moo"
    title: "Moo whenever the moo file is changed"
    watch: moo.txt

```

#### Step 3 - Run it

```sh
$ dev
```

License
-------
MIT
