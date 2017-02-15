[![wercker status](https://app.wercker.com/status/52243183472f21dc72756d12a649ee19/m/master "wercker status")](https://app.wercker.com/project/bykey/52243183472f21dc72756d12a649ee19)
[![go report status](https://goreportcard.com/badge/github.com/soracom/soracom-cli)](https://goreportcard.com/report/github.com/soracom/soracom-cli)

# soracom-cli

A command line tool `soracom` to invoke SORACOM API.

# Feature

The `soracom` command:

- supports new APIs on-time. The binary file of the soracom command is automatically generated from the API definition file.

- just works by copying the cross-compiled binary file into the target environment. There is no need to build an environment or solve dependencies.

- constructs a request based on the specified argument and calls the SORACOM API. Response (JSON) from the API is output directly to standard output.
  - This makes it easier to process the output of the soracom command and pass it to another command

- supports bash completion. Please write the following line in .bashrc etc
  ```
  eval "$(soracom completion)"
  ```

  if you are a macOS user, you probably need to either:
  1. use `bash` version >= 4.0, or
  2. use `brew install bash-completion` instead of using Xcode version of bash-completion

  otherwise you might be getting the error like the following:
  ```
  -bash: __ltrim_colon_completions: command not found
  ```

# How to install

## Using Mac (macOS), installing by homebrew

```
$ brew tap soracom/soracom-cli
$ brew install soracom-cli
```

## In other cases
Download a package file that match the environment of the target from [Releases page](https://github.com/soracom/soracom-cli/releases), unpack it, and place the executable file in the directory where included in PATH.


# How to use

## Basic usage

First of all, create a profile by running the following command:

```
soracom configure
```

You will be asked which coverage type to use.

```
Please select which coverage type to use.

1. Global
2. Japan

select (1-2) >
```

Please select the coverage type which you mainly use. In most cases, please select Global. If you live in Japan and use SIM cards in Japan, please select Japan.

Next you will be asked about the authentication method.

```
Please select which authentication method to use.

1. Input AuthKeyId and AuthKey * Recommended *
2. Input Operator credentials (Operator Email and Password)
3. Input SAM credentials (OperatorId, User name and Password)

select (1-3) >
```

Please select 1 if AuthKey (authentication key) has been issued to SAM user or root account.
(For details on how to issue an authentication key to SAM users, please see [Using SORACOM Access Management to Manage Operation Access](https://dev.soracom.io/en/start/sam/).

Thereafter, when executing the soracom command, an API call is made using the authentication information entered here.



## Advanced usage

### Use multiple profiles

If you have multiple SORACOM accounts or want to use multiple SAM users differently, specify the --profile option to configure and set the profile name.

```
soracom configure --profile user1
  :
  (Enter information for user1)

soracom configure --profile user2
  :
  (Enter information for user2)
```

This will create profiles named user1 and user2.
To use the profile, specify the --profile option in addition to the normal command.

```
soracom subscribers list --profile user1
  :
  (SIM list for user1 will be displayed)

soracom groups list --profile user2
  :
  (Group list for user2 will be displayed)
```


### Call API via proxy

Set `http://your-proxy-nme:port` to HTTP_PROXY environment variable, then execute soracom command.

e.g.) For Linux / Mac:
Assume that the address of the proxy server is 10.0.1.2 and the port number is 8080
```
export HTTP_PROXY=http://10.0.1.2:8080
soracom subscribers list
```

Or

```
HTTP_PROXY=http://10.0.1.2:8080 soracom subscribers list
```


# How to build / test

For developers who want to build from source or for those who wish to make a pull request such as bug fix / function addition, please build and test in one of the following ways.

## How to build in a local environment (Linux / Mac OS X)

In the environment where Go is installed, run the build script as follows:

```
./scripts/build.sh 1.2.3
```

Here 1.2.3 is the version number. Please specify an appropriate number.

If the build succeeds, then run the test:

```
./test/test.sh
```


## How to build using wercker

Install wercker's CLI and execute the build as follows. It will run the test as well.

```
wercker build
```

TODO: Currently, the build result is generated in the container. It will be fixed to generate to the mounted volume.
