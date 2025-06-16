{
    "ordem": {
        "cliente": {
            "id": "NUBANK"
        },
        "correlacao": [
            {
                "id": "T0010_001",
                "sistema-origem": {
                    "id": "ZUP"
                },
                "endereco-resposta": "https://test.apigw.claro.com.br/mvno/v1/provisioningordersnotifications/claro_ac76a7739985cdacad94eecd7f04ff64a97e0e93_6062f134-b4b1-41db-98ad-c3b289fed970"
            }
        ],
        "operacao": {
            "id": "CANSUB",
            "situacao-operacao": {
                "data": "2024-04-09T11:40:58.630+00:00"
            },
            "motivo": {
                "id": "NOVO CLIENTE"
            },
            "usuario": {
                "id": "user"
            },
            "prioridade": "10"
        },
        "item-ordem": [
            {
                "produto-alvo": {
                    "recurso-telefonia": {
                        "numeracao": {
                            "numero-telefone": "11991876004"
                        },
                        "simcard": {
                            "iccid": "89550537110010749882",
                            "imsi": "724051190137118",
                            "pin": "3636",
                            "puk": "42424242",
                            "pin2": "6363",
                            "puk2": "24242424",
                            "ki": "eH55eOOOOOOOOOOOaFOOOOOOOOOOOOOOIf",
                            "tk": "123",
                            "chv5": "3434311111111117",
                            "op": "14"
                        },
                        "perfil-aprovisionamento": {
                            "id": "35",
                            "comando-aprovisionamento": [
                                {
                                    "servico": {
                                        "id": "EQPT"
                                    },
                                    "operacao": {
                                        "id": "ACT"
                                    },
                                    "parametro": [
                                        {
                                            "nome": "HLR",
                                            "valor": "GHLAM01"
                                        }
                                    ]
                                },
                                {
                                    "servico": {
                                        "id": "HSS"
                                    },
                                    "operacao": {
                                        "id": "ACT"
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        ]
    }
}