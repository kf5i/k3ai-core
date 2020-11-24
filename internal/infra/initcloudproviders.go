package infra

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	//Cloud providers GO packages
	"github.com/civo/civogo"
	"github.com/enescakir/emoji"
	"github.com/howeyc/gopass"
)

// const (
// 	apiKey = "7TEH9KPSunoBRxYwyUAMGrf1vpCaVtmF03IQ2068cDXk4sbOgi"
// )

//CloudProviders run specific client configurations based on cloud providers selection
func CloudProviders(osFlavor string, cloudProvider string) {
	var clusterNumNodesInt int
	fmt.Printf("Hold on %v, we are going to install K3s on Civo %v\n", emoji.VulcanSalute, emoji.BuildingConstruction)
	time.Sleep(3 * time.Second)
	fmt.Printf("First we will need some inputs from you... %v\n", emoji.Information)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v Enter your Civo Secret key (We are hiding the terminal, type and press Enter once done): ", emoji.NewMoon)
	apiKey, _ := gopass.GetPasswd()

	client, err := civogo.NewClient(string(apiKey))
	fmt.Printf("%v Enter the name of your cluster: ", emoji.WaxingCrescentMoon)
	clusterName, _ := reader.ReadString('\n')
	clusterName = strings.TrimSpace(strings.ToLower(clusterName))
	fmt.Printf("%v Enter the number of nodes you need: ", emoji.WaxingGibbousMoon)
	clusterNumNodes, _ := reader.ReadString('\n')
	clusterNumNodes = strings.TrimSpace(clusterNumNodes)
	fmt.Printf("%v Enter cluster size (i.e: g2.large): ", emoji.FullMoon)
	clusterSize, _ := reader.ReadString('\n')
	clusterSize = strings.TrimSpace(clusterSize)

	clusterNumNodesInt, _ = strconv.Atoi(clusterNumNodes)

	// client, err := civogo.NewClient(text)
	kc := civogo.KubernetesClusterConfig{Name: clusterName, NumTargetNodes: clusterNumNodesInt, TargetNodesSize: clusterSize}
	fmt.Println("We are creating your cluster..please hold on..")
	client.NewKubernetesClusters(&kc)
	if err != nil {
		fmt.Printf("Failed to create a new config: %s", err)
	} else {
		checkClusterReady(string(apiKey))
	}

}

func checkClusterReady(apiKey string) {
	ticker := time.NewTicker(time.Second * 1).C

	client, _ := civogo.NewClient(apiKey)
	for {
		select {
		case <-ticker:
			instances, err := client.FindKubernetesCluster("k3ai-demo")
			if err != nil {
				log.Fatal(err)
			}
			if instances.Ready != false {
				urlID := string(instances.ID)
				urlID = "https://www.civo.com/account/kubernetes/" + urlID

				fmt.Printf("K3ai installation complete %v%v%v!\n", emoji.PartyPopper, emoji.PartyPopper, emoji.PartyPopper)
				fmt.Printf("To use K3ai grab a copy of KUBECONFIG from the below URL: %v\n", emoji.MechanicalArm)
				fmt.Printf("%v %s\n ", emoji.RightArrow, urlID)
				fmt.Printf("If you need to manage your cluster use the Civo CLI: https://github.com/civo/cli %v\n", emoji.Glasses)
				fmt.Printf("Thank you again for using K3ai, don't forget to check our docs at %v https://docs.k3ai.in\n", emoji.WorldMap)
				time.Sleep(10 * time.Second)
				os.Exit(0)
			}
		}
	}
}
