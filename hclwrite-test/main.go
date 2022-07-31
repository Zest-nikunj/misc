package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

const (
	CREDS = "/opt/creds.json"
)

func main() {
	// create new empty hcl file object
	hclFile := hclwrite.NewEmptyFile()

	// create new file on system
	tfFile, err := os.Create("gcp.spc")
	if err != nil {
		fmt.Println(err)
		return
	}
	// initialize the body of the new file object
	rootBody := hclFile.Body()
	provider := rootBody.AppendNewBlock("connection", []string{"gcp"})
	providerBody := provider.Body()
	providerBody.SetAttributeValue("plugin", cty.StringVal("gcp"))
	providerBody.SetAttributeValue("credentials", cty.StringVal(CREDS))
	providerBody.SetAttributeValue("pagerduty", cty.ObjectVal(map[string]cty.Value{
		"source":  cty.StringVal("PagerDuty/pagerduty"),
		"version": cty.StringVal("1.10.0"),
	}))

	// fmt.Printf("%s", hclFile.Bytes())
	tfFile.Write(hclFile.Bytes())
}
