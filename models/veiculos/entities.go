package veiculos

type Veiculo struct {
	Idveiculos int64 `json:"idveiculos"`
	Nome string `json:"nome"`
	Marca string `json:"marca"`
	Ano string `json:"ano"`
	ValorVenda uint64 `json:"valorvenda"`
}

