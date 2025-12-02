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
* * * # ATENÇÃO:
* * * 		1- Esse projeto remove as substring "(A)" e "(CRT)" dos inserts de SPS_SOLICITACAO em utils.ProcessSolicitacao para evitar erros de sintaxe
* * * 				FTRCD=VPNST;FTR_STATUS=A;CC=91446740;CCNM=SR(A) Marcelo Rosa da Silva;DPTO1=91446740;DPTONM1=SR(A) Marcelo Rosa da Silva;@
* * *				NOM_OPER=VIVO - RS (CRT);@
* * *
* * */

package main

import (
	"fmt"
	"insertTest/database"
	"insertTest/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"

	_ "github.com/sijms/go-ora/v2"
)

func main() {
	logFileName := "log/insertTest.log"
	//    0644: Permissões do arquivo (leitura/escrita para o dono, leitura para grupo/outros).
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("❌ Erro ao abrir o arquivo de log '%s': %v", logFileName, err)
	}
	defer file.Close()
	//-----------------------------------
	// os.Stdout = file
	//os.Stderr = file
	//	log.SetOutput(file)
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// --
	// 1. Cria um "multi-writer" que escreve tanto para o arquivo quanto para os.Stdout (o terminal).
	// O os.Stdout aqui é o *original* stdout do terminal.
	multiWriter := io.MultiWriter(os.Stdout, file)

	// 2. Redireciona a SAÍDA DO PACOTE LOG padrão do Go para o multi-writer.
	// ISSO É O QUE FAZ log.Print/Println/Printf IREMM PARA AMBOS.
	log.SetOutput(multiWriter)
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//-----------------------------------

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("_________")
	log.Println(db)
	log.Println("_________")

	log.Println("✅ Repositorio conectado ao BD")

	// Defina os parâmetros do teste
	userName := "rafael.oliveira@m4sistemas.com.br"
	//folderPath := "testcase/rest/SequenciaEDA_MVN" // Caminho da pasta com os arquivos
	folderPath := "testcase/provq" // Caminho da pasta com os arquivos
	//folderPath := "testcase/provq/HS" // Caminho da pasta com os arquivos
	//folderPath := "testcase/provq/SIXBELL+VOLTE" // Caminho da pasta com os arquivos
	//folderPath := "testcase/provq/Base_TELCO" // Caminho da pasta com os arquivos
	//folderPath := "testcase/request/MIGRACAO_LIVRE/grupo_1" // Caminho da pasta com os arquivos

	description := ""
	idProject := "116"  // insert test
	idSchemaSPS := "21" //	POS

	// Usuário confirma que alterou a variavel folderPath
	log.Print("❌ ATENÇÃO ❌")
	log.Print("📂 Alterou o valor de folderPath para o diretório desejado?\n")
	log.Print("\n📂 folderPath = " + folderPath + "\n\n📂 Confirma?\t<SIM> Enter \n\t\t<NAO> ctrl + c")
	fmt.Scanln(&folderPath)

	log.Println("__________________________________________________________________________")
	log.Println("userName default: " + userName)
	log.Println("description default: " + description)
	log.Println("idProject default: " + idProject)
	log.Println("idSchemaSPS default: " + idSchemaSPS)

	// Solicita ao usuário para inserir os valores
	log.Println("__________________________________________________________________________")
	log.Print(" # Digite o idProject: ")
	log.Println("->	116  InsertTest")
	log.Println("	476  MSE APIGee")
	log.Println("	296  MassaTesteClaro17/06/25")
	log.Println("  	76   SequenciaEda+VOLTE")
	log.Println("  	201  SequenciaEda")
	log.Println("  	136  VNPSIX+VOLTE")
	log.Println("  	137  MAGNOLIA")
	log.Println("  	138  RSA")
	log.Println("  	156  ESIM")
	log.Println("  	203  VOLTE")
	log.Println("  	277  Hermes")
	log.Println("	336  DRAs_HUAWEI_ERICSSON")
	log.Println("	201  SequenciaEDA_MVN")
	log.Println("	356  AJUSTE CSP SIXBELL VOLTE LEGADO")
	log.Println("	376  MIGRACAO LIV_FIX")
	log.Println("	416  Base_TELCO")
	fmt.Scanln(&idProject)

	log.Println("__________________________________________________________________________")
	log.Print(" # Digite o idSchemaSPS: ")
	log.Println("     1 MVN")
	log.Println("  	  2 NAC")
	log.Println("->	 21 POS")
	log.Println("	101 FLX")
	log.Println("	102 PRE")
	log.Println("	103 HUB")
	log.Println("	 28 FIX")
	fmt.Scanln(&idSchemaSPS)

	log.Println("__________________________________________________________________________")

	// Confirma execução
	log.Print("\n📂 folderPath = " + folderPath + "\n\n📂 Executar insert:\t<SIM> Enter \n\t\t\t<NAO> ctrl + c")
	fmt.Scanln()
	log.Println("__________________________________________________________________________")

	var successList = make([]string, 0)
	var errorList = make([]string, 0)
	// Regex para acao
	regexAcao := regexp.MustCompile(`SRV_TRX_TP_CD=([^;]+);`)
	// Mapeia contagem de quantas vezes cada nome de teste foi inserido
	nameCounts := make(map[string]int)

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
		log.Println("--")
		log.Printf("📂 Processando arquivo: %s\n", filePath)

		// Processa o arquivo (testName é o nome do arquivo sem extensão)
		testName, insert, idType, err := utils.ProcessTestFile(filePath)
		if err != nil {
			log.Printf("❌ Erro ao processar %s: %v", filePath, err)
			continue
		}

		for _, script := range insert {
			var description string
			var baseName string

			description = utils.GeraDescription(script)

			matches := regexAcao.FindStringSubmatch(script)
			if len(matches) > 1 {
				baseName = matches[1]
			} else {
				baseName = testName
			}

			// Incrementa a contagem para este nome específico no map
			nameCounts[baseName]++
			currentCount := nameCounts[baseName]
			//Constrói o nome final com a contagem específica
			finalName := fmt.Sprintf("%s_%d", baseName, currentCount)

			//baseName = utils.ProcessBaseNameLIVRE(baseName) // Processa o nome base para atender às regras de nomenclatura

			//--------------------------------
			// -----------------------------------------------
			// log.Println("----INICIO SIMULA INSERT--------------------------------")
			//
			utils.SimulaInsert(idType /*finalName, */, baseName, description, script, userName, idProject, idSchemaSPS)
			log.Println("----FIM SIMULA INSERT-----------------------------------")
			//-------------------------------------------------------------------------------

			//-------------------------------------------------------------------------------
			// Inserir no banco
			//err = database.InsertTestCase(db, idType, finalName /*/, baseName*/, script, description, userName, idProject, idSchemaSPS)
			if err != nil {
				log.Printf("❌ Erro ao inserir %s no banco: %v" /*finalName */, baseName, err)
				errorList = append(errorList, finalName /*, baseName*/) // Adiciona à lista de erros
				continue
			}
			//-------------------------------------------------------------------------------

			log.Println("✅ Test case inserido no banco:", finalName /*, baseName*/)
			successList = append(successList, finalName /*, baseName*/) // Adiciona à lista de sucesso

		}
	}

	log.Println("🚀 Processamento concluído para todos os arquivos da pasta!")

	// Exibe os resultados compilados
	log.Println("\n### Resultados do insert no banco ###")
	log.Printf("Insert com sucesso:\n")
	for _, success := range successList {
		log.Println("✅", success)
	}

	if len(errorList) > 0 {
		log.Printf("\nInsert com erro:\n")
		for _, failure := range errorList {
			log.Println("❌", failure)
		}
	} else {
		log.Printf("\n ❌ Nenhum Insert com erro ❌\n")
	}

	log.Println("Processamento concluído!")

}
