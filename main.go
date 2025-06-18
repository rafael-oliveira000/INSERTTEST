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
	idProject := "116"  // insertTest
	idSchemaSPS := "21" // POS
	//	var successList []string
	//	var errorList []string

	// Usuário confirma que alterou a variavel folderPath
	fmt.Println("❌ ATENÇÃO ❌ 📂 Alterou o valor de folderPath para o diretório desejado?")
	fmt.Print("\n📂 folderPath = " + folderPath + "\n\n📂 Confirma?\t<SIM> Enter \n\t\t<NAO> ctrl + c")
	//	fmt.Scanln(&folderPath)

	fmt.Println("__________________________________________________________________________")
	fmt.Println("userName default: " + userName)
	fmt.Println("description default: " + description)
	fmt.Println("idProject default: " + idProject)
	fmt.Println("idSchemaSPS default: " + idSchemaSPS)

	// Solicita ao usuário para inserir os valores
	fmt.Println("__________________________________________________________________________")
	fmt.Println("	->			  116  InsertTest")
	fmt.Println("				  76   SequenciaEda+VOLTE")
	fmt.Println("				  201  SequenciaEda")
	fmt.Println("				  136  VNPSIX+VOLTE")
	fmt.Println("				  137  MAGNOLIA")
	fmt.Println("				  138  RSA")
	fmt.Println("				  156  ESIM")
	fmt.Println("				  203  VOLTE")
	fmt.Println("				  277  Hermes")
	fmt.Print(" -> Digite o idProject: ")
	//	fmt.Scanln(&idProject)

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
	//	fmt.Scanln(&idSchemaSPS)
	fmt.Println("__________________________________________________________________________")

	// Confirma execução
	fmt.Print("\n📂 folderPath = " + folderPath + "\n\n📂 Executar insert:\t<SIM> Enter \n\t\t\t<NAO> ctrl + c")
	//	fmt.Scanln()
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
		fmt.Println("✅ Test file processado:", testName)

		// Processa Insert
		insertSolicitacao := utils.ProcessInsert(processedSQL)

		//caso teste sps_solicitacao, idType = 3
		if idType == "3" {
			for i, script := range insertSolicitacao {
				Name := fmt.Sprintf("%s_%d", testName, i)

				fmt.Println("----FIM SIMULA INSERT-----------------------------------")
				utils.SimulaInsert(idType, Name, script, userName, idProject, idSchemaSPS)
				fmt.Println("----FIM SIMULA INSERT-----------------------------------")

				fmt.Println("insertSolicitacao processado.")

				/*
					// Inserir no banco
					err = database.InsertTestCase(db, idType, testName, processedSQL, description, userName, idProject, idSchemaSPS)
					if err != nil {
						log.Printf("❌ Erro ao inserir %s no banco: %v", testName, err)
						errorList = append(errorList, testName) // Adiciona à lista de erros
						continue
					}
					fmt.Println("✅ Test case inserido no banco:", testName)
					successList = append(successList, testName) // Adiciona à lista de sucesso
				*/
			}
		}
	}

	fmt.Println("🚀 Processamento concluído para todos os arquivos da pasta!")
	/*
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
	*/
}
