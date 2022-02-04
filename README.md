# RTSP Benchmark

This RTSP benchmarking tool allows you to automate the deployment of Kerberos Agents, and to benchmark the current hardware you are using or considering to use.

By deploying a number of Kerberos Agents in bulk you will benefit from a better understanding of how many Kerberos Agents your hardware can handle, considering different scenarios:

- motion detection
- continuous recording
- different resolutions: 640x480, 1280x720, etc.
- and more.

![Create Kerberos Agents](/images/create-agents.gif)

## How it works

This tool requires a couple of parameters, and will use your Kerberos Factory installation, and its related APIs, to automate the creation of one or more agents. The required options are:

| Option                                        | Description                                                         | Value |
| --------------------------------------------- | ------------------------------------------------------------------- | ----- |
| `-factory-api`                                | The API of Kerberos Factory                                         | `""`  |
| `-factory-username`                           | The username of Kerberos Factory                                    | `""`  |
| `-factory-password`                           | The password of Kerberos Factory                                    | `""`  |
| `-agents`                                     | The number of Kerberos Agents to be deployed                        | `0`  |
| `-action`                                     | Create/delete Kerberos Agents based on the deployment key.          | `""`  |
| `-rtsp`                                       | The RTSP connection that will injected into your Kerberos Agent.    | `""`  |
| `-deployment`                                 | The deployment name to be removed.                                  | `""`  |
| `-continuous`                                 | Enabling continuous recording (true), or motion based recording (false).                                  | `""`  |
| `-region`                                     | If motion recording enabled, we can set a region (x1-y1,x2-y2,..)   | `""`  |

## How to use

Make sure you have Golang installed on your machine, and that you have access to the Kerberos Factory api. Once you have that dependency ready, you can simply execute the `go run` or `go build` command to execute the benchmark or build the benchmark binary.

    go run main.go -factory-api=http://api...

Read below for more information, how to execute the benchmark tool.

## Creating new Kerberos Agents

Below command will create 5 Kerberos Agents connected to the [Big Buck Bunny](https://www.wowza.com/developer/rtsp-stream-test) rtsp stream (you can replace this by your own).

    go run main.go \
    -factory-api=http://api.factory.domain2.com \
    -factory-username=root \
    -factory-password=kerberos \
    -action=create \
    -agents=5 \
    -rtsp=rtsp://wowzaec2demo.streamlock.net

## Delete existing Kerberos Agents

Kerberos Agents previously created can be removed by defining the `-deployment` value. That was generated while creating your Kerberos Agents in the previous step. 

    go run main.go \
    -factory-api=http://api.factory.domain2.com \
    -factory-username=root \
    -factory-password=kerberos \
    -action=delete \
    -deployment=kingfish

## Use for bulk deployment (WIP).

Despite being a benchmarking tool, it could also be used for the creation of your entire video landscape. By defining a couple of tweaks you could reuse the majority of this repository, to read a CSV file with (name, rtsp) information, and create the corresponding Kerberos Agent.

This feature is not yet implemented, so PRs or ideas are welcome.