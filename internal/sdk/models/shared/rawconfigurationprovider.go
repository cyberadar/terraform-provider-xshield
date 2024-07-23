package shared

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type ConfigurationProvider interface {
	TenancyId() string
	FingerPrint() string
	PrincipalId() string
	KeyLoader() PrivateKeyLoader
}

type PrivateKeyLoader interface {
	LoadKey() (*rsa.PrivateKey, error)
}

type privateKeyFromFile struct {
	filePath   string
	passphrase string
}

func NewPrivateKeyFromFileLoader(ctx context.Context, path, passphrase string) PrivateKeyLoader {
	return privateKeyFromFile{filePath: path, passphrase: passphrase}
}

func (pkff privateKeyFromFile) LoadKey() (*rsa.PrivateKey, error) {

	if pkff.filePath == "" {
		return nil, fmt.Errorf("invalid file path for private key: %s", pkff.filePath)
	}

	priv, err := os.ReadFile(pkff.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key. Path: %s, Error: %v", pkff.filePath, err.Error())
	}

	return privateKeyFromBytesWithPassword(priv, []byte(pkff.passphrase))
}

type privateKeyFromByteSlice struct {
	content []byte
}

func NewPrivateKeyFromByteSlice(content string) PrivateKeyLoader {
	return privateKeyFromByteSlice{content: []byte(content)}
}

func (pkfs privateKeyFromByteSlice) LoadKey() (*rsa.PrivateKey, error) {

	if len(pkfs.content) == 0 {
		return nil, fmt.Errorf("byte slice empty. No key content available")
	}

	return privateKeyFromBytesWithPassword(pkfs.content, []byte(""))
}

func privateKeyFromBytesWithPassword(pemData, _ []byte) (key *rsa.PrivateKey, e error) {
	if pemBlock, _ := pem.Decode(pemData); pemBlock != nil {
		// decrypted := pemBlock.Bytes
		// if x509.IsEncryptedPEMBlock(pemBlock) {
		// 	if password == nil {
		// 		e = fmt.Errorf("private key password is required for encrypted private keys")
		// 		return
		// 	}
		// 	if decrypted, e = x509.DecryptPEMBlock(pemBlock, password); e != nil {
		// 		return
		// 	}
		// }

		key, e = parsePKCSPrivateKey(pemBlock.Bytes)

	} else {
		e = fmt.Errorf("PEM data was not found in buffer")
		return
	}
	return
}

func parsePKCSPrivateKey(decryptedKey []byte) (*rsa.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(decryptedKey); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(decryptedKey); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey:
			return key, nil
		default:
			return nil, fmt.Errorf("unsupportesd private key type in PKCS8 wrapping")
		}
	}
	return nil, fmt.Errorf("failed to parse private key")
}

type RawConfigProvider struct {
	tenantId    string
	principalId string
	fingerprint string
	loader      PrivateKeyLoader
}

func (rcp RawConfigProvider) TenancyId() string           { return rcp.tenantId }
func (rcp RawConfigProvider) PrincipalId() string         { return rcp.principalId }
func (rcp RawConfigProvider) FingerPrint() string         { return rcp.fingerprint }
func (rcp RawConfigProvider) KeyLoader() PrivateKeyLoader { return rcp.loader }

func NewRawConfigProvider(tenancyId, userId, fingerprint, keyLocation string) RawConfigProvider {
	return RawConfigProvider{
		tenantId:    tenancyId,
		principalId: userId,
		fingerprint: fingerprint,
		loader:      NewPrivateKeyFromFileLoader(context.Background(), keyLocation, ""),
	}
}

func NewDirectConfigProvider(tenancyId, userId, fingerprint, privateKey string) RawConfigProvider {
	return RawConfigProvider{
		tenantId:    tenancyId,
		principalId: userId,
		fingerprint: fingerprint,
		loader:      NewPrivateKeyFromByteSlice(privateKey),
	}
}
