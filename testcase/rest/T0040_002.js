{
   "ordem":{
      "correlacao":[
         {
            "id":"VPNSIX_T0040_002",
            "sistema-origem":{
               "id":"SGIOT"
            },
            "endereco-resposta":"http://10.18.81.219:8004/FachadaWoaSgiot/rest"
         }
      ],
      "operacao":{
         "id":"NEW_SUB_ACTIVATION",
         "situacao-operacao":{
            "data":"2023-05-30T08:00:10-03:00"
         },
         "motivo":{
            "id":"PROVISIONING"
         },
         "usuario":{
            "id":"SGIOT"
         },
         "prioridade":"10"
      },
      "cliente":{
         "id":"111008173"
      },
      "item-ordem":[
         {
            "produto-alvo":{
               "recurso-telefonia":{
                  "numeracao":{
                     "numero-telefone":"11988418080"
                  },
                  "simcard":{
                     "iccid":"89550539120001678791",
                     "imsi":"724051215563881",
                     "pin":"3636",
                     "puk":"3636",
                     "pin2":"3636",
                     "puk2":"3636",
                     "ki":"259146167291242D4F5127D9AC5CB8A5",
                     "tk":"178",
                     "chv5":"761446455758488D"
                  },
                  "perfil-aprovisionamento":{
                     "id":"30",
                     "comando-aprovisionamento":[
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
                                    "valor": ""
                                },
                                {
                                    "nome": "HSSDRA",
                                    "valor": "HSSRJ3MG3"
                                }
                            ]
                        },
                        {
                           "servico":{
                              "id":"HSS"
                           },
                           "operacao":{
                              "id":"ACT"
                           }
                        },
                        {
                           "servico":{
                              "id":"HSS"
                           },
                           "operacao":{
                              "id":"CAN"
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
