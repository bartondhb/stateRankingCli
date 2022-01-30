# U.S. State and Territory Ranking Data CLI
## Overview
This command line tool allows the user to view and summarize U.S. State and Territory population Data from the Wikipedia article at https://en.wikipedia.org/wiki/List_of_U.S._states_and_territories_by_population. An internet connection is required to use this tool as the list is retrieved at runtime and is not cached locally.

## Building
**_NOTE:_** This application requires Go 1.17 or later.

1. Clone Repository
    ```shell
   git clone https://github.com/bartondhb/stateRankingCli.git
   ```
2. Switch to the directory where the repo has been cloned
3. Build the Executable:
    ```shell
    go build -o uspop .
   ```

## Usage

The `uspop` CLI accepts two flags:
```shell
-state string

-summary
```

### State Flag

The `-state` flag is followed by the name of the state for which you want to retrieve population data. For example:
```shell
-state california
```
This will display the population for the state of California.

**_NOTE:_** This field is not case sensitive.

Additionally, there are two other options that can be supplied with the `state` flag:
```shell
-state list
-state all
```

`-state list` will print an alphabetical list of all U.S. States and Territories.

`-state all` will print the population of all U.S. States and Territories.

### Summary Flag
The `-summary` flag will print several summarized pieces of information about the U.S. Population including:
- Total Population
- Mean Population
- Media Population
- Standard Deviation of the Population
