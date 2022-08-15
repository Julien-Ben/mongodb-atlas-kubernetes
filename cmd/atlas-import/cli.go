package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mongodb/mongodb-atlas-kubernetes/cmd/atlas-import/importer"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// TODO add a debug flag

func generateBaseConfig(cmd *cobra.Command) importer.AtlasImportConfig {
	baseConfig := importer.AtlasImportConfig{
		AtlasDomain:     "https://cloud-qa.mongodb.com",
		ImportNamespace: "default",
	}
	orgID, publicKey, privateKey := getConnectionStrings(cmd)
	baseConfig.PublicKey = publicKey
	baseConfig.PrivateKey = privateKey
	baseConfig.OrgID = orgID
	domain, err := cmd.Flags().GetString("domain")
	if err == nil && domain != "" {
		baseConfig.AtlasDomain = domain
	}
	namespace, err := cmd.Flags().GetString("namespace")
	if err == nil && domain != "" {
		baseConfig.ImportNamespace = namespace
	}
	return baseConfig
}

func getConnectionStrings(cmd *cobra.Command) (string, string, string) {
	orgID, err := cmd.Flags().GetString("org")
	if err != nil {
		importer.Log.Error(err.Error())
		importer.Log.Fatal("Please provide following arg : org")
	}
	publicKey, err := cmd.Flags().GetString("publickey")
	if err != nil {
		importer.Log.Error(err.Error())
		importer.Log.Fatal("Please provide following arg : publickey")
	}
	privateKey, err := cmd.Flags().GetString("privatekey")
	if err != nil {
		importer.Log.Error(err.Error())
		importer.Log.Fatal("Please provide following arg : privatekey")
	}
	return orgID, publicKey, privateKey
}

func parseYAMLFile(filename string) (*importer.AtlasImportConfig, error) {
	absPath := filename
	if !filepath.IsAbs(filename) {
		abs, err := filepath.Abs(filename)
		if err != nil {
			importer.Log.Error("Error with file path, try specifying the absolute path")
			return nil, err
		}
		absPath = abs
	}
	yamlFile, err := os.ReadFile(filepath.Clean(absPath))
	if err != nil {
		importer.Log.Error("Error reading file " + filename + " : " + err.Error())
		return nil, err
	}
	var config importer.AtlasImportConfig
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		importer.Log.Error("Error parsing YAML from file : " + err.Error())
		return nil, err
	}
	return &config, nil
}

func main() {
	importer.Log, _ = zap.NewDevelopment()
	Execute()
}

func run(config importer.AtlasImportConfig) {
	err := importer.RunImports(config)
	if err != nil {
		importer.Log.Fatal(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   "atlas-import",
	Short: "CLI tool to import your atlas resources into Kubernetes",
	Long: `CLI tool to import your atlas resources into Kubernetes. This tool allows you
to either use a configuration file, or CLI arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You called atlas-import")
	},
}

var fromConfigCmd = &cobra.Command{
	Use:     "from-config",
	Aliases: []string{"conf"},
	Short:   "Import resources using a config file",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configFile := args[0]
		importer.Log.Info("Running imports with config file : " + configFile)
		config, err := parseYAMLFile(configFile)
		if err != nil {
			importer.Log.Error(err.Error())
			importer.Log.Fatal("Couldn't read provided configuration file")
		}
		run(*config)
	},
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Import an Atlas Project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		importProjectID := args[0]
		importAllDeployments, _ := cmd.Flags().GetBool("all")
		deploymentsList, _ := cmd.Flags().GetStringSlice("Deployments")
		importer.Log.Info("Importing project with ID : " + importProjectID)
		config := generateBaseConfig(cmd)
		config.ImportAll = false
		config.ImportedProjects = []importer.ImportedProject{
			{
				ID:          importProjectID,
				ImportAll:   importAllDeployments,
				Deployments: deploymentsList,
			},
		}
		run(config)
	},
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Import all your Atlas resources in Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		importer.Log.Info("Importing all resources")
		config := generateBaseConfig(cmd)
		config.ImportAll = true
		run(config)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(fromConfigCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(allCmd)

	//TODO replace authentication method, can store a secret in cluster like for operator
	rootCmd.PersistentFlags().String("org", "", "Your Atlas organization ID")
	rootCmd.PersistentFlags().String("publickey", "", "Your Atlas organization public key")
	rootCmd.PersistentFlags().String("privatekey", "", "Your Atlas organization private key")
	//TODO rename this flag to import-namespace
	rootCmd.PersistentFlags().String("namespace", "", "Kubernetes namespace in which to instantiate resources")
	rootCmd.PersistentFlags().String("domain", "", "Atlas domain name")

	//Will be true if --all is added, and false otherwise
	projectCmd.PersistentFlags().Bool("all", false, "Import all Deployments for given project")
	//Deployments should be specified as : --deployments="dep1,dep2,dep3"
	projectCmd.PersistentFlags().StringSlice("deployments", []string{}, "List of Deployments ID to import for given project")
}
