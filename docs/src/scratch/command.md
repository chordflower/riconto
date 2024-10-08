---
title: Commands
tags:
  - scratch
  - ideas
  - commands
---

## Commands ##

Possible commands:

- create (old init) => Creates a project from scratch;
- fetch => Fetches a projet from a repository;
- build => Builds the project;
- clean => Cleans the built files;
- add => Adds a new subproject;
- remove => Removes a subproject;
- list => Lists the available subprojects.

### Create ###

Creates a new project from scratch, it can receive:

- name (required) => Name of the project;
- version => Version of the project;
- description => Description of the project;
- license => License of the project.
- strict => If an existing project in the current directory should, cause the command to fail with error code 1, instead of ignoring the create and exit with 0.

Returns the error code:

0 => On success;
1 => If something happens.

### Fetch ###

`riconto fetch <url> [<path>]`

Fetches a project from various sources:

- git => Clone from a remote git repository;
- mercurial => Clone from a remote mercurial repository;
- ftp => Copy from a remote ftp site;
- webdav => Copy from a remote webdav site.

Returns the error code:

0 => On success;
1 => If something happens.

#### URLs ####

- git:
  - git.https://...
  - git.git://...
  - git.ssh://...
- mercurial:
  - hg.https://...
  - hg.http://...
  - hg.ssh://...
- ftp:
  - ftp://...
  - ftps://...
- webdav:
  - webdav.http://...
  - webdav.https://...

### Build ###

Builds the project, it has the following parameters:

- name => The project(s) to build separated by commas or multiple flags.
- warnings-as-errors => Fail if there are any warnings.

Returns the error code:

0 => On success;
1 => If something happens.

### Clean ###

Cleans the built files, it has the following parameters:

- name => The project(s) to clean separated by commas or multiple flags.

Returns the error code:

0 => On success;
1 => If something happens.

### Add ###

Adds a subproject, it has the following parameters:

- name => The subproject name;
- output => The subproject output path (default `./dist/<name>/`);
- source => The subproject source path (default `./src/<name>/main.md`).

Returns the error code:

0 => On success;
1 => If something happens.

### Remove ###

Removes a subproject, it has the following parameters:

- name => The subproject to remove;

Returns the error code:

0 => On success;
1 => If something happens.

### List ###

Lists all the available subprojects, it has one parameter:

- format => The output format it can be:
  - text => The default value, that prints a human readable list, using a table;
  - json => Outputs the list in json format;
  - yaml => The same as json, but in yaml format;
  - toml => The same as json, but in toml format;
  - csv => The same as json, but in csv format.

Returns the error code:

0 => On success;
1 => If something happens.
