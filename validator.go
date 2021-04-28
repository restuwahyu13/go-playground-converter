package gpc

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type response struct {
    Message interface{} `json:"msg"`
    Value interface{} `json:"value"`
    Param string `json:"param"`
    Tag string `json:"tag"`
}

type validators interface {
    validator(s interface{}, config map[int]interface{}) map[string]interface{}
}

type validate struct {
    v2 *validator.Validate
}

func NewValidator(v2 *validator.Validate) *validate {
    return &validate{v2: v2}
}

func (v *validate) validator(s interface{}, config map[int]interface{}) map[string]interface{} {

errObject := make(map[string]interface{})

for _, value := range config {
    encode, _ := json.Marshal(value)
    json.Unmarshal(encode, &value)
    mapping := value.(map[string]interface{})

    v.v2 = validator.New()
    err := v.v2.Struct(s)

    if err != nil {
        for _, errResult := range err.(validator.ValidationErrors) {
            switch {
            case mapping["Tag"] == "eqcsfield" && errResult.ActualTag() == "eqcsfield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                    }
            case mapping["Tag"] == "eqfield" && errResult.ActualTag() == "eqfield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "fieldcontains" && errResult.ActualTag() == "fieldcontains" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "fieldexcludes" && errResult.ActualTag() == "fieldexcludes" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "gtcsfield" && errResult.ActualTag() == "gtcsfield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "gtecsfield" && errResult.ActualTag() == "gtecsfield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "gtefield" && errResult.ActualTag() == "gtefield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "gtfield" && errResult.ActualTag() == "gtfield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ltcsfield" && errResult.ActualTag() == "ltcsfield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ltecsfield" && errResult.ActualTag() == "ltecsfield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ltefield" && errResult.ActualTag() == "ltefield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ltfield" && errResult.ActualTag() == "ltfield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "necsfield" && errResult.ActualTag() == "necsfield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "nefield" && errResult.ActualTag() == "nefield" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "cidr" && errResult.ActualTag() == "cidr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "cidrv4" && errResult.ActualTag() == "cidrv4" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "cidrv6" && errResult.ActualTag() == "cidrv6" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "datauri" && errResult.ActualTag() == "datauri" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "fqdn" && errResult.ActualTag() == "fqdn" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "hostname" && errResult.ActualTag() == "hostname" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "hostname_port" && errResult.ActualTag() == "hostname_port" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "hostname_rfc1123" && errResult.ActualTag() == "hostname_rfc1123" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ip" && errResult.ActualTag() == "ip" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ip4_addr" && errResult.ActualTag() == "ip4_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ip6_addr" && errResult.ActualTag() == "ip6_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ip_addr" && errResult.ActualTag() == "ip_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ipv4" && errResult.ActualTag() == "ipv4" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ipv6" && errResult.ActualTag() == "ipv6" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "mac" && errResult.ActualTag() == "mac" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "tcp4_addr	" && errResult.ActualTag() == "tcp4_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "tcp6_addr" && errResult.ActualTag() == "tcp6_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "tcp_addr" && errResult.ActualTag() == "tcp_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "udp4_addr" && errResult.ActualTag() == "udp4_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "udp6_addr" && errResult.ActualTag() == "udp6_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "udp_addr" && errResult.ActualTag() == "udp_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "unix_addr" && errResult.ActualTag() == "unix_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uri" && errResult.ActualTag() == "uri" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "url" && errResult.ActualTag() == "url" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "url_encoded" && errResult.ActualTag() == "url_encoded" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "urn_rfc2141" && errResult.ActualTag() == "urn_rfc2141" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "alpha" && errResult.ActualTag() == "alpha" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "alphanum" && errResult.ActualTag() == "alphanum" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "alphanumunicode" && errResult.ActualTag() == "alphanumunicode" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "alphaunicode" && errResult.ActualTag() == "alphaunicode" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ascii" && errResult.ActualTag() == "ascii" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "contains" && errResult.ActualTag() == "contains" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "containsany" && errResult.ActualTag() == "containsany" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "containsrune" && errResult.ActualTag() == "containsrune" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "endswith" && errResult.ActualTag() == "endswith" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "lowercase" && errResult.ActualTag() == "lowercase" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "multibyte" && errResult.ActualTag() == "multibyte" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "number" && errResult.ActualTag() == "number" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "numeric" && errResult.ActualTag() == "numeric" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "printascii" && errResult.ActualTag() == "printascii" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "startswith" && errResult.ActualTag() == "startswith" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uppercase" && errResult.ActualTag() == "uppercase" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "base64" && errResult.ActualTag() == "base64" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "base64url" && errResult.ActualTag() == "base64url" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "btc_addr" && errResult.ActualTag() == "btc_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "btc_addr_bech32" && errResult.ActualTag() == "btc_addr_bech32" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "datetime" && errResult.ActualTag() == "datetime" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "e164" && errResult.ActualTag() == "e164" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "email" && errResult.ActualTag() == "email" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "eth_addr" && errResult.ActualTag() == "eth_addr" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "hexadecimal" && errResult.ActualTag() == "hexadecimal" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "hexcolor" && errResult.ActualTag() == "hexcolor" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "hexcolor" && errResult.ActualTag() == "hexcolor" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "hsl" && errResult.ActualTag() == "hsl" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "hsla" && errResult.ActualTag() == "hsla" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "html" && errResult.ActualTag() == "html" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "html_encoded" && errResult.ActualTag() == "html_encoded" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "isbn" && errResult.ActualTag() == "isbn" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "isbn10" && errResult.ActualTag() == "isbn10" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "isbn13" && errResult.ActualTag() == "isbn13" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "json" && errResult.ActualTag() == "json" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "latitude" && errResult.ActualTag() == "latitude" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "longitude" && errResult.ActualTag() == "longitude" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "rgb" && errResult.ActualTag() == "rgb" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "rgba" && errResult.ActualTag() == "rgba" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ssn" && errResult.ActualTag() == "ssn" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uuid" && errResult.ActualTag() == "uuid" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uuid3" && errResult.ActualTag() == "uuid3" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uuid3_rfc4122" && errResult.ActualTag() == "uuid3_rfc4122" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uuid4" && errResult.ActualTag() == "uuid4" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uuid4_rfc4122" && errResult.ActualTag() == "uuid4_rfc4122" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uuid5" && errResult.ActualTag() == "uuid5" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uuid5_rfc4122" && errResult.ActualTag() == "uuid5_rfc4122" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "uuid_rfc4122" && errResult.ActualTag() == "uuid_rfc4122" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "eq" && errResult.ActualTag() == "eq" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "gt" && errResult.ActualTag() == "gt" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "gte" && errResult.ActualTag() == "gte" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "lt" && errResult.ActualTag() == "lt" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "lte" && errResult.ActualTag() == "lte" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ne" && errResult.ActualTag() == "ne" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "dir" && errResult.ActualTag() == "dir" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "endswith" && errResult.ActualTag() == "endswith" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "excludes" && errResult.ActualTag() == "excludes" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "excludesall" && errResult.ActualTag() == "excludesall" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "excludesrune" && errResult.ActualTag() == "excludesrune" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "file" && errResult.ActualTag() == "file" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "isdefault" && errResult.ActualTag() == "isdefault" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "len" && errResult.ActualTag() == "len" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "max" && errResult.ActualTag() == "max" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "ne" && errResult.ActualTag() == "ne" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "min" && errResult.ActualTag() == "min" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "oneof" && errResult.ActualTag() == "oneof" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "required" && errResult.ActualTag() == "required" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "required_if" && errResult.ActualTag() == "required_if" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "required_unless" && errResult.ActualTag() == "required_unless" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "required_with" && errResult.ActualTag() == "required_with" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "required_with_all" && errResult.ActualTag() == "required_with_all" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "required_without" && errResult.ActualTag() == "required_without" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "required_without_all" && errResult.ActualTag() == "required_without_all" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "excluded_with" && errResult.ActualTag() == "excluded_with" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "excluded_with_all" && errResult.ActualTag() == "excluded_with_all" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "excluded_without" && errResult.ActualTag() == "excluded_without" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "excluded_without_all" && errResult.ActualTag() == "excluded_without_all" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            case mapping["Tag"] == "unique" && errResult.ActualTag() == "unique" && mapping["Field"] == errResult.StructField():
                errObject[errResult.StructField()] = response{
                    Message: mapping["Message"],
                    Value: errResult.Value(),
                    Param: errResult.StructField(),
                    Tag: errResult.ActualTag(),
                }
            }
        }
    }
}
    return errObject
}