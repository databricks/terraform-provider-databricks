package main

type certUtil interface {
	GetRegisterCertificateCommands(certPath string) [][]string
	GetDeregisterCertificateCommands(certPath string) [][]string
}
