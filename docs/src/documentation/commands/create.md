---
title: "Riconto create command"
description: "This is the documentation for riconto create command"
authors:
  - name: "carddamom"
    email: "carddamom at tutanota dot com"
tags:
  - riconto
  - documentation
  - command
metadata:
  created: "2024-10-09T11:42:12.791404Z"
  published: "2024-10-09T11:42:12.791404Z"
  modified: "2024-10-09T11:42:12.791404Z"
---

The create command creates a new riconto project in the current directory.

By default it will create:

- A configuration file describing the new project in one of the three supported formats;
- A directory named src, with a file main.md that represents the main file of the project;
- A directory named resources that contains the resources used by the project, like images.

The configuration file can be in:

- JSON;
- YAML;
- TOML;

It accepts the following options:

- name => The name of the project, as recorded in the configuration file (required);
- version => The project version, as recorded in the configuration file, it is advised to use a format like semantic versioning or calendar versioning, by default it is 0.0.1;
- description => A description for the project;
- license => The main license of the project;
- format => The format to use for the configuration file, it can be json, yaml or toml and by default it is toml;
- strict => If an existing configuration file should make riconto return with error code 1, instead of skipping the project creation silently.

The exit codes are:

- 0 => If the command succeded or if there was an existing projet in the directory and strict, was not given;
- 1 => If an error happened.

#### Inner Workings ####

The command will start by:

1. Validating all of the given parameters;
2. Checking if the current directory has a configuration file;
3. Create a temporary directory, for working;
4. Create the configuration file in the choosen format on the temporary directory;
5. Create the auxiliary directories (src and resources) on the temporary directory;
6. Creating an empty file named main.md on the temporary directory;
7. Copy the temporary directory contents to the final directory;
8. Delete the temporary directory and all of its contents (even if an error happened).
