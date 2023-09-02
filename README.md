# Barber: CLI & TUI Snippet Manager

Barber is a terminal snippet manager inspired by [nap](https://github.com/maaslalani/nap). It's designed to help you organise and access your code snippets efficiently. Barber is currently under development, and this README provides an overview of its planned features and functionality.

If you have any ideas or feedback please reach out to me [via email](mailto:broderickwestrope@gmail.com) or start a discussion on this repo.

## Key Changes from Nap

While Barber draws inspiration from Nap, it introduces several key changes:

- **Flexible Default Directory**: Barber allows you to easily set the default directory for storing snippets. By default, snippets are stored in the current directory, giving you more control over snippet storage.

- **Multiple Storage Locations**: You can store snippets in different directories, and Barber follows a specific order of precedence: specified directory, global default (if set), and the current directory.

- **Git Integration**: Barber aims to seamlessly integrate with Git, allowing you to host your snippet collections in Git repositories using any remote provider. It also provides the option for local version control.

## Glossary

Before diving into the features, here are some important terms to understand:

- **Collection**: A folder containing snippets and subfolders. Collections are defined by a collection config file and collection metadata. Barber selects the collection based on precedence: explicitly defined collection/alias, working directory, or defaults.

- **Snippet**: A text-based file stored within a collection. Snippet references are relative paths from the root of the active collection.

- **Alias**: A unique string beginning with `~` that represents a collection alias. Aliases are defined and stored within the app metadata.

- **Config**: Barber has two types of configuration:
    - **App**: Stores mappings for aliases and their collection paths (relative to the user's home directory). It also holds default TUI configuration, which can be overridden by collection-specific settings.
    - **Collection**: Contains overrides for TUI configuration specific to a collection.

- **Metadata**: Two types of metadata are maintained:
    - **App**: Stores a list of aliases (mapping alias values to collection paths).
    - **Collection**: Holds a list of snippets, each with a description, tags, and a path relative to the collection root.

## CLI

Barber offers a comprehensive CLI for managing your snippets. Here's an overview of some planned commands:

- `barber`: Start an interactive mode in the selected collection.
    - *--collection, -c* (string): Specify the path to a collection or a collection alias with the prefix.

- `barber list`: List all snippets in the collection.
    - *--collection, -c* (string): Specify the path to a collection or a collection alias with the prefix.

- `barber rm`: Delete a specific file from the collection.
    - **path** (string, positional, required): The file to remove, relative to the collection root.
    - *--collection, -c* (string): Specify the path to a collection or a collection alias with the prefix.

- `barber scan`: Update metadata by scanning files and directories. Attempts to carry metadata along renames & moves using Git diff.
    - *--collection, -c* (string): Specify the path to a collection or a collection alias with the prefix.
    - *--keep* (bool, default false): Keep metadata for snippets thought to be deleted (allows manual alterations).

- `barber find`: Search for text in the collection's snippets and list the results.
    - **search** (string, positional, optional): A string to search for (fuzzy search of title and description).
    - *--tag* (string slice): A list of tag strings to filter by (exact match).
    - *--desc* (string): A string for fuzzy searching the description.
    - *--title* (string): A string for fuzzy searching the title.

- `barber config`: Commands to work with app configuration unless a collection argument is provided.
    - `reset`: Reset the config to the default.
        - *--yes, -y* (bool, default false): Automatically accept the confirmation message.
        - *--collection, -c* (string): Specify the path to a collection or a collection alias with the prefix.
    - `open`: Open the config file in the user's preferred editor.
        - *--collection, -c* (string): Specify the path to a collection or a collection alias with the prefix.

- `barber collection`: Commands to manage collections.
    - `init`: Initialise a new collection.
        - *--collection, -c* (string): Specify the path to a collection or a collection alias with the prefix.
    - `rm`: Remove the collection (config & metadata).
        - *--yes, -y* (string, default false): Automatically accept the confirmation message.
        - *--collection, -c* (string): Specify the path to a collection or a collection alias with the prefix.
    - `reset`: Reset the collection (config & metadata).
        - *--yes, -y* (string, default false): Automatically accept the confirmation message.
        - *--collection, -c* (string): Specify the path to a collection or a collection alias with the prefix.

- `barber alias`: Commands to manage collection aliases.
    - `list`: List all aliases.
    - `add`: Create an alias with the given collection.
        - **alias** (string, positional, required): The alias (without the prefix).
        - **collection** (string, positional
