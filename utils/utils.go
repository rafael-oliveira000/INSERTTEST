package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

type Ordem struct {
	Correlacao []map[string]interface{} `json:"correlacao"`
	Operacao   map[string]interface{}   `json:"operacao"`
	Cliente    map[string]interface{}   `json:"cliente"`
	ItemOrdem  []map[string]interface{} `json:"item-ordem"`
}

func ReadSQLFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// ProcessTestCase verifica o tipo de conteúdo do arquivo e chama a função apropriada.
func ProcessTestCase(filePath string) (string, string, string, error) {
	fileName := filepath.Base(filePath)                              // Pega apenas "arquivo.sql"
	testName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) // Remove ".sql"

	sqlContent := ReadSQLFile(filePath) // Ler o conteúdo do arquivo SQL

	// Determinar qual processamento será aplicado
	switch {
	case regexp.MustCompile(`(?i)INSERT\s+INTO\s+PROV\s*_\s*Q`).MatchString(sqlContent):
		fmt.Println("🔍 Detected: PROV_Q - Chamando ProcessProvQ")
		return ProcessProvQ(testName, sqlContent)

	case regexp.MustCompile(`(?i)INSERT\s+INTO\s+SPS\s*_\s*SOLICITACAO`).MatchString(sqlContent):
		fmt.Println("🔍 Detected: SPS_SOLICITACAO - Chamando ProcessSolicitacao")
		return ProcessSolicitacao(testName, sqlContent)

	case isJSON(sqlContent):
		fmt.Println("🔍 Detected: JSON structure - Chamando ProcessRestJson")
		return ProcessRestJson(testName, sqlContent)

	default:
		err := fmt.Errorf("❌ Nenhuma estrutura identificada no arquivo: %s", filePath)
		fmt.Println(err)
		return "", "", "0", err
	}
}

// isJSON verifica se o conteúdo tem formato JSON válido (simples verificação).
func isJSON(content string) bool {
	content = strings.TrimSpace(content)
	return strings.HasPrefix(content, "{") && strings.HasSuffix(content, "}")
}

// ProcessProvQ - Processa inserts na tabela prov_q
func ProcessProvQ(testName, sqlContent string) (string, string, string, error) {
	var processedLines []string

	// Expressão regular para encontrar comandos INSERT INTO
	insertRegex := regexp.MustCompile(`(?i)INSERT INTO\s+.+\s+VALUES\s*\((.+)\)`)

	// Separar o conteúdo em linhas
	lines := strings.Split(sqlContent, "\n")

	for _, line := range lines {
		// Verifica se a linha corresponde a um INSERT INTO
		if matches := insertRegex.FindStringSubmatch(line); matches != nil {
			values := matches[1] // Captura os valores dentro de VALUES (...)

			// Separar os valores e substituir o primeiro por "<%TCID%>"
			valueParts := strings.SplitN(values, ",", 2)
			if len(valueParts) > 1 {
				modifiedValues := "<%TCID%>, " + valueParts[1]
				modifiedLine := strings.Replace(line, values, modifiedValues, 1)
				processedLines = append(processedLines, modifiedLine)
			} else {
				// Caso não haja separação correta, apenas mantém a linha original
				processedLines = append(processedLines, line)
			}
		}
	}

	// Retornar o nome do teste e os inserts processados
	return testName, strings.Join(processedLines, "\n"), "2", nil
}

// ProcessSolicitacao - Processa inserts na tabela sps_solicitacao
func ProcessSolicitacao(testName, sqlContent string) (string, string, string, error) {
	// TODO: Implementar regras específicas para sps_solicitacao
	return testName, sqlContent, "3", nil
}

// Função para processar o JSON e substituir o valor de "id" dentro de "correlacao" que está dentro de "ordem"
func ProcessRestJson(testName, jsonContent string) (string, string, string, error) {
	// Parse o JSON em uma estrutura genérica
	var jsonData interface{}
	err := json.Unmarshal([]byte(jsonContent), &jsonData)
	if err != nil {
		return "", "", "", fmt.Errorf("erro ao parsear o JSON: %v", err)
	}

	// Função recursiva para percorrer o JSON e modificar o valor de "id" dentro de "correlacao", que está dentro de "ordem"
	var processMap func(map[string]interface{})
	processMap = func(mapData map[string]interface{}) {
		for key, value := range mapData {
			// Verificar se é "ordem" e contém "correlacao"
			if key == "ordem" {
				if ordemMap, ok := value.(map[string]interface{}); ok {
					if correlacaoArr, found := ordemMap["correlacao"].([]interface{}); found {
						for _, item := range correlacaoArr {
							if correlacaoMap, ok := item.(map[string]interface{}); ok {
								// Substituir qualquer valor de "id" por "<%TCID%>"
								if _, idExists := correlacaoMap["id"]; idExists {
									correlacaoMap["id"] = "<%TCID%>"
								}
							}
						}
					}
				}
			}
			// Se o valor for um mapa, chamar recursivamente
			if nestedMap, ok := value.(map[string]interface{}); ok {
				processMap(nestedMap)
			}
		}
	}

	// Aplicar a modificação se o JSON for um mapa
	if dataMap, ok := jsonData.(map[string]interface{}); ok {
		processMap(dataMap)

		// Criar um buffer para armazenar o JSON formatado sem HTML escaping
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false) // Evita que "<" e ">" sejam escapados
		encoder.SetIndent("", "  ")  // Mantém a formatação bonita

		// Codificar o JSON modificado no buffer
		err = encoder.Encode(dataMap)
		if err != nil {
			return "", "", "", fmt.Errorf("erro ao criar item modificado: %v", err)
		}

		// Retornar o JSON formatado corretamente
		return testName, buf.String(), "1", nil
	}

	return "", "", "", fmt.Errorf("... Estrutura inesperada de JSON")
}

// Simular a valores processados no banco exibindo o SQL gerado
func SimulaInsert(idType, testName, processedSQL, userName, idProject, idSchemaSPS string) {
	fmt.Println("==== SQL QUE SERIA INSERIDO ====")
	fmt.Println("ID:", "Incrementado no banco")
	fmt.Println("VERSION:", 1)
	fmt.Println("ID Type:", idType)
	fmt.Println("Nome:", testName)
	fmt.Println("Script: \n", processedSQL)
	fmt.Println("Description: ", "")
	fmt.Println("DT_CREATION: ", "SYSDATE")
	fmt.Println("DT_UPDATE: ", "SYSDATE")
	fmt.Println("Usuário:", userName)
	fmt.Println("Projeto:", idProject)
	fmt.Println("Schema SPS:", idSchemaSPS)
	fmt.Println("================================")
}
