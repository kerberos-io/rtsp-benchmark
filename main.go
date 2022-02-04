package main

import (
    "encoding/json"
    "flag"
    "github.com/dustinkirkland/golang-petname"
    "github.com/k0kubun/go-ansi"
    "github.com/kerberos-io/rtsp-benchmark/api"
    "github.com/kerberos-io/rtsp-benchmark/models"
    "github.com/schollz/progressbar/v3"
    log "github.com/sirupsen/logrus"
    "math/rand"
    "strconv"
    "strings"
    "time"
)

func main() {

    // The required input variables.
    factoryAPI := flag.String("factory-api", "", "the API of Kerberos Factory")
    factoryUsername := flag.String("factory-username", "", "the username of Kerberos Factory")
    factoryPassword := flag.String("factory-password", "", "the password of Kerberos Factory")
    numberOfAgents := flag.Int("agents", 0, "the number of Kerberos Agents to be deployed")
    action := flag.String("action", "create", "create/delete Kerberos Agents based on the deployment key.")
    deployment := flag.String("deployment", "", "the deployment name to be removed.")
    rtsp := flag.String("rtsp", "", "the RTSP connection that will injected into your Kerberos Agent.")
    continuous := flag.String("continuous", "", "whether the agent needs to run in a continuous or motion based recording.")
    region := flag.String("region", "", "if motion based, a region can be setup.")
    flag.Parse()

    if *factoryAPI == "" {
        log.Error("Providing a factory API is required.")
        return
    }
    if *factoryUsername == "" {
        log.Error("Providing the factory username is required.")
        return
    }
    if *factoryPassword == "" {
        log.Error("Providing the factory password is required.")
        return
    }

    // Create a progress bar to visualise the creation
    // of the Kerberos Agents and Virtual RTSP's.

    bar := progressbar.NewOptions(103,
        progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
        progressbar.OptionEnableColorCodes(true),
        progressbar.OptionShowBytes(false),
        progressbar.OptionSetWidth(30),
        progressbar.OptionSetTheme(progressbar.Theme{
            Saucer:        "[green]=[reset]",
            SaucerHead:    "[green]>[reset]",
            SaucerPadding: " ",
            BarStart:      "[",
            BarEnd:        "]",
        }),
    )

    time.Sleep(100 * time.Millisecond)
    log.Info("Contacting: " + *factoryAPI + " with username: " + *factoryUsername + " and password: " + *factoryPassword)
    time.Sleep(500 * time.Millisecond)
    bar.Describe("[cyan][1/4][reset] Login to Kerberos Factory")
    bar.Add(1)

    time.Sleep(500 * time.Millisecond)

    body, err := api.AuthorizeFactory(*factoryAPI, *factoryUsername, *factoryPassword)
    var authorization models.Authorization
    err = json.Unmarshal(body, &authorization)

    if err == nil {
        bar.Describe("[cyan][2/4][reset] Received authorization token")
        bar.Add(1)
        time.Sleep(500 * time.Millisecond)
    } else {
        bar.Describe("[cyan][2/4][reset] Something went wrong while fetching authorization token")
        bar.Finish()
        time.Sleep(500 * time.Millisecond)
        return
    }

    if *action == "create" {

        if *numberOfAgents <= 0 {
            bar.Describe("[red][4/4][reset] One or more agents need to be deployed")
            bar.Finish()
            return
        }
        if *rtsp == "" {
            bar.Describe("[red][4/4][reset] An RTSP connection is required for your Kerberos Agent")
            bar.Finish()
            return
        }

        // Creating a deployment petname
        rand.Seed(time.Now().UnixNano())
        friendlyName := petname.Generate(1, "")

        bar.ChangeMax(*numberOfAgents + 3)
        agentsAdded := 0
        for i := 0; i < *numberOfAgents; i++ {
            bar.Describe("[cyan][3/4][reset] Adding virtual-rtsp Kerberos Agent (" + strconv.Itoa(i) + ")")
            agentName := friendlyName+"-"+strconv.Itoa(i)

            var containerDetails models.ContainerDetails
            containerDetails.Name = agentName
            containerDetails.RTSP = *rtsp
            containerDetails.Continuous = *continuous
            containerDetails.Region = *region

            _, err := api.CreateKerberosAgent(*factoryAPI, authorization.Token, containerDetails)
            if err == nil {
                agentsAdded = agentsAdded + 1
                bar.Add(1)
                time.Sleep(50 * time.Millisecond)
            } else {
                break
            }
        }
        bar.Describe("[cyan][4/4][reset] You are done. " + strconv.Itoa(agentsAdded) + " Kerberos Agents where deployed under deployment key: " + friendlyName)
        bar.Finish()
        return

    } else if *action == "delete" {

        if *deployment == "" {
            bar.Describe("[red][4/4][reset] When deleting make sure you have a deployment key specified (-deployment)")
            bar.Finish()
            return
        }

        body, err := api.GetKerberosAgents(*factoryAPI, authorization.Token)
        var apiResponse models.APIResponse
        err = json.Unmarshal(body, &apiResponse)
        if err == nil {
            kerberosAgents := apiResponse.Data
            // Filter the relevant agents from the deployment key.
            var kerberosAgentsFiltered []models.Deployment
            for _, agent := range kerberosAgents {
                name := agent.DeployName
                if strings.Contains(name, *deployment) {
                    kerberosAgentsFiltered = append(kerberosAgentsFiltered, agent)
                }
            }

            // Iterate over the agents.
            agentsRemoved := 0
            bar.ChangeMax(len(kerberosAgentsFiltered) + 3)
            for _, agent := range kerberosAgentsFiltered {
                bar.Describe("[cyan][3/4][reset] Removing virtual-rtsp Kerberos Agent (" + agent.DeployName + ")")
                _, err := api.DeleteKerberosAgent(*factoryAPI, authorization.Token, agent.DeployName)
                if err == nil {
                    agentsRemoved = agentsRemoved + 1
                    bar.Add(1)
                    time.Sleep(50 * time.Millisecond)
                } else {
                    break
                }
            }

            bar.Describe("[green][4/4][reset] You are done. " + strconv.Itoa(agentsRemoved) + " Kerberos Agents where removed.")
            bar.Finish()
            return
        }
    }

    bar.Describe("[red][4/4][reset] Something went wrong " + err.Error())
    bar.Finish()
    return
}