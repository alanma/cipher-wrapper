package crypto_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/89hmdys/toast/crypto"
	"github.com/89hmdys/toast/cipher"
)

func Test_DES_ECB(t *testing.T) {

	cipher, err := crypto.NewDES([]byte("Z'{ru/^e"))
	if err != nil {
		t.Error(err)
		return
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_DES_CBC ok")
}

func Test_DES_CBC(t *testing.T) {

	mode := cipher.NewCBCMode()
	cipher, err := crypto.NewDESWith([]byte("Z'{ru/^e"), mode) //创建一个des 加密的builder
	if err != nil {
		t.Error(err)
		return
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)

	fmt.Println(string(pp))

	fmt.Println("Test_DES_ECB ok")
}

func Test_DES_CFB(t *testing.T) {

	mode := cipher.NewCFBMode()

	cipher, err := crypto.NewDESWith([]byte("Z'{ru/^e"), mode)
	if err != nil {
		t.Error(err)
		return
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_DES_CFB ok")
}

func Test_DES_OFB(t *testing.T) {

	mode := cipher.NewOFBMode()

	cipher, err := crypto.NewDESWith([]byte("Z'{ru/^e"), mode)
	if err != nil {
		t.Error(err)
		return
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_DES_OFB ok")
}

func Test_DES_CTR(t *testing.T) {

	mode := cipher.NewCTRMode()

	cipher, err := crypto.NewDESWith([]byte("Z'{ru/^e"), mode)
	if err != nil {
		t.Error(err)
		return
	}

	plant := `您好！如果您要入手广汽传祺GS5！它所搭载就是这款7速g-dct手自一体变速箱！网上有说好的也要说不好的！我给你一个中肯的建议！首先这款车是一款新车，我把它归纳为一款小众的车！这款7速双离合变速箱也是新款！都有待市场考验！如果您入手了！这款变速箱后期维修和保养可是比丰田，本田，大众，日产，马自达这类常见车的维修和保养成本高太多了！因为配件比较难找！我还是不建议入手！请谨慎考虑！`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_DES_CTR ok")
}
