package ecdsa

import (
    "crypto/ecdsa"
    "crypto/elliptic"
)

// 设置 PrivateKey
func (this Ecdsa) WithPrivateKey(data *ecdsa.PrivateKey) Ecdsa {
    this.privateKey = data

    return this
}

// 设置 PublicKey
func (this Ecdsa) WithPublicKey(data *ecdsa.PublicKey) Ecdsa {
    this.publicKey = data

    return this
}

// 设置曲线类型
// 可选参数 [P521 | P384 | P256 | P224]
func (this Ecdsa) WithCurve(curve string) Ecdsa {
    switch curve {
        case "P521":
            this.curve = elliptic.P521()
        case "P384":
            this.curve = elliptic.P384()
        case "P256":
            this.curve = elliptic.P256()
        case "P224":
            this.curve = elliptic.P224()
    }

    return this
}

// 设置 data
func (this Ecdsa) WithData(data []byte) Ecdsa {
    this.data = data

    return this
}

// 设置 paredData
func (this Ecdsa) WithParedData(data []byte) Ecdsa {
    this.paredData = data

    return this
}

// 设置 hash 类型
func (this Ecdsa) WithSignHash(hash string) Ecdsa {
    this.signHash = hash

    return this
}

// 设置验证结果
func (this Ecdsa) WithVeryed(data bool) Ecdsa {
    this.veryed = data

    return this
}

// 设置错误
func (this Ecdsa) WithError(err error) Ecdsa {
    this.Error = err

    return this
}