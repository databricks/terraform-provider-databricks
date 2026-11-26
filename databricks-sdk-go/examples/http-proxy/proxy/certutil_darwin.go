package main

import "os"

type certUtilImpl struct {
	homeDir string
}

func (d certUtilImpl) GetRegisterCertificateCommands(certPath string) [][]string {
	return [][]string{{"security", "add-trusted-cert", "-r", "trustRoot", "-k", d.homeDir + "/Library/Keychains/login.keychain-db", certPath}}
}

func (d certUtilImpl) GetDeregisterCertificateCommands(certPath string) [][]string {
	return [][]string{{"security", "remove-trusted-cert", certPath}}
}

func newCertUtil() certUtil {
	homeDir := os.Getenv("HOME")
	return certUtilImpl{homeDir: homeDir}
}
