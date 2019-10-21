package bbrpc

// Addnewtemplate https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#addnewtemplate
func (c *Client) Addnewtemplate(p AddnewtemplateParam) (*string, error) {
	// TODO p should be a struct or pointer to struct
	m := map[string]interface{}{
		"type":        p.TemplateType(),
		p.ParamName(): p,
	}
	resp, err := c.sendCmd("addnewtemplate", m)
	if err != nil {
		return nil, err
	}
	var templateAddress string
	err = futureParse(resp, &templateAddress)
	return &templateAddress, err
}

// Createtransaction https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#createtransaction
func (c *Client) Createtransaction(cmd CmdCreatetransaction) (*string, error) {
	resp, err := c.sendCmd("createtransaction", cmd)
	if err != nil {
		return nil, err
	}
	var txdata string
	err = futureParse(resp, &txdata)
	return &txdata, err
}

// Getbalance https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#getbalance
func (c *Client) Getbalance(fork *string, address *string) ([]BalanceInfo, error) {
	if address != nil  && *address == "" {
		address = nil
	}
	resp, err := c.sendCmd("getbalance", struct {
		Fork    *string `json:"fork,omitempty"`
		Address *string `json:"address,omitempty"`
	}{Fork: fork, Address: address})
	if err != nil {
		return nil, err
	}
	var data []BalanceInfo
	err = futureParse(resp, &data)
	return data, err
}

// Getnewkey https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#getnewkey
func (c *Client) Getnewkey(passphrase string) (string, error) {
	resp, err := c.sendCmd("getnewkey", struct {
		Passphrase string `json:"passphrase"`
	}{Passphrase: passphrase})
	if err != nil {
		return "", nil
	}
	var pubkey string
	err = futureParse(resp, &pubkey)
	return pubkey, err
}

// Gettransaction https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#gettransaction
func (c *Client) Gettransaction(txid string, serialized *bool) (*TransactionDetail, error) {
	resp, err := c.sendCmd("gettransaction", struct {
		Txid       string `json:"txid"`
		Serialized *bool  `json:"serialized,omitempty"`
	}{Txid: txid, Serialized: serialized})
	if err != nil {
		return nil, err
	}
	var detail TransactionDetail
	err = futureParse(resp, &detail)
	return &detail, err
}

// Importprivkey https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#importprivkey
func (c *Client) Importprivkey(privkey, passphrase string) (*string, error) {
	resp, err := c.sendCmd("importprivkey", struct {
		Privkey    string `json:"privkey"`
		Passphrase string `json:"passphrase"`
	}{Privkey: privkey, Passphrase: passphrase})
	if err != nil {
		return nil, err
	}
	var pubkey string
	err = futureParse(resp, &pubkey)
	return &pubkey, err
}

// Listaddress https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#listaddress
func (c *Client) Listaddress() ([]AddressData, error) {
	resp, err := c.sendCmd("listaddress", nil)
	if err != nil {
		return nil, err
	}
	var data []AddressData
	err = futureParse(resp, data)
	return data, err
}

// Listkey https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#listkey
func (c *Client) Listkey() ([]PubkeyInfo, error) {
	resp, err := c.sendCmd("listkey", nil)
	if err != nil {
		return nil, err
	}
	var data []PubkeyInfo
	err = futureParse(resp, &data)
	return data, err
}

// Listtransaction https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#listtransaction
func (c *Client) Listtransaction(count *uint, offset *int) ([]Transaction, error) {
	resp, err := c.sendCmd("listtransaction", struct {
		Count  *uint `json:"count,omitempty"`
		Offset *int  `json:"offset,omitempty"`
	}{Count: count, Offset: offset})
	if err != nil {
		return nil, err
	}
	var data []Transaction
	err = futureParse(resp, &data)
	return data, err
}

// CmdSendfrom .
type CmdSendfrom struct {
	To     string   `json:"to"`               //(string, required) to address
	From   string   `json:"from"`             //(string, required) from address
	Amount float64  `json:"amount"`           //(double, required) amount
	Txfee  *float64 `json:"txfee,omitempty"`  //(double, optional) transaction fee
	Fork   *string  `json:"fork,omitempty"`   //(string, optional) fork hash
	Data   *string  `json:"data,omitempty"`   //(string, optional) output data
	SignM  *string  `json:"sign_m,omitempty"` //(string, optional) exchange sign m
	SignS  *string  `json:"sign_s,omitempty"` //(string, optional) exchange sign s
}

// Sendfrom https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#sendfrom
func (c *Client) Sendfrom(cmd CmdSendfrom) (*string, error) {
	resp, err := c.sendCmd("sendfrom", &cmd)
	if err != nil {
		return nil, err
	}
	var txid string
	err = futureParse(resp, &txid)
	return &txid, err
}

// Signtransaction https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#signtransaction
func (c *Client) Signtransaction(txdata string) (*SigntransactionResult, error) {
	resp, err := c.sendCmd("signtransaction", struct {
		Txdata string `json:"txdata"`
	}{Txdata: txdata})
	if err != nil {
		return nil, err
	}
	var ret SigntransactionResult
	err = futureParse(resp, &ret)
	return &ret, err
}

// Validateaddress https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#validateaddress
func (c *Client) Validateaddress(addr string) (*AddressInfo, error) {
	resp, err := c.sendCmd("validateaddress", struct {
		Address string `json:"address"`
	}{Address: addr})
	if err != nil {
		return nil, err
	}
	var info AddressInfo
	err = futureParse(resp, &info)
	return &info, err
}

// Unlockkey https://github.com/bigbangcore/BigBang/wiki/JSON-RPC#unlockkey
func (c *Client) Unlockkey(pubkey, passphrase string, timeout *int64) (*string, error) {
	resp, err := c.sendCmd("unlockkey", struct {
		Pubkey     string `json:"pubkey"`
		Passphrase string `json:"passphrase"`
		Timeout    *int64 `json:"timeout,omitempty"`
	}{Pubkey: pubkey, Passphrase: passphrase, Timeout: timeout})
	if err != nil {
		return nil, err
	}
	var result string
	err = futureParse(resp, &result)
	return &result, err
}
