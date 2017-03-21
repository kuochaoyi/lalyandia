package protocol

import (
	"errors"

	"golang.org/x/crypto/blowfish"
)

// ScrambleMod crypt N bytes of RSA key
func ScrambleMod(n []byte) {
	// step 1 : 0x4d-0x50 <-> 0x00-0x04
	for i := 0; i < 4; i++ {
		temp := n[i]
		n[i] = n[0x4d+i]
		n[0x4d+i] = temp
	}

	// step 2 : xor first 0x40 bytes with last 0x40 bytes
	for i := 0; i < 0x40; i++ {
		n[i] = (byte)(n[i] ^ n[0x40+i])
	}

	// step 3 : xor bytes 0x0d-0x10 with bytes 0x34-0x38
	for i := 0; i < 4; i++ {
		n[0x0d+i] = (byte)(n[0x0d+i] ^ n[0x34+i])
	}

	// step 4 : xor last 0x40 bytes with first 0x40 bytes
	for i := 0; i < 0x40; i++ {
		n[0x40+i] = (byte)(n[0x40+i] ^ n[i])
	}
}

/*
func UnscrambleMod( n []byte )
{
      typedef unsigned char byte;
      int i;

      // step 4 xor last 0x40 bytes with first 0x40 bytes
      for( i=0; i<0x40; i++ ) {
            n[0x40 + i] = (byte)(n[0x40 + i] ^ n[i]);
      };

      // step 3 xor bytes 0x0d-0x10 with bytes 0x34-0x38
      for( i=0; i<4; i++ ) {
            n[0x0d + i] = (byte)(n[0x0d + i] ^ n[0x34 + i]);
      };

      // step 2 xor first 0x40 bytes with last 0x40 bytes
      for( i=0; i<0x40; i++ ) {
            n[i] = (byte)(n[i] ^ n[0x40 + i]);
      };

      for( i=0; i<4; i++ ) {
            byte temp = n[0x00 + i];
            n[0x00 + i] = n[0x4d + i];
            n[0x4d + i] = temp;
      };
}
*/

func BlowfishDecrypt(encrypted, key []byte) ([]byte, error) {
	cipher, err := blowfish.NewCipher(key)

	if err != nil {
		return nil, errors.New("Couldn't initialize the blowfish cipher")
	}

	// Check if the encrypted data is a multiple of our block size
	if len(encrypted)%8 != 0 {
		return nil, errors.New("The encrypted data is not a multiple of the block size")
	}

	count := len(encrypted) / 8

	decrypted := make([]byte, len(encrypted))

	for i := 0; i < count; i++ {
		cipher.Decrypt(decrypted[i*8:], encrypted[i*8:])
	}

	return decrypted, nil
}

func BlowfishEncrypt(decrypted, key []byte) ([]byte, error) {
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, errors.New("Couldn't initialize the blowfish cipher")
	}

	// Check if the decrypted data is a multiple of our block size
	if len(decrypted)%8 != 0 {
		return nil, errors.New("The decrypted data is not a multiple of the block size")
	}

	count := len(decrypted) / 8

	encrypted := make([]byte, len(decrypted))

	for i := 0; i < count; i++ {
		cipher.Encrypt(encrypted[i*8:], decrypted[i*8:])
	}

	return encrypted, nil
}
