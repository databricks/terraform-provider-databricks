package main

type certUtilImpl struct{}

func (w certUtilImpl) GetRegisterCertificateCommands(certPath string) [][]string {
	return [][]string{{"certutil", "-addstore", "-f", "ROOT", certPath}}
}

func (w certUtilImpl) GetDeregisterCertificateCommands(certPath string) [][]string {
	return [][]string{{"certutil", "-delstore", "ROOT", certPath}}
}

func newCertUtil() certUtil {
	return certUtilImpl{}
}
