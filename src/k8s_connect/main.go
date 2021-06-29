package main

import (
    "fmt"
    "os"
    "log"
    "os/exec"
)


func main() {
    // Variables
    environment := os.Args[1]
    var aws_account_id string
    var aws_role_name string

    // If condition to set AWS account ID and Role name depending on the environment
    if environment == "sandbox" {
        aws_account_id = "YYYYYYYYYY"
        aws_role_name = "ROLE-NAME"
    } else if environment == "dev" {
        aws_account_id = "XXXXXXXXXX"
        aws_role_name = "ROLE-NAME"
    } else if environment == "qa" {
        aws_account_id = "AAAAAAAAA"
        aws_role_name = "ROLE-NAME"
    } else if environment == "prod" {
        aws_account_id = "BBBBBBBBB"
        aws_role_name = "ROLE-NAME"
    } else {
        fmt.Println("Wrong environment. Please specify one of the options: sandbox, dev, qa or prod.")
    }

    // Setting content of $HOME/.aws/config
    aws_config_content := "[default]\nregion = us-east-1\nsso_start_url = https://SSO_URL/start\nsso_region = us-east-1\nsso_account_id = " + aws_account_id + "\nsso_role_name = " + aws_role_name + "\n"

    // Get the home directory to set the right path
    dirname, err := os.UserHomeDir()
    f, err := os.Create(dirname + "/.aws/config")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    // Write content to aws config file
    _, err2 := f.WriteString(aws_config_content)

    if err2 != nil {
        log.Fatal(err2)
    }

    // Print some information
    fmt.Println("Environment: " + environment)
    fmt.Println("Opening AWS authentication page")

    // Run AWS SSO Login
    sso_login_cmd := exec.Command("aws", "sso", "login")
    sso_login_cmd.Run()

    // Run aws cli to create the kubeconfig file
    kube_config_cmd := exec.Command("/bin/sh", "-c", "aws --profile default eks --region us-east-1 update-kubeconfig --name eks-" + environment + " --kubeconfig ~/.kube/config --alias eks" + environment)
    kube_config_cmd.Run()

    fmt.Println("Exporting configuration")
    fmt.Println("Make sure the env variable KUBECONFIG is set to $HOME/.kube/config")
}
