/** *
* * * # DOCUMENTAÇÃO Projeto INSERTTEST
* * *
* * * # OBJETIVO: Facilitar o processo de migração dos testes do robô antigo para o banco do robô novo
* * * # COMO UTILIZAR:
* * * 		1- Acesse a página do robo "http://192.168.2.140:8888/login"
* * *		2- Acesse a página de projeto "MENU > Cadastro > Projeto"
* * *		3- Caso o projeto desejado não exista, crie um novo projeto
* * *		4- Identifique o ID do projeto que deseja incluir testes, esse valor será o idProject
* * *		5- Copie a pasta de testes para o diretório "insertTest/testcase/", APENAS OS TESTES
* * *		6- Altere o valor da variável folderPath para o caminho dos testes a serem inseridos
* * *		7- Execute o projeto INSERTTEST com o comando "go run ."
* * *		8- Preencha o valor de idProject
* * *		9- Preencha o valor de idSchemaSPS
* * *	   10- Verifique o resultado do insert no console.
* * *	   11- Criar Lista de teste no robô
* * * # DICA:
* * * 		1- Os valores de idProject e idSchemaSPS são inicializados com valor padrão
* * *		2- A cada novo projeto, adicionar um Println com id e nome do novo projeto para facilitar futuras interações
* * *		3- Alterar valor padrão da variavel userName
* * *		4- Pode alterar o valor da variavel description do test case de acordo com a demanda da remessa de teste
* * *		5- Em caso de erro, repita o processo apenas com as falhas até que todos sejam inseridos
* * *
* * */

package main

import (
	"fmt"
	"insertTest/database"
	"insertTest/utils"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/sijms/go-ora/v2"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	fmt.Println("✅ Repositorio conectado ao BD")

	// Defina os parâmetros do teste
	userName := "rafael.oliveira@m4sistemas.com.br"
	//folderPath := "testcase/provq" // Caminho da pasta com os arquivos
	//folderPath := "testcase/rest/Hermes" // Caminho da pasta com os arquivos
	folderPath := "testcase/solicitacao" // Caminho da pasta com os arquivos
	//folderPath := "testcase/provq/SIXBELL+VOLTE" // Caminho da pasta com os arquivos

	description := ""
	idProject := "296"  // MassaTesteClaro17/06/25
	idSchemaSPS := "21" // POS
	var successList = make([]string, 0)
	var errorList = make([]string, 0)

	// Usuário confirma que alterou a variavel folderPath
	fmt.Println("❌ ATENÇÃO ❌ 📂 Alterou o valor de folderPath para o diretório desejado?")
	fmt.Print("\n📂 folderPath = " + folderPath + "\n\n📂 Confirma?\t<SIM> Enter \n\t\t<NAO> ctrl + c")
	fmt.Scanln(&folderPath)

	fmt.Println("__________________________________________________________________________")
	fmt.Println("userName default: " + userName)
	fmt.Println("description default: " + description)
	fmt.Println("idProject default: " + idProject)
	fmt.Println("idSchemaSPS default: " + idSchemaSPS)

	// Solicita ao usuário para inserir os valores
	fmt.Println("__________________________________________________________________________")
	fmt.Println("				  116  InsertTest")
	fmt.Println("	->			  296  MassaTesteClaro17/06/25")
	fmt.Println("				  76   SequenciaEda+VOLTE")
	fmt.Println("				  201  SequenciaEda")
	fmt.Println("				  136  VNPSIX+VOLTE")
	fmt.Println("				  137  MAGNOLIA")
	fmt.Println("				  138  RSA")
	fmt.Println("				  156  ESIM")
	fmt.Println("				  203  VOLTE")
	fmt.Println("				  277  Hermes")
	fmt.Print(" -> Digite o idProject: ")
	fmt.Scanln(&idProject)

	// Exibe a tabela com as opções para o idSchemaSPS
	fmt.Println("__________________________________________________________________________")
	fmt.Println("  	  1 MVN")
	fmt.Println("  	  2 NAC")
	fmt.Println("->	 21 POS")
	fmt.Println("  	101 FLX")
	fmt.Println("  	102 PRE")
	fmt.Println("  	103 HUB")

	// Solicita o idSchemaSPS
	fmt.Print(" -> Digite o idSchemaSPS: ")
	fmt.Scanln(&idSchemaSPS)
	fmt.Println("__________________________________________________________________________")

	// Confirma execução
	fmt.Print("\n📂 folderPath = " + folderPath + "\n\n📂 Executar insert:\t<SIM> Enter \n\t\t\t<NAO> ctrl + c")
	fmt.Scanln()
	fmt.Println("__________________________________________________________________________")

	// Lê todos os arquivos da pasta
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(fmt.Errorf("Erro ao ler a pasta: %v", err))
	}

	// Percorre cada arquivo na pasta
	for _, file := range files {
		if file.IsDir() {
			continue // Ignora diretórios
		}

		// Monta o caminho completo do arquivo
		filePath := filepath.Join(folderPath, file.Name())
		fmt.Println("--")
		fmt.Printf("📂 Processando arquivo: %s\n", filePath)

		// Processa o arquivo
		testName, processedSQL, idType, err := utils.ProcessTestFile(filePath)
		if err != nil {
			log.Printf("❌ Erro ao processar %s: %v", filePath, err)
			continue
		}

		testName = "TESTCASE"
		fmt.Println("✅ Test file processado:", testName)

		// Processa Insert
		insertSolicitacao := utils.ProcessInsert(processedSQL)

		// Regex para acao
		re := regexp.MustCompile(`SRV_TRX_TP_CD=([^;]+);`)
		// Regex para IMSI
		reIMSI := regexp.MustCompile(`IMSI=([^;]+);`)
		// Regex para MSISDN
		reMSISDN := regexp.MustCompile(`MSISDN=([^;]+);`)
		// Nova Regex para HHUA, HLR ou HLREDA
		reHLX := regexp.MustCompile(`(HHUA|HLREDA|HLR)=([^;]+);`)
		// Regex para VOLTE (case-insensitive) - \b garante que é a palavra completa
		reVOLTE := regexp.MustCompile(`(?i)\bVOLTE\b`) // `(?i)` para case-insensitive
		// Regex para VPNSIX (case-insensitive)
		reVPNSIX := regexp.MustCompile(`(?i)\bVPNSIX\b`)

		nameCounts := make(map[string]int)

		//caso teste sps_solicitacao, idType = 3
		if idType == "3" {
			for _, script := range insertSolicitacao {
				var baseName string
				var imsiExtracted string
				var msisdnExtracted string
				var hlxExtracted string // Variável para armazenar o valor de HHUA/HLR/HLREDA
				var volteFound bool     // Variável para indicar se VOLTE foi encontrado
				var vpnsixFound bool    // Variável para indicar se VPNSIX foi encontrado

				matches := re.FindStringSubmatch(script)
				if len(matches) > 1 {
					baseName = matches[1]
				} else {
					baseName = testName
				}
				// Incrementa a contagem para este nome específico no map
				nameCounts[baseName]++
				currentCount := nameCounts[baseName]
				// Constrói o nome final com a contagem específica
				finalName := fmt.Sprintf("%s_%d", baseName, currentCount)
				//-------------------------------------------------------------------------------
				// --- Extração do IMSI ---
				matchesIMSI := reIMSI.FindStringSubmatch(script)
				if len(matchesIMSI) > 1 {
					fullIMSI := matchesIMSI[1]
					// Garante que pegamos apenas os primeiros 7 caracteres, se existirem
					if len(fullIMSI) >= 7 {
						imsiExtracted = fullIMSI[:7]
					} else {
						imsiExtracted = fullIMSI // Se for menor que 7, pega tudo
					}
				} else {
					imsiExtracted = "" // Se não encontrar, deixa vazio
				}

				// --- Extração do MSISDN ---
				matchesMSISDN := reMSISDN.FindStringSubmatch(script)
				if len(matchesMSISDN) > 1 {
					msisdnExtracted = matchesMSISDN[1]
				} else {
					msisdnExtracted = "" // Se não encontrar, deixa vazio
				}

				// --- Extração de HHUA/HLR/HLREDA ---
				matchesHLX := reHLX.FindStringSubmatch(script)
				if len(matchesHLX) > 2 {
					hlxFieldName := matchesHLX[1] // Ex: "HHUA", "HLR", "HLREDA"
					hlxValue := matchesHLX[2]     // Ex: "ValorHHUA1"
					// Apenas alteramos a forma como hlxExtracted é formatado
					hlxExtracted = fmt.Sprintf("%s=%s", hlxFieldName, hlxValue)
				} else {
					hlxExtracted = ""
				}

				// --- Verificação VOLTE E VPNSIX ---
				// Usa FindString para verificar a existência, não precisa de submatches
				if reVOLTE.FindString(script) != "" {
					volteFound = true
				}

				if reVPNSIX.FindString(script) != "" {
					vpnsixFound = true
				}

				// --- Construção da Description ---
				// Vamos concatenar apenas se os valores foram encontrados
				description := ""
				parts := []string{}
				if imsiExtracted != "" {
					parts = append(parts, "IMSI="+imsiExtracted)
				}
				if msisdnExtracted != "" {
					parts = append(parts, "MSISDN="+msisdnExtracted)
				}
				if hlxExtracted != "" { // Adiciona o campo HLX se for encontrado
					parts = append(parts, hlxExtracted)
				}
				if volteFound {
					parts = append(parts, "VOLTE") // Adiciona a string "VOLTE"
				}
				if vpnsixFound {
					parts = append(parts, "VPNSIX") // Adiciona a string "VPNSIX"
				}

				// Junta as partes com um separador, se houver mais de uma
				description = strings.Join(parts, ", ")

				//-------------------------------------------------------------------------------
				fmt.Println("----INICIO SIMULA INSERT--------------------------------")
				utils.SimulaInsert(idType, finalName, description, script, userName, idProject, idSchemaSPS)
				fmt.Println("----FIM SIMULA INSERT-----------------------------------")
				//-------------------------------------------------------------------------------

				/*
					//-------------------------------------------------------------------------------
					// Inserir no banco
					err = database.InsertTestCase(db, idType, finalName, script, description, userName, idProject, idSchemaSPS)
					if err != nil {
						log.Printf("❌ Erro ao inserir %s no banco: %v", finalName, err)
						errorList = append(errorList, finalName) // Adiciona à lista de erros
						continue
					}
					//-------------------------------------------------------------------------------
				*/

				fmt.Println("✅ Test case inserido no banco:", finalName)
				successList = append(successList, finalName) // Adiciona à lista de sucesso

			}
		}
	}

	fmt.Println("🚀 Processamento concluído para todos os arquivos da pasta!")

	// Exibe os resultados compilados
	fmt.Println("\n### Resultados do insert no banco ###")
	fmt.Printf("Insert com sucesso:\n")
	for _, success := range successList {
		fmt.Println("✅", success)
	}

	if len(errorList) > 0 {
		fmt.Printf("\nInsert com erro:\n")
		for _, failure := range errorList {
			fmt.Println("❌", failure)
		}
	} else {
		fmt.Printf("\n ❌ Nenhum Insert com erro ❌\n")
	}

	fmt.Println("Processamento concluído!")

}
