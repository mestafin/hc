package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/netio"

	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

type VerifyClientController struct {
	bridge   *netio.Bridge
	session  *PairVerifySession
	username string
	LTPK     []byte
	LTSK     []byte
}

func NewVerifyClientController(bridge *netio.Bridge, username string) *VerifyClientController {
	LTPK, LTSK, _ := crypto.ED25519GenerateKey(username)

	controller := VerifyClientController{
		username: username,
		bridge:   bridge,
		session:  NewPairVerifySession(),
		LTPK:     LTPK,
		LTSK:     LTSK,
	}

	return &controller
}

func (c *VerifyClientController) Handle(cont_in common.Container) (common.Container, error) {
	var cont_out common.Container
	var err error

	method := cont_in.GetByte(TLVMethod)

	// It is valid that method is not sent
	// If method is sent then it must be 0x00
	if method != 0x00 {
		return nil, common.NewErrorf("Cannot handle auth method %b", method)
	}

	seq := cont_in.GetByte(TLVSequenceNumber)
	switch seq {
	case VerifyStartRespond:
		cont_out, err = c.handlePairVerifyRespond(cont_in)
	case VerifyFinishRespond:
		cont_out, err = c.handlePairVerifyFinishRespond(cont_in)
	default:
		return nil, common.NewErrorf("Cannot handle sequence number %d", seq)
	}

	return cont_out, err
}

// Client -> Server
// - Public key `A`
func (c *VerifyClientController) InitialKeyVerifyRequest() io.Reader {
	cont_out := common.NewTLV8Container()
	cont_out.SetByte(TLVMethod, 0)
	cont_out.SetByte(TLVSequenceNumber, VerifyStartRequest)
	cont_out.SetBytes(TLVPublicKey, c.session.PublicKey[:])

	fmt.Println("<-     A:", hex.EncodeToString(cont_out.GetBytes(TLVPublicKey)))

	return cont_out.BytesBuffer()
}

// Server -> Client
// - B: server public key
// - encrypted message
//      - username
//      - signature: from server session public key, server name, client session public key
//
// Client -> Server
// - encrypted message
//      - username
//      - signature: from client session public key, server name, server session public key,
func (c *VerifyClientController) handlePairVerifyRespond(cont_in common.Container) (common.Container, error) {
	serverPublicKey := cont_in.GetBytes(TLVPublicKey)
	if len(serverPublicKey) != 32 {
		return nil, common.NewErrorf("Invalid server public key size %d", len(serverPublicKey))
	}

	var otherPublicKey [32]byte
	copy(otherPublicKey[:], serverPublicKey)
	c.session.GenerateSharedKeyWithOtherPublicKey(otherPublicKey)
	c.session.SetupEncryptionKey([]byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))

	fmt.Println("Client")
	fmt.Println("->   B:", hex.EncodeToString(serverPublicKey))
	fmt.Println("     S:", hex.EncodeToString(c.session.SecretKey[:]))
	fmt.Println("Shared:", hex.EncodeToString(c.session.SharedKey[:]))
	fmt.Println("     K:", hex.EncodeToString(c.session.EncryptionKey[:]))

	// Decrypt
	data := cont_in.GetBytes(TLVEncryptedData)
	message := data[:(len(data) - 16)]
	var mac [16]byte
	copy(mac[:], data[len(message):]) // 16 byte (MAC)

	decrypted, err := crypto.Chacha20DecryptAndPoly1305Verify(c.session.EncryptionKey[:], []byte("PV-Msg02"), message, mac, nil)
	if err != nil {
		return nil, err
	}

	decrypted_buffer := bytes.NewBuffer(decrypted)
	tlv_decrypted, err := common.NewTLV8ContainerFromReader(decrypted_buffer)
	if err != nil {
		return nil, err
	}

	username := tlv_decrypted.GetString(TLVUsername)
	signature := tlv_decrypted.GetBytes(TLVEd25519Signature)

	fmt.Println("    Username:", username)
	fmt.Println("   Signature:", hex.EncodeToString(signature))

	// Validate signature
	material := make([]byte, 0)
	material = append(material, c.session.OtherPublicKey[:]...)
	material = append(material, username...)
	material = append(material, c.session.PublicKey[:]...)

	LTPK := c.bridge.PublicKey

	if crypto.ValidateED25519Signature(LTPK, material, signature) == false {
		return nil, common.NewErrorf("Could not validate signature")
	}

	cont_out := common.NewTLV8Container()
	cont_out.SetByte(TLVMethod, 0)
	cont_out.SetByte(TLVSequenceNumber, VerifyFinishRequest)

	tlv_encrypt := common.NewTLV8Container()
	tlv_encrypt.SetString(TLVUsername, c.username)

	material = make([]byte, 0)
	material = append(material, c.session.PublicKey[:]...)
	material = append(material, c.username...)
	material = append(material, c.session.OtherPublicKey[:]...)

	signature, err = crypto.ED25519Signature(c.LTSK, material)
	if err != nil {
		return nil, err
	}

	tlv_encrypt.SetBytes(TLVEd25519Signature, signature)

	encrypted, mac, _ := crypto.Chacha20EncryptAndPoly1305Seal(c.session.EncryptionKey[:], []byte("PV-Msg03"), tlv_encrypt.BytesBuffer().Bytes(), nil)

	cont_out.SetBytes(TLVEncryptedData, append(encrypted, mac[:]...))

	return cont_out, nil
}

// Server -> Client
// - only error ocde (optional)
func (c *VerifyClientController) handlePairVerifyFinishRespond(cont_in common.Container) (common.Container, error) {
	err_code := cont_in.GetByte(TLVStatus_NoError)
	if err_code != 0x00 {
		fmt.Println("Unexpected error %d", err_code)
	}

	return nil, nil
}
