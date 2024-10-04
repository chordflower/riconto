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

### Fetch ###

`riconto fetch <url> [<path>]`

Fetches a project from various sources:

- git => Clone from a remote git repository;
- mercurial => Clone from a remote mercurial repository;
- ftp => Copy from a remote ftp site;
- webdav => Copy from a remote webdav site.

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

### Clean ###

Cleans the built files, it has the following parameters:

- name => The project(s) to clean separated by commas or multiple flags.

### Add ###

Adds a subproject, it has the following parameters:

- name => The subproject name;
- output => The subproject output path (default `./dist/<name>/`);
- source => The subproject source path (default `./src/<name>/main.md`).

### Remove ###

Removes a subproject, it has the following parameters:

- name => The subproject to remove;

### List ###

Lists all the available subprojects.
