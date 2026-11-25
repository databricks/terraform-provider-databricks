package main

type certUtilImpl struct{}

func (l certUtilImpl) GetRegisterCertificateCommands(certPath string) [][]string {
	return [][]string{{"sudo", "cp", certPath, "/usr/local/share/ca-certificates/proxy.crt"}, {"sudo", "update-ca-certificates"}}
}

func (l certUtilImpl) GetDeregisterCertificateCommands(certPath string) [][]string {
	return [][]string{{"sudo", "rm", "/usr/local/share/ca-certificates/proxy.crt"}, {"sudo", "update-ca-certificates"}}
}

func newCertUtil() certUtil {
	return certUtilImpl{}
}
