package gf

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var kcpCmd = &cobra.Command{
	Use:   "kcp",
	Short: "Karabiner change profile",
	Long:  `Usage: gf kcp PROFILE`,
	Run:   Kcp,
	Args:  cobra.ExactArgs(1),
}

func Kcp(cmd *cobra.Command, args []string) {
	newProfile := args[0]
	if newProfile == "" {
		log.Fatalf("Usage: gf kcp PROFILE")
	}

	path := os.Getenv("HOME") + "/.config/karabiner/karabiner.json"
	file, err := os.ReadFile(path)
	cobra.CheckErr(err)

	cfg := map[string]interface{}{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		panic("Failed to unmarshall karabiner.json")
	}
	cobra.CheckErr(err)

	if len(cfg) == 0 {
		panic("karabiner.json length == 0")
	}

	profiles, ok := cfg["profiles"].([]interface{})
	if !ok {
		panic("karabiner.json does not contain profiles")
	}

	newProfileSet := false
	// Iterate over the profiles
	for i := range profiles {
		profileMap, ok := profiles[i].(map[string]interface{})
		if !ok {
			panic("karabiner.json profile is invalid")
		}

		// Set "selected = true" for newProfile
		profileMap["selected"] = false
		if profileMap["name"] == newProfile {
			profileMap["selected"] = true
			newProfileSet = true
		}
		profiles[i] = profileMap
	}
	if !newProfileSet {
		log.Default().Fatalf("Profile '%s' does not exist", newProfile)
	}

	res, err := json.MarshalIndent(cfg, "", "   ")
	cobra.CheckErr(err)

	err = os.WriteFile(path, res, 0600)
	cobra.CheckErr(err)
}

func init() {
	RootCmd.AddCommand(kcpCmd)
}
