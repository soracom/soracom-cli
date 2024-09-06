[![go report status](https://goreportcard.com/badge/github.com/soracom/soracom-cli)](https://goreportcard.com/report/github.com/soracom/soracom-cli)
![build-artifacts](https://github.com/soracom/soracom-cli/actions/workflows/build-artifacts.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/soracom/soracom-cli.svg)](https://pkg.go.dev/github.com/soracom/soracom-cli)

# soracom-cli

Provides `soracom` command, a command line tool for calling SORACOM APIs.

# Feature

The `soracom` command:

- supports new APIs on-time. The binary file of the soracom command is automatically generated from the API definition file.

- just works by copying the cross-compiled binary file into the target environment. There is no need to build an environment or solve dependencies.

- constructs a request based on the specified arguments and calls the SORACOM API. Response (JSON) from the API is output directly to standard output.
  - This makes it easier to process the output of the soracom command and pass it to another command

- supports bash completion. Please write the following line in .bashrc etc
  ```
  eval "$(soracom completion bash)"
  ```

  - If you get an error like this:

    ```
    -bash: __ltrim_colon_completions: command not found
    ```

    This error may appear if you are using macOS. You might need to satisfy one of the following conditions:

    - use `bash` version >= 4.0, or
    - use `brew install bash-completion` instead of using Xcode version of bash-completion and then add the following to either your `.bash_profile` or `.profile`:

      ```
      if [ -f "$(brew --prefix)/etc/bash_completion" ]; then
        . "$(brew --prefix)/etc/bash_completion"
      fi
      ```

- supports zsh completion. The generated completion script by running the following command should be put somewhere in your `$fpath` named `_soracom`
  ```
  soracom completion zsh
  ```

# How to install

## Using Mac (macOS) or Linux, installing by homebrew

```shell
brew tap soracom/soracom-cli
brew install soracom-cli
brew install bash-completion
```

## In other cases

By running one of the following commands, the latest version of `soracom` command will be installed.

If you have a permission to write a file into `/usr/local/bin` directory (e.g. you are `root` user), please run the command below:

```shell
curl -fsSL https://raw.githubusercontent.com/soracom/soracom-cli/master/install.sh | bash
```

If you do not have a permission to write a file into `/usr/local/bin` directory, please run either of the following commands.

If you are in sudoers and want to install `soracom` command to `/usr/local/bin`:

```shell
curl -fsSL https://raw.githubusercontent.com/soracom/soracom-cli/master/install.sh | sudo bash
```

or

If you are not in sudoers or want to install `soracom` command to other directory e.g. `$HOME/bin`:

```shell
mkdir -p "$HOME/bin"
curl -fsSL https://raw.githubusercontent.com/soracom/soracom-cli/master/install.sh | BINDIR="$HOME/bin" bash
```

You can change `"$HOME/bin"` in the command above to wherever you want.

If you want to upgrade the `soracom` command, you can just run the same command you used to install `soracom` again.

If you want to uninstall the `soracom` command, you can just remove `soracom` executable file you have installed. (You may want to remove `$HOME/.soracom/` directory which contains profiles for the `soracom` command.)

If the commands above did not work well, or if you want to install older version of `soracom` command, you can download a package file that match the environment of the target from [Releases page](https://github.com/soracom/soracom-cli/releases), unpack it, and place the executable file in the directory where included in `PATH`.


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
4. Switch user

select (1-4) >
```

Please select 1 if AuthKey (authentication key) has been issued to SAM user or root account.
(For details on how to issue an authentication key to SAM users, please see [Users & Roles | SORACOM Developers](https://developers.soracom.io/en/docs/security/users-and-roles/).

If you select 4. Switch user, you can specify the Operator ID and SAM user name of the switch destination user. Please create a profile for the switch source user before configuring a switch user profile. If you specify a switch user profile, soracom-cli will automatically authenticate with the switch source profile and then switch to the SAM user before making API calls.

Thereafter, when executing the soracom command, an API call is made using the authentication information entered here.



## Advanced usage

### Use multiple profiles

If you have multiple SORACOM accounts or want to use multiple SAM users differently, specify the `--profile` option to configure and set the profile name.

```
soracom configure --profile user1
  :
  (Enter information for user1)

soracom configure --profile user2
  :
  (Enter information for user2)
```

This will create profiles named user1 and user2.
To use the profile, specify the `--profile` option in addition to the normal command.

```
soracom subscribers list --profile user1
  :
  (SIM list for user1 will be displayed)

soracom groups list --profile user2
  :
  (Group list for user2 will be displayed)
```


### Create a profile for API Sandbox

It is possible to use soracom-cli for setting up [SORACOM API Sandbox](https://dev.soracom.io/en/docs/api_sandbox/) environment.

In order to create a profile for sandbox, use `configure-sandbox` subcommand.

```
soracom configure-sandbox
```

By answering to the questions prompted, a profile named `sandbox` will be created. By using the `sandbox` profile, you can issue commands to the API sandbox as follows.

```
soracom subscribers list --profile sandbox
```

You can use commands dedicated for the sandbox.

```
soracom sandbox subscribers create --profile sandbox
```

You can use different profile name.

```
soracom configure-sandbox --profile test
soracom sandbox subscribers create --profile test
```

In order to make it easier to use from shell scripts etc., all the parameters necessary for profile creation can be specified with arguments.

```
soracom configure-sandbox --coverage-type jp --auth-key-id="$AUTHKEY_ID" --auth-key="$AUTHKEY" --email="$EMAIL" --password="$PASSWORD"
```

### Priority of authentication methods specified by command line arguments

The soracom-cli internally authenticates to obtain the API key and token to make calls to the SORACOM API.
These keys and tokens are then sent along with the API requests.

There are several options available to authenticate or to specify the API key and token, which are used as follows:

1. Use the previously authenticated and obtained API key and token directly by specifying them with the `--api-key` and `--api-token` options to make API calls.

2. Authenticate by specifying the authentication key ID and authentication key with the `--auth-key-id` and `--auth-key options`, then obtain the API key and token, and use them to make API calls.

3. Generate a profile (which contains the information for authentication) by executing an external command specified by the `--profile-command` option, then use that profile to authenticate, obtain the API key and token, and make the API calls.

4. Authenticate using a pre-configured profile specified with the `--profile` option, then obtain the API key and token, and use them to make the API calls.

These methods are ranked in priority from 1 to 4, with 1 being the highest priority and 4 the lowest.
For example, if a soracom-cli user specifies both the `--profile-command` and `--profile` options at the same time, the contents of the `--profile-command` option will take precedence.

Additionally, if options that need to be specified together, like `--api-key` and `--api-token` or `--auth-key-id` and `--auth-key`, are provided singly, it will result in an error.


### Priority of authentication methods specified in the profile

Within a profile, you can specify one of the following authentication methods:

1. The `profileCommand` field to specify an external command to generate profile information.

2. The `sourceProfile` field to specify the original profile when using the switch user feature, and the `operatorId` and `username` fields to specify the operator ID and username for the target switch.

3. The `authKeyId` and `authKey` fields to specify the authentication key ID and authentication key.

4. The `email` and `password` fields to specify the root user's email address and password.

5. The `operatorId`, `username`, and `password` fields to specify the SAM user's operator ID, username, and password.

These methods are ranked in priority from 1 to 5, with 1 being the highest priority and 5 the lowest.
For example, if the `profileCommand` field and both `authKeyId` and `authKey` fields are specified in a profile at the same time, the contents of the `profileCommand` will take precedence.

You cannot specify a `sourceProfile` within a profile that is being referenced by `sourceProfile`.


### Call API via proxy

Set `http://your-proxy-name:port` to HTTP_PROXY environment variable, then execute soracom command.

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

### Use AWS Lambda Layers of soracom-cli

Have you ever thought about using soracom-cli on AWS Lambda? By including the soracom-cli binary in your Zip package or container image and deploying it, you can use soracom-cli in your Lambda functions.

However, the soracom-cli binary is relatively large and may overwhelm the space in the Zip package or container image.

Therefore, we offer soracom-cli's Layer.

You can execute the `soracom` command in your Lambda function by specifying the following ARN:

- x86_64 architecture:
  ```
  arn:aws:lambda:ap-northeast-1:717257875195:layer:soracom-cli-${ver}:1
  ```

- arm64 architecture:
  ```
  arn:aws:lambda:ap-northeast-1:717257875195:layer:soracom-cli-${ver}-arm64:1
  ```

The `${ver}` part is the version number of the target soracom-cli without the `.`.

For example, for version `1.2.3`, `${ver}` should be `123`.

The binary is installed in /bin/soracom, which is PATHed and can be executed in Lambda functions simply as the `soracom` command.

In Node.js 18.x runtime, it can be called as follows: (Pass AUTH_KEY_ID and AUTH_KEY in the environment variables)

```
const execSync = require('child_process').execSync;
const jpBill = execSync(`soracom --auth-key-id ${process.env.AUTH_KEY_ID} --auth-key ${process.env.AUTH_KEY} bills get-latest --coverage-type jp`).toString();
```


### Trouble shooting

If you get an error message like the following:

```
Error: Permissions for the file 'path/to/default.json' which contains your credentials are too open.
It is required that your credential files are NOT accessible by others.
```

Please try the following to fix it:

```
soracom unconfigure
soracom configure
```

i.e. perform `unconfigure` and then `configure` again in order to re-create a credentials file with appropriate permissions.


# How to build / test

For developers who want to build from source or for those who wish to make a pull request such as bug fix / function addition, please build and test in one of the following ways.

## Update API definitions / help messages

Update the following API definition files:

- generators/assets/soracom-api.en.yaml
- generators/assets/soracom-api.ja.yaml
- generators/assets/sandbox/soracom-sandbox-api.en.yaml
- generators/assets/sandbox/soracom-sandbox-api.ja.yaml

To update the message displayed with `configure --help`, update the following files:

- generators/assets/cli/en.yaml
- generators/assets/cli/ja.yaml

## How to build in a local environment (Linux / Mac OS X)

In the environment where Go is installed, run the build script as follows:

```
aws ecr-public get-login-password --profile {your AWS profile} --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
VERSION=1.2.3
./scripts/build.sh $VERSION
```

Here 1.2.3 is the version number. Please specify an appropriate number.

If the build succeeds, then run the test:

```
export SORACOM_AUTHKEY_ID_FOR_TEST=...   # set AuthKey ID & AuthKey of a Soracom operator (account) to use the API sandbox.
export SORACOM_AUTHKEY_FOR_TEST=...
./test/test.sh $VERSION
```

### Troubleshooting Build Issues

If you encounter an error like `go: could not create module cache: mkdir /go/pkg: permission denied` during the build process, please check the permissions of /go/pkg inside the Docker container. In the build.sh script, we mount the host's `${GOPATH:-$HOME/go}` to /go/pkg, so you should check the permissions of `$GOPATH` or `$HOME/go` on the host. In most cases, the following command should resolve the issue:



```bash
sudo chown -R $USER:$USER ${GOPATH:-$HOME/go}
```
