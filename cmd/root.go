/*
TODO LICENSE
Copyright © 2019 victor23d <victor6742x@gmail.com>
This file is part of {{ .appName }}.
// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ctlv3 contains the main entry point for the etcdctl for v3 API.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.etcd.io/etcd/etcdctl/ctlv3/command"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	cliName        = "etcdctl"
	cliDescription = "A simple command line client for etcd3."

	defaultDialTimeout      = 2 * time.Second
	defaultCommandTimeOut   = 5 * time.Second
	defaultKeepAliveTime    = 2 * time.Second
	defaultKeepAliveTimeOut = 6 * time.Second
)

var cfgFile string

var (
	globalFlags = command.GlobalFlags{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "etcd-batch",
	Short: "batch json/yaml/toml/envfile into etcd",
	Long:  `It's easy to use`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("prefix", "n", "", "a prefix string insert before the key start")

	// From "go.etcd.io/etcd/etcdctl/ctlv3/ctl.go v3.3.15 TODO upgrade 3.4 username password flag"

	rootCmd.PersistentFlags().StringSliceVar(&globalFlags.Endpoints, "endpoints", []string{"127.0.0.1:2379"}, "gRPC endpoints")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.Debug, "debug", false, "enable client-side debug logging")

	rootCmd.PersistentFlags().StringVarP(&globalFlags.OutputFormat, "write-out", "w", "simple", "set the output format (fields, json, protobuf, simple, table)")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.IsHex, "hex", false, "print byte strings as hex encoded strings")

	rootCmd.PersistentFlags().DurationVar(&globalFlags.DialTimeout, "dial-timeout", defaultDialTimeout, "dial timeout for client connections")
	rootCmd.PersistentFlags().DurationVar(&globalFlags.CommandTimeOut, "command-timeout", defaultCommandTimeOut, "timeout for short running command (excluding dial timeout)")
	rootCmd.PersistentFlags().DurationVar(&globalFlags.KeepAliveTime, "keepalive-time", defaultKeepAliveTime, "keepalive time for client connections")
	rootCmd.PersistentFlags().DurationVar(&globalFlags.KeepAliveTimeout, "keepalive-timeout", defaultKeepAliveTimeOut, "keepalive timeout for client connections")

	// TODO: secure by default when etcd enables secure gRPC by default.
	rootCmd.PersistentFlags().BoolVar(&globalFlags.Insecure, "insecure-transport", true, "disable transport security for client connections")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.InsecureDiscovery, "insecure-discovery", true, "accept insecure SRV records describing cluster endpoints")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.InsecureSkipVerify, "insecure-skip-tls-verify", false, "skip server certificate verification")
	rootCmd.PersistentFlags().StringVar(&globalFlags.TLS.CertFile, "cert", "", "identify secure client using this TLS certificate file")
	rootCmd.PersistentFlags().StringVar(&globalFlags.TLS.KeyFile, "key", "", "identify secure client using this TLS key file")
	rootCmd.PersistentFlags().StringVar(&globalFlags.TLS.CAFile, "cacert", "", "verify certificates of TLS-enabled secure servers using this CA bundle")
	rootCmd.PersistentFlags().StringVar(&globalFlags.User, "user", "", "username[:password] for authentication (prompt if password is not supplied)")
	rootCmd.PersistentFlags().StringVarP(&globalFlags.TLS.ServerName, "discovery-srv", "d", "", "domain name to query for SRV records describing cluster endpoints")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".etcd-batch" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".etcd-batch")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.EnablePrefixMatching = true
}
